// eval/interpreter.go
package eval

import (
	"strconv"
	"strings"

	gen "interpreter/parser_src"
)

/*************** Interpreter ***************/

type Interpreter struct {
	Parser *gen.JLangParser
	Root   *Env // 全局环境
}

func NewInterpreter(p *gen.JLangParser) *Interpreter {
	root := NewEnv(nil)
	InstallBuiltins(root)
	return &Interpreter{
		Parser: p,
		Root:   root,
	}
}

/*************** Top-level ***************/

func (ip *Interpreter) Run(prog gen.IProgramContext) {
	ip.execTop(prog)
	if v, ok := ip.Root.Get("main"); ok && v.Kind == FuncKind {
		ip.callValue(v, []Value{}, ip.Root)
	}
}

func (ip *Interpreter) execTop(p gen.IProgramContext) {
	for _, td := range p.AllTopDecl() {
		ip.execTopDecl(td)
	}
}

func (ip *Interpreter) execTopDecl(ctx gen.ITopDeclContext) {
	if v := ctx.VarDecl(); v != nil {
		ip.execVarDecl(ip.Root, v)
		return
	}
	if c := ctx.ConsDecl(); c != nil {
		ip.execConsDecl(ip.Root, c)
		return
	}
	if t := ctx.TypeDecl(); t != nil {
		_ = t // v0: 忽略类型声明
		return
	}
	if f := ctx.FuncDecl(); f != nil {
		ip.execFuncDecl(ip.Root, f)
		return
	}
	if s := ctx.Stmt(); s != nil {
		ip.execStmt(ip.Root, s)
		return
	}
}

/*************** Decls ***************/

func (ip *Interpreter) execFuncDecl(env *Env, f gen.IFuncDeclContext) {
	name := f.IDENT().GetText()
	params := []string{}
	if sig := f.Signature(); sig != nil {
		if pl := sig.ParamList(); pl != nil {
			for _, p := range pl.AllParam() {
				params = append(params, p.IDENT().GetText())
			}
		}
	}
	fn := &Function{
		Name:   name,
		Params: params,
		Body:   f.Block(), // gen.IBlockContext
		Env:    env,       // 闭包环境
	}
	env.Set(name, VFunc(fn))
}

func (ip *Interpreter) execVarDecl(env *Env, v gen.IVarDeclContext) {
	for _, spec := range v.AllVarSpec() {
		names := idsOf(spec.IdentList())
		var vals []Value
		if spec.ExprList() != nil {
			vals = ip.evalExprList(env, spec.ExprList())
		} else {
			vals = make([]Value, len(names))
			for i := range vals {
				vals[i] = VNil()
			}
		}
		for i, name := range names {
			env.Set(name, vals[min(i, len(vals)-1)])
		}
	}
}

func (ip *Interpreter) execConsDecl(env *Env, c gen.IConsDeclContext) {
	for _, spec := range c.AllConstSpec() {
		names := idsOf(spec.IdentList())
		vals := ip.evalExprList(env, spec.ExprList())
		for i, name := range names {
			env.Set(name, vals[min(i, len(vals)-1)])
		}
	}
}

/*************** Stmts ***************/

