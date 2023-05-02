package interpreter

import "github.com/flychario/flylang/ast"

type Element interface {
	Eval(*Context) ast.Element
}

//func (a *ast.Atom) Eval(c *Context) Element {
//	if c.Values[a.Name] != nil {
//		return c.Values[a.Name]
//	}
//	if c.Parent != nil {
//		return c.Parent.Values[a.Name]
//	}
//	panic("undefined variable: " + a.Name)
//}
