package parser

import (
	"github.com/flychario/flylang/ast"
	"github.com/flychario/flylang/scanner"
	"github.com/flychario/flylang/token"
)

// The Parser structure holds the Parser's internal state.
type Parser struct {
	//file    *token.File
	//errors  scanner.ErrorList
	scanner scanner.Scanner

	// Tracing/debugging
	//mode   Mode // parsing mode
	//trace  bool // == (mode&Trace != 0)
	//indent int  // indentation used for tracing output

	// Comments
	//comments    []*ast.CommentGroup
	//leadComment *ast.CommentGroup // last lead comment
	//lineComment *ast.CommentGroup // last line comment

	// Next token
	//pos token.Pos   // token position
	tok token.Token // one token look-ahead
	lit string      // token literal
}

func (p *Parser) Init(filename string, src []byte) {
	//p.file = fset.AddFile(filename, -1, len(src))
	//var m scanner.Mode
	//if mode&ParseComments != 0 {
	//	m = scanner.ScanComments
	//}
	//eh := func(pos token.Position, msg string) { p.errors.Add(pos, msg) }
	//p.scanner.Init(p.file, src, eh, m)
	p.scanner.Init(src)

	//p.mode = mode
	//p.trace = mode&Trace != 0 // for convenience (p.trace is used frequently)
	p.next()
}

func (p *Parser) next() {
	//p.leadComment = nil
	//p.lineComment = nil
	//prev := p.pos
	p.tok, p.lit = p.scanner.Scan()
}

func (p *Parser) expect(tok token.Token) {
	if p.tok != tok {
		//p.errorExpected(p.pos, tok)
		//p.advance(tok)
		panic("expected " + tok.String())
	}
	p.next()
}

func (p *Parser) parseList() ast.List {
	var elements []ast.Element
	p.expect(token.LPAREN)
	for p.tok != token.RPAREN {
		elements = append(elements, p.parseElement())
	}

	if len(elements) == 1 {
		if ast.IsKeyword(elements[0].ElementType()) {
			return elements[0]
		}
	}

	p.expect(token.RPAREN)
	return ast.ListElement{Elements: elements}
}

func (p *Parser) parseLiteral() (ret ast.Literal) {
	if p.tok == token.INTEGER {
		ret = ast.Literal{Value: p.lit, Type: ast.LiteralTypeInteger}
		p.expect(token.INTEGER)
		return ret
	} else if p.tok == token.REAL {
		ret = ast.Literal{Value: p.lit, Type: ast.LiteralTypeReal}
		p.expect(token.REAL)
		return ret
	} else if p.tok == token.BOOLEAN {
		ret = ast.Literal{Value: p.lit, Type: ast.LiteralTypeBoolean}
		p.expect(token.BOOLEAN)
		return ret
	} else if p.tok == token.NULL {
		ret = ast.Literal{Value: p.lit, Type: ast.LiteralTypeNull}
		p.expect(token.NULL)
		return ret
	}
	panic("expected literal")
}

func (p *Parser) parseAtom() ast.Atom {
	ret := ast.Atom{Name: p.lit}
	p.expect(token.IDENTIFIER)
	return ret
}

func (p *Parser) parseSetq() ast.Setq {
	p.expect(token.SETQ)
	return ast.Setq{Atom: p.parseAtom(), Element: p.parseElement()}
}

func (p *Parser) parseFunc() ast.Func {
	p.expect(token.FUNC)
	return ast.Func{Atom: p.parseAtom(), List: p.parseList(), Element: p.parseElement()}
}

func (p *Parser) parseLambda() ast.Lambda {
	p.expect(token.LAMBDA)
	return ast.Lambda{List: p.parseList(), Element: p.parseElement()}
}

func (p *Parser) parseProg() ast.Prog {
	p.expect(token.PROG)
	return ast.Prog{List: p.parseList(), Element: p.parseElement()}
}

func (p *Parser) parseCond() ast.Cond {
	p.expect(token.COND)
	list := p.parseList()
	element1 := p.parseElement()
	if p.tok == token.RPAREN {
		return ast.Cond{List: list, Element1: element1, Element2: nil}
	}
	element2 := p.parseElement()
	return ast.Cond{List: list, Element1: element1, Element2: element2}
}

func (p *Parser) parseWhile() ast.While {
	p.expect(token.WHILE)
	return ast.While{Element1: p.parseElement(), Element2: p.parseElement()}
}

func (p *Parser) parseReturn() ast.Return {
	p.expect(token.RETURN)
	return ast.Return{Element: p.parseElement()}
}

func (p *Parser) parseBreak() ast.Break {
	p.expect(token.BREAK)
	return ast.Break{}
}

func (p *Parser) parseElement() ast.Element {
	switch p.tok {
	case token.IDENTIFIER:
		return p.parseAtom()
	case token.INTEGER, token.REAL, token.BOOLEAN, token.NULL:
		return p.parseLiteral()
	case token.LPAREN:
		return p.parseList()
	case token.SETQ:
		return p.parseSetq()
	case token.FUNC:
		return p.parseFunc()
	case token.LAMBDA:
		return p.parseLambda()
	case token.PROG:
		return p.parseProg()
	case token.COND:
		return p.parseCond()
	case token.WHILE:
		return p.parseWhile()
	case token.RETURN:
		return p.parseReturn()
	case token.BREAK:
		return p.parseBreak()
	}
	panic("expected element")
}

func (p *Parser) ParseProgram() ast.Program {
	var elements []ast.Element
	for p.tok != token.EOF {
		elements = append(elements, p.parseElement())
	}
	return ast.Program{Elements: elements}
}