func (ip *Interpreter) execStmt(env *Env, s gen.IStmtContext) {
	// 先处理 simpleStmt
	if ss := s.SimpleStmt(); ss != nil {
		ip.execSimpleStmt(env, ss)
		return
	}

	switch {
	case s.DeclStmt() != nil:
		ds := s.DeclStmt()
		if v := ds.VarDecl(); v != nil {
			ip.execVarDecl(env, v)
			return
		}
		if c := ds.ConsDecl(); c != nil {
			ip.execConsDecl(env, c)
			return
		}
		// typeDecl v0 忽略
		return

	case s.IfStmt() != nil:
		ip.execIf(env, s.IfStmt())
		return

	case s.SwitchStmt() != nil:
		// v0 未实现
		return

	case s.ForStmt() != nil:
		// v0 未实现
		return

	case s.RangeStmt() != nil:
		// 这里是 channel range 的简易实现
		rs := s.RangeStmt()
		h := rs.RangeHeader()

		// 目前忽略 "(identList DECL)?" 的绑定，只读取右侧 expr 作为 channel
		ch := ip.evalExpr(env, h.Expr())
		if ch.Kind != ChanKind || ch.C == nil {
			panic("range over non-channel")
		}

		// 最小可跑：把当前缓冲里的元素全部消费
		for len(ch.C.Buf) > 0 {
			_ = ch.C.Buf[0] // 你后面可以在这里绑定 x := ...
			ch.C.Buf = ch.C.Buf[1:]
			ip.execBlock(env, rs.Block())
		}
		return

	case s.SelectStmt() != nil:
		// v0 未实现
		return

	case s.ReturnStmt() != nil:
		r := s.ReturnStmt()
		var v Value = VNil()
		if r.ExprList() != nil {
			v = ip.evalExprList(env, r.ExprList())[0]
		}
		panic(returnSignal{Val: v})

	case s.BreakStmt() != nil:
		// v0 未实现循环，暂忽略
		return

	case s.ContinueStmt() != nil:
		// v0 未实现循环，暂忽略
		return

	case s.JotoStmt() != nil:
		// v0 未实现 label 跳转
		return

	case s.PanicStmt() != nil:
		msg := ip.evalExpr(env, s.PanicStmt().Expr())
		panic(panicSignal{Msg: toString(msg)})

	case s.Block() != nil:
		ip.execBlock(env, s.Block())
		return
	}
}

func (ip *Interpreter) execSimpleStmt(env *Env, ss gen.ISimpleStmtContext) {
	switch {
	case ss.ExprStmt() != nil:
		ip.evalExpr(env, ss.ExprStmt().Expr())
		return

	case ss.AssignStmt() != nil:
		l := ss.AssignStmt()
		vals := ip.evalExprList(env, l.ExprList())
		// MVP：仅支持 IDENT 左值
		pe := l.Lhs().PrimaryExpr()
		if pe == nil || pe.Operand() == nil || pe.Operand().IDENT() == nil {
			panic("only IDENT assignment supported in MVP")
		}
		id := pe.Operand().IDENT()
		if !env.SetExisting(id.GetText(), vals[0]) {
			env.Set(id.GetText(), vals[0])
		}
		return

	case ss.ShortVarDecl() != nil:
		d := ss.ShortVarDecl()
		ids := idsOf(d.IdentList())
		vals := ip.evalExprList(env, d.ExprList())
		for i, name := range ids {
			env.Set(name, vals[min(i, len(vals)-1)])
		}
		return

	case ss.SendStmt() != nil:
		s := ss.SendStmt()
		ch := ip.evalExpr(env, s.Expr(0))
		val := ip.evalExpr(env, s.Expr(1))
		if ch.Kind != ChanKind || ch.C == nil {
			panic("send to non-channel")
		}
		if ch.C.Closed {
			panic("send on closed channel")
		}
		ch.C.Buf = append(ch.C.Buf, val)
		return

	case ss.SpawnStmt() != nil:
		// v0：把 j foo(args) 当成立即求值（或在此处排入任务队列）
		_ = ip.evalExpr(env, ss.SpawnStmt().Expr())
		return
	}
}

func (ip *Interpreter) execIf(env *Env, i gen.IIfStmtContext) {
	cond := ip.evalExpr(env, i.Expr())
	if toBool(cond) {
		ip.execBlock(env, i.Block())
	} else if i.ElseOpt() != nil {
		eo := i.ElseOpt()
		if eo.IfStmt() != nil {
			ip.execIf(env, eo.IfStmt())
		} else if eo.Block() != nil {
			ip.execBlock(env, eo.Block())
		}
	}
}

