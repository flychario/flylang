package ast

import "fmt"

type Builtin struct {
	Name string
	Args []Element
	Code func(*Context, []Element) Element
}

func (b Builtin) ElementType() ElementType {
	return ElementTypeLambda
}

func (b Builtin) Eval(c *Context) Element {
	return b
}

func (b Builtin) Call(c *Context, args []Element) Element {
	if len(args) != len(b.Args) {
		panic(fmt.Sprintf("Wrong number of arguments to %s: %d != %d", b.Name, len(args), len(b.Args)))
	}
	return b.Code(c, args)
}

func (b Builtin) GetElements() []Element {
	return nil
}

func (b Builtin) String() string {
	return b.Name
}

var Builtins = []Builtin{
	{
		Name: "plus",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				panic(fmt.Sprintf("Can't add %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)

			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralInteger{a.(LiteralInteger).Value + b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralReal{a.(LiteralReal).Value + b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralReal{float64(a.(LiteralInteger).Value) + b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralReal{a.(LiteralReal).Value + float64(b.(LiteralInteger).Value)}
			}
			panic(fmt.Sprintf("Can't add %s and %s", a, b))
		},
	},
	{
		Name: "minus",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				panic(fmt.Sprintf("Can't add %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)

			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralInteger{a.(LiteralInteger).Value - b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralReal{a.(LiteralReal).Value - b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralReal{float64(a.(LiteralInteger).Value) - b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralReal{a.(LiteralReal).Value - float64(b.(LiteralInteger).Value)}
			}
			panic(fmt.Sprintf("Can't subtract %s and %s", a, b))
		},
	},
	{
		Name: "times",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				panic(fmt.Sprintf("Can't add %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)
			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralInteger{a.(LiteralInteger).Value * b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralReal{a.(LiteralReal).Value * b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralReal{float64(a.(LiteralInteger).Value) * b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralReal{a.(LiteralReal).Value * float64(b.(LiteralInteger).Value)}
			}
			panic(fmt.Sprintf("Can't multiply %s and %s", a, b))
		},
	},
	{
		Name: "divide",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				panic(fmt.Sprintf("Can't add %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)
			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralInteger{a.(LiteralInteger).Value / b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralReal{a.(LiteralReal).Value / b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralReal{float64(a.(LiteralInteger).Value) / b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralReal{a.(LiteralReal).Value / float64(b.(LiteralInteger).Value)}
			}
			panic(fmt.Sprintf("Can't divide %s and %s", a, b))
		},
	},
	{
		Name: "equal",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				panic(fmt.Sprintf("Can't add %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)
			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralInteger).Value == b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralBoolean{a.(LiteralReal).Value == b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralBoolean{float64(a.(LiteralInteger).Value) == b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralReal).Value == float64(b.(LiteralInteger).Value)}
			} else if a.Type() == LiteralTypeBoolean && b.Type() == LiteralTypeBoolean {
				return LiteralBoolean{a.(LiteralBoolean).Value == b.(LiteralBoolean).Value}
			} else if a.Type() == LiteralTypeNull && b.Type() == LiteralTypeNull {
				return LiteralBoolean{true}
			}
			panic(fmt.Sprintf("Can't compare %s and %s", a, b))
		},
	},
	{
		Name: "lesseq",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				panic(fmt.Sprintf("Can't add %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)
			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralInteger).Value <= b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralBoolean{a.(LiteralReal).Value <= b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralBoolean{float64(a.(LiteralInteger).Value) <= b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralReal).Value <= float64(b.(LiteralInteger).Value)}
			}
			panic(fmt.Sprintf("Can't compare %s and %s", a, b))
		},
	},
}

func initBuiltins(c *Context) {
	for _, b := range Builtins {
		c.Add(b.Name, b)
	}
}
