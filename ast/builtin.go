package ast

import (
	"fmt"
)

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
		CreateEvaluateError(fmt.Sprintf("Wrong number of arguments to %s: %d != %d", b.Name, len(args), len(b.Args)))
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
				CreateEvaluateError(fmt.Sprintf("Can't add %v and %v", ae, be))
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
			CreateEvaluateError(fmt.Sprintf("Can't add %v and %v", a, b))
			return nil
		},
	},
	{
		Name: "minus",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't add %s and %s", ae, be))
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
			CreateEvaluateError(fmt.Sprintf("Can't subtract %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "times",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't add %s and %s", ae, be))
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
			CreateEvaluateError(fmt.Sprintf("Can't multiply %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "divide",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't add %s and %s", ae, be))
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
			CreateEvaluateError(fmt.Sprintf("Can't divide %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "equal",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", ae, be))
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
			CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "nonequal",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)
			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralInteger).Value != b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralBoolean{a.(LiteralReal).Value != b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralBoolean{float64(a.(LiteralInteger).Value) != b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralReal).Value != float64(b.(LiteralInteger).Value)}
			} else if a.Type() == LiteralTypeBoolean && b.Type() == LiteralTypeBoolean {
				return LiteralBoolean{a.(LiteralBoolean).Value != b.(LiteralBoolean).Value}
			} else if a.Type() == LiteralTypeNull && b.Type() == LiteralTypeNull {
				return LiteralBoolean{false}
			} else if a.Type() == LiteralTypeNull && b.Type() == LiteralTypeNull {
				return LiteralBoolean{false}
			}
			CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "less",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)
			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralInteger).Value < b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralBoolean{a.(LiteralReal).Value < b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralBoolean{float64(a.(LiteralInteger).Value) < b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralReal).Value < float64(b.(LiteralInteger).Value)}
			}
			CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "lesseq",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't add %s and %s", ae, be))
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
			CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "greater",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)
			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralInteger).Value > b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralBoolean{a.(LiteralReal).Value > b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralBoolean{float64(a.(LiteralInteger).Value) > b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralReal).Value > float64(b.(LiteralInteger).Value)}
			}
			CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "greatereq",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)
			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)
			if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralInteger).Value >= b.(LiteralInteger).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeReal {
				return LiteralBoolean{a.(LiteralReal).Value >= b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeInteger && b.Type() == LiteralTypeReal {
				return LiteralBoolean{float64(a.(LiteralInteger).Value) >= b.(LiteralReal).Value}
			} else if a.Type() == LiteralTypeReal && b.Type() == LiteralTypeInteger {
				return LiteralBoolean{a.(LiteralReal).Value >= float64(b.(LiteralInteger).Value)}
			}
			CreateEvaluateError(fmt.Sprintf("Can't compare %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "isint",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)

			var isElementType = ae.ElementType() == ElementTypeLiteral
			var isValidLiteralType = false

			if isElementType {
				isValidLiteralType = ae.(Literal).Type() == LiteralTypeInteger
			}

			return LiteralBoolean{isValidLiteralType}
		},
	},
	{
		Name: "isreal",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)

			var isElementType = ae.ElementType() == ElementTypeLiteral
			var isValidLiteralType = false

			if isElementType {
				isValidLiteralType = ae.(Literal).Type() == LiteralTypeReal
			}

			return LiteralBoolean{isValidLiteralType}
		},
	},
	{
		Name: "isbool",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)

			var isElementType = ae.ElementType() == ElementTypeLiteral
			var isValidLiteralType = false

			if isElementType {
				isValidLiteralType = ae.(Literal).Type() == LiteralTypeBoolean
			}

			return LiteralBoolean{isValidLiteralType}
		},
	},
	{
		Name: "isnull",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)

			var isElementType = ae.ElementType() == ElementTypeLiteral
			var isValidLiteralType = false

			if isElementType {
				isValidLiteralType = ae.(Literal).Type() == LiteralTypeNull
			}

			return LiteralBoolean{isValidLiteralType}
		},
	},
	{
		Name: "isatom",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			var isElementType = args[0].ElementType() == ElementTypeAtom

			return LiteralBoolean{isElementType}
		},
	},
	{
		Name: "islist",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			var isElementType = args[0].ElementType() == ElementTypeList

			return LiteralBoolean{isElementType}
		},
	},
	{
		Name: "and",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)

			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't use logical operator on %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)

			if a.Type() == LiteralTypeBoolean && b.Type() == LiteralTypeBoolean {
				av := a.(LiteralBoolean).Value
				bv := b.(LiteralBoolean).Value

				return LiteralBoolean{av && bv}
			}
			CreateEvaluateError(fmt.Sprintf("Can't use logical operator on %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "or",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)

			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				panic(fmt.Sprintf("Can't use logical operator on %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)

			if a.Type() == LiteralTypeBoolean && b.Type() == LiteralTypeBoolean {
				av := a.(LiteralBoolean).Value
				bv := b.(LiteralBoolean).Value

				return LiteralBoolean{av || bv}
			}
			panic(fmt.Sprintf("Can't use logical operator on %s and %s", a, b))
		},
	},
	{
		Name: "xor",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)
			be := args[1].Eval(c)

			if ae.ElementType() != ElementTypeLiteral || be.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't use logical operator on %s and %s", ae, be))
			}
			a := ae.(Literal)
			b := be.(Literal)

			if a.Type() == LiteralTypeBoolean && b.Type() == LiteralTypeBoolean {
				av := a.(LiteralBoolean).Value
				bv := b.(LiteralBoolean).Value

				return LiteralBoolean{(av || bv) && !(av && bv)}
			}
			CreateEvaluateError(fmt.Sprintf("Can't use logical operator on %s and %s", a, b))
			return nil
		},
	},
	{
		Name: "not",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			ae := args[0].Eval(c)

			if ae.ElementType() != ElementTypeLiteral {
				CreateEvaluateError(fmt.Sprintf("Can't use logical operator on %s", ae))
			}
			a := ae.(Literal)

			if a.Type() == LiteralTypeBoolean {
				av := a.(LiteralBoolean).Value

				return LiteralBoolean{!av}
			}
			CreateEvaluateError(fmt.Sprintf("Can't use logical operator on %s", a))
			return nil
		},
	},
	{
		Name: "head",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {

			if args[0].ElementType() != ElementTypeList {
				CreateEvaluateError(fmt.Sprintf("Can't use list operators on %s", args[0]))
			}

			av := args[0].(List)

			if len(av.GetElements()) <= 0 {
				return LiteralNull{}
			} else {
				return av.GetElements()[0]
			}
		},
	},
	{
		Name: "tail",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			if args[0].ElementType() != ElementTypeList {
				CreateEvaluateError(fmt.Sprintf("Can't use list operators on %s", args[0]))
			}

			av := args[0].(List)

			if len(av.GetElements()) <= 0 {
				return ListElement{}
			} else {
				return ListElement{av.GetElements()[1:]}
			}
		},
	},
	{
		Name: "cons",
		Args: []Element{Atom{"a"}, Atom{"b"}},
		Code: func(c *Context, args []Element) Element {
			if args[1].ElementType() != ElementTypeList {
				CreateEvaluateError(fmt.Sprintf("Can't use list operators on %s", args[1]))
			}

			bv := args[1].(List)

			return ListElement{append([]Element{args[0]}, bv.GetElements()...)}
		},
	},
	{
		Name: "eval",
		Args: []Element{Atom{"a"}},
		Code: func(c *Context, args []Element) Element {
			return args[0].Eval(c)
		},
	},
}

func GetBuiltinByName(name string) *Builtin {
	for _, b := range Builtins {
		if b.Name == name {
			return &b
		}
	}
	return nil
}

func initBuiltins(c *Context) {
	for _, b := range Builtins {
		c.Add(b.Name, b)
	}
}
