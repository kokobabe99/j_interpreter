package eval

import (
	"fmt"

	gen "interpreter/parser_src"
)

type ValueKind int

const (
	NilKind ValueKind = iota
	IntKind
	BoolKind
	StringKind
	FuncKind
	ChanKind // ✅ channel
)

type Channel struct {
	Buf    []Value
	Closed bool
}

type Value struct {
	Kind ValueKind
	I    int64
	B    bool
	S    string
	F    *Function
	C    *Channel
}

func VNil() Value             { return Value{Kind: NilKind} }
func VInt(v int64) Value      { return Value{Kind: IntKind, I: v} }
func VBool(v bool) Value      { return Value{Kind: BoolKind, B: v} }
func VString(v string) Value  { return Value{Kind: StringKind, S: v} }
func VFunc(f *Function) Value { return Value{Kind: FuncKind, F: f} }

// ✅ 统一的 channel 构造：带 capacity
func VChan(capacity int64) Value {
	if capacity < 0 {
		capacity = 0
	}
	return Value{
		Kind: ChanKind,
		C: &Channel{
			Buf:    make([]Value, 0, capacity),
			Closed: false,
		},
	}
}

func (v Value) String() string {
	switch v.Kind {
	case NilKind:
		return "nil"
	case IntKind:
		return fmt.Sprintf("%d", v.I)
	case BoolKind:
		return fmt.Sprintf("%v", v.B)
	case StringKind:
		return fmt.Sprintf("%q", v.S)
	case FuncKind:
		return "<func>"
	case ChanKind:
		if v.C == nil {
			return "<chan nil>"
		}
		if v.C.Closed {
			return "<chan closed>"
		}
		return "<chan>"
	default:
		return "<unknown>"
	}
}

type Function struct {
	Name       string
	Params     []string
	ParamKinds []ValueKind
	Body       gen.IBlockContext
	Env        *Env
}

type Env struct {
	Parent *Env
	Vars   map[string]Value

	Defers []DeferredCall

	Panic *string
}

type DeferredCall struct {
	Callee Value
	Args   []Value
}

func NewEnv(parent *Env) *Env {
	return &Env{
		Parent: parent,
		Vars:   make(map[string]Value),
	}
}

func (e *Env) Get(name string) (Value, bool) {
	cur := e
	for cur != nil {
		if v, ok := cur.Vars[name]; ok {
			return v, true
		}
		cur = cur.Parent
	}
	return VNil(), false
}

func (e *Env) Set(name string, v Value) {
	e.Vars[name] = v
}

func (e *Env) SetExisting(name string, v Value) bool {
	if _, ok := e.Vars[name]; ok {
		e.Vars[name] = v
		return true
	}
	if e.Parent != nil {
		return e.Parent.SetExisting(name, v)
	}
	return false
}

// 用 panic 传递的“控制流信号”
type breakSignal struct{}
type continueSignal struct{}
type returnSignal struct{ Val Value }
type panicSignal struct{ Msg string }
type jotoSignal struct {
	Label string
}

func (k ValueKind) String() string {
	switch k {
	case NilKind:
		return "nil"
	case IntKind:
		return "int"
	case BoolKind:
		return "bool"
	case StringKind:
		return "string"
	case FuncKind:
		return "func"
	case ChanKind:
		return "chan"
	default:
		return "unknown"
	}
}
