package eval

import "fmt"

func InstallBuiltins(root *Env) {
	// namespace: io
	ioNS := NewEnv(nil)
	ioNS.Set("Println", VFunc(&Function{
		Name:   "io.Println",
		Params: []string{"v"},
		Body:   nil, // native
		Env:    nil,
	}))
	// 用一个函数值来模拟 namespace: io
	root.Set("io", VFunc(&Function{Name: "<namespace:io>", Env: ioNS}))

	// builtin: recover()
	root.Set("recover", VFunc(&Function{
		Name:   "recover",
		Params: []string{},
		Body:   nil,
		Env:    nil,
	}))

	// builtin: close(ch)
	root.Set("close", VFunc(&Function{
		Name:   "close",
		Params: []string{"ch"},
		Body:   nil,
		Env:    nil,
	}))
}

func CallBuiltin(fn *Function, args []Value, env *Env) (Value, bool) {

	switch fn.Name {
	case "recover":
		if env.Panic != nil {
			msg := *env.Panic
			env.Panic = nil // ✅ 清空，表示已恢复
			return VString(msg), true
		}
		return VString(""), true

	case "io.Println":
		fmt.Println(joinForPrint(args))
		return VNil(), true

	case "close":
		if len(args) > 0 && args[0].Kind == ChanKind && args[0].C != nil {
			args[0].C.Closed = true
		}
		return VNil(), true

	case "make":
		var cap int64
		if len(args) > 1 {
			cap = toInt(args[1])
		}
		return VChan(cap), true
	}

	return VNil(), false
}

func joinForPrint(xs []Value) string {
	out := ""
	for i, v := range xs {
		if i > 0 {
			out += " "
		}
		switch v.Kind {
		case StringKind:
			out += v.S
		default:
			out += v.String()
		}
	}
	return out
}