func (ip *Interpreter) execBlock(outer *Env, b gen.IBlockContext) (ret Value) {
	env := NewEnv(outer)

	defer func() {
		if r := recover(); r != nil {
			switch sig := r.(type) {
			case returnSignal:
				ret = sig.Val
				return

			case panicSignal:
				env.Panic = &sig.Msg

				// 先跑 later（defer）
				for i := len(env.Defers) - 1; i >= 0; i-- {
					ip.callValue(env.Defers[i].Callee, env.Defers[i].Args, env)
				}

				if env.Panic != nil {
					// 仍未 recover，继续向上传播 panic
					panic(sig)
				}
				return

			default:
				panic(r)
			}
		}

		// 正常路径：同样需要执行 later
		for i := len(env.Defers) - 1; i >= 0; i-- {
			ip.callValue(env.Defers[i].Callee, env.Defers[i].Args, env)
		}
	}()

	if b.StmtList() == nil {
		return VNil()
	}
	for _, s := range b.StmtList().AllStmt() {
		ip.execStmt(env, s)
	}
	return VNil()
}

/*************** Exprs ***************/

func (ip *Interpreter) evalExprList(env *Env, el gen.IExprListContext) []Value {
	out := []Value{}
	for _, e := range el.AllExpr() {
		out = append(out, ip.evalExpr(env, e))
	}
	return out
}

// 顶层：expr : logicalOrExpr ;
func (ip *Interpreter) evalExpr(env *Env, e gen.IExprContext) Value {
	return ip.evalLogicalOr(env, e.LogicalOrExpr())
}

/*** 分层表达式求值，对应新的 .g4 ***/

// logicalOrExpr : logicalAndExpr (OROR logicalAndExpr)* ;
func (ip *Interpreter) evalLogicalOr(env *Env, ctx gen.ILogicalOrExprContext) Value {

	if ctx == nil {
		return VNil()
	}
	items := ctx.AllLogicalAndExpr()
	if len(items) == 0 {
		return VNil()
	}

	// 先算第一个
	v := ip.evalLogicalAnd(env, items[0])

	// 如果只有一个项，直接返回原值（可能是 int / chan / string / bool 等）
	if len(items) == 1 {
		return v
	}

	// 有多个项时，才做 || 逻辑，结果是 bool
	for i := 1; i < len(items); i++ {
		if toBool(v) {
			// 短路：前面已经为真
			return VBool(true)
		}
		v = ip.evalLogicalAnd(env, items[i])
	}
	return VBool(toBool(v))
}

// logicalAndExpr : bitOrExpr (ANDAND bitOrExpr)* ;
func (ip *Interpreter) evalLogicalAnd(env *Env, ctx gen.ILogicalAndExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	items := ctx.AllBitOrExpr()
	if len(items) == 0 {
		return VNil()
	}

	// 先算第一个
	v := ip.evalBitOr(env, items[0])

	// 只有一个项时，直接返回原值
	if len(items) == 1 {
		return v
	}

	// 多个项时才做 && 逻辑
	for i := 1; i < len(items); i++ {
		if !toBool(v) {
			// 短路：前面已经为假
			return VBool(false)
		}
		v = ip.evalBitOr(env, items[i])
	}
	return VBool(toBool(v))
}

// bitOrExpr : bitXorExpr (BOR bitXorExpr)* ;
func (ip *Interpreter) evalBitOr(env *Env, ctx gen.IBitOrExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	items := ctx.AllBitXorExpr()
	if len(items) == 0 {
		return VNil()
	}
	v := ip.evalBitXor(env, items[0])
	for i := 1; i < len(items); i++ {
		v = VInt(toInt(v) | toInt(ip.evalBitXor(env, items[i])))
	}
	return v
}

// bitXorExpr : bitAndExpr (BXOR bitAndExpr)* ;
func (ip *Interpreter) evalBitXor(env *Env, ctx gen.IBitXorExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	items := ctx.AllBitAndExpr()
	if len(items) == 0 {
		return VNil()
	}
	v := ip.evalBitAnd(env, items[0])
	for i := 1; i < len(items); i++ {
		v = VInt(toInt(v) ^ toInt(ip.evalBitAnd(env, items[i])))
	}
	return v
}

