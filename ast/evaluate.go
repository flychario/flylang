package ast

import "fmt"

type Context struct {
	Parent *Context
	Values map[string]Element
}

func GetGlobalContext() *Context {
	c := &Context{Parent: nil}
	c.Values = make(map[string]Element)
	initBuiltins(c)
	return c
}

func NewContext(parent *Context) *Context {
	c := &Context{Parent: parent}
	c.Values = make(map[string]Element)
	return c
}

func (c *Context) Add(name string, value Element) {
	if literal, ok := value.(Literal); ok {
		c.Values[name] = literal
		return
	} else if fun, ok := value.(Lambda); ok {
		c.Values[name] = &fun
		return
	} else if fun, ok := value.(Prog); ok {
		c.Values[name] = &fun
		return
	} else if atom, ok := value.(Builtin); ok {
		c.Values[name] = &atom
		return
	}

	panic(fmt.Sprintf("Can't add value to context %s: %s", name, value.ElementType()))
}

func (c *Context) Get(name string) Element {
	if _, ok := c.Values[name]; ok {
		return c.Values[name]
	} else if c.Parent != nil {
		return c.Parent.Get(name)
	}

	return nil
}

func (a Atom) Eval(c *Context) Element {
	val := c.Get(a.Name)
	if val != nil {
		return val
	}
	panic("undefined variable: " + a.Name)
}

func (l LiteralInteger) Eval(c *Context) Element {
	return l
}

func (l LiteralReal) Eval(c *Context) Element {
	return l
}

func (l LiteralBoolean) Eval(c *Context) Element {
	return l
}

func (l LiteralNull) Eval(c *Context) Element {
	return l
}

func (l ListElement) Eval(c *Context) Element {
	evaluated := make([]Element, len(l.Elements))
	for i, elem := range l.Elements {
		evaluated[i] = elem.Eval(c)
	}

	if len(l.Elements) == 0 {
		return l
	}

	if fun, ok := evaluated[0].(Callable); ok {
		return fun.Call(c, evaluated[1:])
	}
	panic("The first element of a list must be a function")
}

func (f Func) Eval(c *Context) Element {
	c.Add(f.Atom.Name, Lambda{f.List, f.SubProg})
	return f.Atom
} //  `go run main.go file`

func (q Quote) Eval(c *Context) Element {
	return q.Element
}

func (s Setq) Eval(c *Context) Element {
	c.Values[s.Atom.Name] = s.Element.Eval(c)
	return s.Element.Eval(c)
}

func (l Lambda) Eval(c *Context) Element {
	return l
}

func (l Prog) Eval(c *Context) Element {
	return l
}

func (cond Cond) Eval(c *Context) Element {
	body := cond.List.Eval(c)
	if body.ElementType() != ElementTypeLiteral {
		panic("cond body evaluated to non-literal")
	}
	val := body.(Literal)
	if val.Type() != LiteralTypeBoolean {
		panic("cond body evaluated to non-boolean")
	}
	if val.(LiteralBoolean).Value {
		return cond.Element1.Eval(c)
	} else if cond.Element2 != nil {
		return cond.Element2.Eval(c)
	}
	return LiteralNull{}
}

func (w While) Eval(c *Context) Element {
	for true {
		body := w.Element1.Eval(c)
		if body.ElementType() != ElementTypeLiteral {
			panic("cond body evaluated to non-literal")
		}
		val := body.(Literal)
		if val.Type() != LiteralTypeBoolean {
			panic("cond body evaluated to non-boolean")
		}
		if val.(LiteralBoolean).Value {
			w.Element2.Eval(c)
			println("Evaluated!")
		} else {
			println("Not valuated!")
			break
		}
	}
	return LiteralNull{}
}

func (p Program) Eval(c *Context) []Element {
	res := make([]Element, len(p.Elements))
	for i, e := range p.Elements {
		res[i] = e.Eval(c)
	}
	return res
}

func (l Lambda) Call(c *Context, args []Element) (res Element) {
	if len(l.List.GetElements()) > len(args) {
		panic("not enough arguments")
	} else if len(l.List.GetElements()) < len(args) {
		panic("too many arguments")
	}

	newContext := NewContext(c)
	for i, arg := range args {
		newContext.Add(l.List.GetElements()[i].(Atom).Name, arg)
	}

	for _, e := range l.SubProg.Elements {
		res = e.Eval(newContext)

		// TODO: Handle return
	}
	return res
}

func (l Prog) Call(c *Context, args []Element) (res Element) {
	if len(l.List.GetElements()) > len(args) {
		panic("not enough arguments")
	} else if len(l.List.GetElements()) < len(args) {
		panic("too many arguments")
	}

	newContext := NewContext(c)
	for i, arg := range args {
		newContext.Add(l.List.GetElements()[i].(Atom).Name, arg)
	}

	for _, e := range l.SubProg.Elements {
		res = e.Eval(newContext)

		// TODO: Handle return
	}
	return res
}
