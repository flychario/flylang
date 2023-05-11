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
	defer func() {
		if r := recover(); r != nil { // Check if we get element from panic, otherwise it is error message
			if elem, ok := r.(Element); ok {
				if elem.ElementType() == ElementTypeBreak { // If it is not break it can be element from return

				} else {
					panic(r) // panic again with element for return
				}
			} else {
				panic(r) // panic again with error message
			}
		}
	}()

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
			for _, elem := range w.Element2.Elements {
				elem.Eval(c)
			}
		} else {
			break
		}
	}
	return LiteralNull{}
}

func (p Program) Eval(c *Context) (res Element) {
	defer func() {
		if r := recover(); r != nil { // Check if we get element from panic, otherwise it is error message
			if elem, ok := r.(Element); ok {
				res = elem
			} else {
				panic(r) // panic again with error message
			}
		}
	}()

	for _, e := range p.Elements {
		res = e.Eval(c)
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

	defer func() {
		if r := recover(); r != nil { // Check if we get element from panic, otherwise it is error message
			if elem, ok := r.(Element); ok {
				res = elem
			} else {
				panic(r) // panic again with error message
			}
		}
	}()

	for _, e := range l.SubProg.Elements {
		res = e.Eval(newContext)
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

	defer func() {
		if r := recover(); r != nil { // Check if we get element from panic, otherwise it is error message
			if elem, ok := r.(Element); ok {
				res = elem
			} else {
				panic(r) // panic again with error message
			}
		}
	}()

	for _, e := range l.SubProg.Elements {
		res = e.Eval(newContext)

	}
	return res
}

func (r Return) Eval(c *Context) Element {
	var elem = r.Element.Eval(c)
	panic(elem)
}

func (b Break) Eval(c *Context) Element {
	panic(b)
}