// bitAndExpr : equalityExpr (BAND equalityExpr)* ;
func (ip *Interpreter) evalBitAnd(env *Env, ctx gen.IBitAndExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	items := ctx.AllEqualityExpr()
	if len(items) == 0 {
		return VNil()
	}
	v := ip.evalEquality(env, items[0])
	for i := 1; i < len(items); i++ {
		v = VInt(toInt(v) & toInt(ip.evalEquality(env, items[i])))
	}
	return v
}

// equalityExpr : relationalExpr ((EQ | NE) relationalExpr)* ;
func (ip *Interpreter) evalEquality(env *Env, ctx gen.IEqualityExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	rel := ctx.AllRelationalExpr()
	if len(rel) == 0 {
		return VNil()
	}
	if len(rel) == 1 {
		return ip.evalRelational(env, rel[0])
	}

	left := ip.evalRelational(env, rel[0])
	right := ip.evalRelational(env, rel[1])

	// 只看第一个操作符
	if ctx.GetChildCount() >= 2 {
		if node, ok := ctx.GetChild(1).(interface{ GetText() string }); ok {
			op := node.GetText()
			return VBool(compare(op, left, right))
		}
	}
	return VBool(compare("==", left, right))
}

// relationalExpr : shiftExpr ((LT | LE | GT | GE) shiftExpr)* ;
func (ip *Interpreter) evalRelational(env *Env, ctx gen.IRelationalExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	shift := ctx.AllShiftExpr()
	if len(shift) == 0 {
		return VNil()
	}
	if len(shift) == 1 {
		return ip.evalShift(env, shift[0])
	}

	left := ip.evalShift(env, shift[0])
	right := ip.evalShift(env, shift[1])

	if ctx.GetChildCount() >= 2 {
		if node, ok := ctx.GetChild(1).(interface{ GetText() string }); ok {
			op := node.GetText()
			return VBool(compare(op, left, right))
		}
	}
	return VBool(compare("==", left, right))
}

// shiftExpr : additiveExpr ((SHL | SHR) additiveExpr)* ;
func (ip *Interpreter) evalShift(env *Env, ctx gen.IShiftExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	items := ctx.AllAdditiveExpr()
	if len(items) == 0 {
		return VNil()
	}
	if len(items) == 1 {
		return ip.evalAdditive(env, items[0])
	}

	cur := ip.evalAdditive(env, items[0])
	for i := 1; i < len(items); i++ {
		right := ip.evalAdditive(env, items[i])
		shift := toInt(right)

		// 操作符在 child(2*i-1)
		var op string
		idx := 2*i - 1
		if idx < ctx.GetChildCount() {
			if node, ok := ctx.GetChild(idx).(interface{ GetText() string }); ok {
				op = node.GetText()
			}
		}
		if op == ">>" {
			cur = VInt(toInt(cur) >> shift)
		} else {
			// 默认 <<
			cur = VInt(toInt(cur) << shift)
		}
	}
	return cur
}

// additiveExpr : multiplicativeExpr ((PLUS | MINUS) multiplicativeExpr)* ;
func (ip *Interpreter) evalAdditive(env *Env, ctx gen.IAdditiveExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	items := ctx.AllMultiplicativeExpr()
	if len(items) == 0 {
		return VNil()
	}
	if len(items) == 1 {
		return ip.evalMultiplicative(env, items[0])
	}

	cur := ip.evalMultiplicative(env, items[0])
	for i := 1; i < len(items); i++ {
		right := ip.evalMultiplicative(env, items[i])
		idx := 2*i - 1
		op := "+"
		if idx < ctx.GetChildCount() {
			if node, ok := ctx.GetChild(idx).(interface{ GetText() string }); ok {
				op = node.GetText()
			}
		}
		a, b := toInt(cur), toInt(right)
		if op == "-" {
			cur = VInt(a - b)
		} else {
			cur = VInt(a + b)
		}
	}
	return cur
}

