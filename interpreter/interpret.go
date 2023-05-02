package interpreter

import (
	"github.com/flychario/flylang/ast"
)

type Context struct {
	Parent *Context
	Values map[string]ast.Element
}