// multiplicativeExpr : unaryExpr ((STAR | SLASH | PERCENT) unaryExpr)* ;
func (ip *Interpreter) evalMultiplicative(env *Env, ctx gen.IMultiplicativeExprContext) Value {
	if ctx == nil {
		return VNil()
	}
	items := ctx.AllUnaryExpr()
	if len(items) == 0 {
		return VNil()
	}
	if len(items) == 1 {
		return ip.evalUnary(env, items[0])
	}

	cur := ip.evalUnary(env, items[0])
	for i := 1; i < len(items); i++ {
		right := ip.evalUnary(env, items[i])
		idx := 2*i - 1
		op := "*"
		if idx < ctx.GetChildCount() {
			if node, ok := ctx.GetChild(idx).(interface{ GetText() string }); ok {
				op = node.GetText()
			}
		}
		a, b := toInt(cur), toInt(right)
		switch op {
		case "*":
			cur = VInt(a * b)
		case "/":
			if b == 0 {
				panic(panicSignal{Msg: "division by zero"})
			}
			cur = VInt(a / b)
		case "%":
			cur = VInt(a % b)
		default:
			cur = VInt(a * b)
		}
	}
	return cur
}

func (ip *Interpreter) evalUnary(env *Env, u gen.IUnaryExprContext) Value {
	// unaryExpr : primaryExpr | PLUS unaryExpr | ...
	if u.PrimaryExpr() != nil {
		return ip.evalPrimary(env, u.PrimaryExpr())
	}
	inner := ip.evalUnary(env, u.UnaryExpr())
	switch {
	case u.PLUS() != nil:
		return inner
	case u.MINUS() != nil:
		return VInt(-toInt(inner))
	case u.BANG() != nil:
		return VBool(!toBool(inner))
	case u.TILDE() != nil:
		return VInt(^toInt(inner))
	case u.STAR() != nil:
		// 解引用未实现
		return inner
	case u.BAND() != nil:
		// 取址未实现
		return inner
	default:
		return inner
	}
}

func (ip *Interpreter) evalPrimary(env *Env, p gen.IPrimaryExprContext) Value {
	px, ok := p.(*gen.PrimaryExprContext)
	if !ok || px == nil {
		return VNil()
	}

	// 1. 最底层：operand 或 makeExpr
	if px.Operand() != nil {
		return ip.evalOperand(env, px.Operand())
	}
	if px.MakeExpr() != nil {
		return ip.evalMakeExpr(env, px.MakeExpr())
	}

	// 2. 链式：primaryExpr selector/index/call
	if px.PrimaryExpr() == nil {
		return VNil()
	}
	cur := ip.evalPrimary(env, px.PrimaryExpr())

	// selector: x.y
	if px.Selector() != nil {
		name := px.Selector().IDENT().GetText()
		// 简单模拟 package/对象成员：如果 cur.F.Env 不为空，就从那里面取
		if cur.Kind == FuncKind && cur.F.Env != nil {
			if v, ok := cur.F.Env.Get(name); ok {
				return v
			}
		}
		return VNil()
	}

	// call: f(args...)
	if px.Call() != nil {
		var args []Value
		if px.Call().ArgList() != nil {
			for _, e := range px.Call().ArgList().AllExpr() {
				args = append(args, ip.evalExpr(env, e))
			}
		}
		return ip.callValue(cur, args, env)
	}

	// index: v0 未实现（数组/切片）
	if px.Index() != nil {
		return VNil()
	}

	return cur
}

func (ip *Interpreter) evalMakeExpr(env *Env, m gen.IMakeExprContext) Value {
	t := m.Type_()
	if t == nil {
		return VNil()
	}

	// 只支援 channel 类型：make(channel string, 2)
	isChan := false
	if t.ChannelType() != nil {
		isChan = true
	}

	if !isChan {
		// 目前只实现 channel 的 make
		return VNil()
	}

	// cap 参数（可选）
	capacity := int64(0)
	if m.ExprList() != nil && len(m.ExprList().AllExpr()) > 0 {
		capacity = toInt(ip.evalExpr(env, m.ExprList().Expr(0)))
		if capacity < 0 {
			capacity = 0
		}
	}

	// ✅ 用上面刚刚统一好的 VChan
	return VChan(capacity)
}
func (ip *Interpreter) evalOperand(env *Env, o gen.IOperandContext) Value {
	switch {
	case o.IDENT() != nil:
		if v, ok := env.Get(o.IDENT().GetText()); ok {
			return v
		}
		return VNil()

	case o.INT_LIT() != nil:
		s := strings.ReplaceAll(o.INT_LIT().GetText(), "_", "")
		n, _ := strconv.ParseInt(parseIntLike(s), 0, 64)
		return VInt(n)

	case o.FLOAT_LIT() != nil:
		// 目前没有专门的 float 类型，这里先简单转成字符串或忽略
		s := strings.ReplaceAll(o.FLOAT_LIT().GetText(), "_", "")
		f, _ := strconv.ParseFloat(s, 64)
		return VInt(int64(f)) // 先粗暴 cast，之后你可以加 FloatKind
	case o.STRING_LIT() != nil:
		text := o.STRING_LIT().GetText()
		unq := strings.Trim(text, "\"")
		unq = strings.ReplaceAll(unq, "\\\"", "\"")
		return VString(unq)

	case o.RAW_STR() != nil:
		return VString(strings.Trim(o.RAW_STR().GetText(), "`"))

	case o.CHAR_LIT() != nil:
		// 简单处理：把字符当成长度为 1 的 string
		text := o.CHAR_LIT().GetText()
		return VString(strings.Trim(text, "'"))

	case o.KW_RECOVER() != nil:
		// 语法已经保证 recover() 形式，所以这里直接返回一个特殊函数占位
		return VFunc(&Function{Name: "recover"})
	}

	if o.LPAREN() != nil && o.RPAREN() != nil {
		return ip.evalExpr(env, o.Expr())
	}
	return VNil()
}

/*************** Calls ***************/

func (ip *Interpreter) callValue(callee Value, args []Value, cur *Env) Value {
	if callee.Kind != FuncKind || callee.F == nil {
		return VNil()
	}
	fn := callee.F

	// 内建（包含 recover / io.Println 等）
	if v, ok := CallBuiltin(fn, args, cur); ok {
		return v
	}

	// 用户函数
	newEnv := NewEnv(fn.Env)
	for i, p := range fn.Params {
		if i < len(args) {
			newEnv.Set(p, args[i])
		} else {
			newEnv.Set(p, VNil())
		}
	}

	// 运行函数体（由外层 execBlock 统一处理 later / panic / ret）
	val := ip.execBlock(newEnv, fn.Body)
	return val
}

/*************** later 解析辅助 ***************/

// 从 expr 一路往右，拿到最右侧的 primaryExpr
func rightmostPrimary(e gen.IExprContext) gen.IPrimaryExprContext {
	if e == nil || e.LogicalOrExpr() == nil {
		return nil
	}
	return rightmostPrimaryFromLogicalOr(e.LogicalOrExpr())
}

func rightmostPrimaryFromLogicalOr(ctx gen.ILogicalOrExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllLogicalAndExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromLogicalAnd(list[len(list)-1])
}

func rightmostPrimaryFromLogicalAnd(ctx gen.ILogicalAndExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllBitOrExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromBitOr(list[len(list)-1])
}

func rightmostPrimaryFromBitOr(ctx gen.IBitOrExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllBitXorExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromBitXor(list[len(list)-1])
}

func rightmostPrimaryFromBitXor(ctx gen.IBitXorExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllBitAndExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromBitAnd(list[len(list)-1])
}

func rightmostPrimaryFromBitAnd(ctx gen.IBitAndExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllEqualityExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromEquality(list[len(list)-1])
}

func rightmostPrimaryFromEquality(ctx gen.IEqualityExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllRelationalExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromRelational(list[len(list)-1])
}

func rightmostPrimaryFromRelational(ctx gen.IRelationalExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllShiftExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromShift(list[len(list)-1])
}

func rightmostPrimaryFromShift(ctx gen.IShiftExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllAdditiveExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromAdditive(list[len(list)-1])
}

func rightmostPrimaryFromAdditive(ctx gen.IAdditiveExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllMultiplicativeExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromMultiplicative(list[len(list)-1])
}

func rightmostPrimaryFromMultiplicative(ctx gen.IMultiplicativeExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	list := ctx.AllUnaryExpr()
	if len(list) == 0 {
		return nil
	}
	return rightmostPrimaryFromUnary(list[len(list)-1])
}

func rightmostPrimaryFromUnary(ctx gen.IUnaryExprContext) gen.IPrimaryExprContext {
	if ctx == nil {
		return nil
	}
	if ctx.PrimaryExpr() != nil {
		return ctx.PrimaryExpr()
	}
	if ctx.UnaryExpr() != nil {
		return rightmostPrimaryFromUnary(ctx.UnaryExpr())
	}
	return nil
}

// later 后紧跟调用：抽出 callee/args（支持 IDENT(...) 或 io.Println(...) 链式）
func (ip *Interpreter) extractCall(env *Env, e gen.IExprContext) (Value, []Value) {
	if pe := rightmostPrimary(e); pe != nil {
		if v, args := ip.extractCallPrimary(env, pe); v.Kind == FuncKind || len(args) > 0 {
			return v, args
		}
	}
	// 回退：立即求值（非显式调用形态）
	v := ip.evalExpr(env, e)
	return v, nil
}

func (ip *Interpreter) extractCallPrimary(env *Env, p gen.IPrimaryExprContext) (Value, []Value) {
	var cur Value
	px, ok := p.(*gen.PrimaryExprContext)
	if !ok || px == nil {
		return VNil(), nil
	}

	if px.Operand() != nil {
		cur = ip.evalOperand(env, px.Operand())
	} else if px.PrimaryExpr() != nil {
		cur, _ = ip.extractCallPrimary(env, px.PrimaryExpr())
	}

	// selector
	if px.Selector() != nil {
		name := px.Selector().IDENT().GetText()
		if cur.Kind == FuncKind && cur.F.Env != nil {
			if v, ok := cur.F.Env.Get(name); ok {
				cur = v
			}
		}
	}

	// call
	if px.Call() != nil {
		var args []Value
		if px.Call().ArgList() != nil {
			for _, e := range px.Call().ArgList().AllExpr() {
				args = append(args, ip.evalExpr(env, e))
			}
		}
		return cur, args
	}
	return cur, nil
}

/*************** Helpers ***************/

func idsOf(il gen.IIdentListContext) []string {
	out := []string{}
	for _, id := range il.AllIDENT() {
		out = append(out, id.GetText())
	}
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func toInt(v Value) int64 {
	switch v.Kind {
	case IntKind:
		return v.I
	case BoolKind:
		if v.B {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func toBool(v Value) bool {
	switch v.Kind {
	case BoolKind:
		return v.B
	case IntKind:
		return v.I != 0
	case StringKind:
		return v.S != ""
	default:
		return false
	}
}

func toString(v Value) string {
	switch v.Kind {
	case StringKind:
		return v.S
	case IntKind:
		return strconv.FormatInt(v.I, 10)
	case BoolKind:
		if v.B {
			return "true"
		}
		return "false"
	default:
		return ""
	}
}

func parseIntLike(s string) string {
	// 支持 0x/0b/0o/0 前缀：strconv.ParseInt(…, 0, …) 已处理
	return s
}

func opOf(e gen.IExprContext) string {
	switch {
	// 旧版用的，现在保留防以后你想复用
	default:
		return ""
	}
}

func compare(op string, a, b Value) bool {
	if a.Kind == IntKind && b.Kind == IntKind {
		switch op {
		case "==":
			return a.I == b.I
		case "!=":
			return a.I != b.I
		case "<":
			return a.I < b.I
		case "<=":
			return a.I <= b.I
		case ">":
			return a.I > b.I
		case ">=":
			return a.I >= b.I
		}
		return a.I == b.I
	}
	// 其它类型：按字符串比较
	as, bs := a.String(), b.String()
	switch op {
	case "==":
		return as == bs
	case "!=":
		return as != bs
	case "<":
		return as < bs
	case "<=":
		return as <= bs
	case ">":
		return as > bs
	case ">=":
		return as >= bs
	default:
		return as == bs
	}
}
