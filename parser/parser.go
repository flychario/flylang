package parser

import (
	"fmt"
	"github.com/flychario/flylang/ast"
	"github.com/flychario/flylang/scanner"
	"github.com/flychario/flylang/token"
	"strconv"
)

// The Parser structure holds the Parser's internal state.
type Parser struct {
	scanner scanner.Scanner

	// Next token
	pos token.Position // token position
	tok token.Token    // one token look-ahead
	lit string         // token literal
}

func (p *Parser) Init(filename string, src []byte) {
	p.scanner.Init(src)

	p.next()
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}

func (p *Parser) expect(tok token.Token) {
	if p.tok != tok {
		p.ThrowError("expected " + tok.String())
	}
	p.next()
}

func (p *Parser) parseList() ast.List {
	var elements []ast.Element
	p.expect(token.LPAREN)
	defer p.expect(token.RPAREN)

	switch p.tok {
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
	case token.QUOTE:
		return p.parseQuote()
	case token.WHILE:
		return p.parseWhile()
	case token.RETURN:
		return p.parseReturn()
	case token.BREAK:
		return p.parseBreak()
	}

	for p.tok != token.RPAREN {
		elements = append(elements, p.parseElement())
	}

	return ast.ListElement{Elements: elements}
}

func (p *Parser) parseLiteral() (ret ast.Literal) {
	if p.tok == token.INTEGER {
		val, err := strconv.ParseInt(p.lit, 10, 64)
		if err != nil {
			p.ThrowError(err.Error())
		}
		ret = ast.LiteralInteger{Value: val}
		p.expect(token.INTEGER)
		return ret
	} else if p.tok == token.REAL {
		val, err := strconv.ParseFloat(p.lit, 64)
		if err != nil {
			p.ThrowError(err.Error())
		}
		ret = ast.LiteralReal{Value: val}
		p.expect(token.REAL)
		return ret
	} else if p.tok == token.BOOLEAN {
		if p.lit == "true" {
			ret = ast.LiteralBoolean{Value: true}
		} else if p.lit == "false" {
			ret = ast.LiteralBoolean{Value: false}
		} else {
			p.ThrowError("invalid boolean literal")
		}
		p.expect(token.BOOLEAN)
		return ret
	} else if p.tok == token.NULL {
		ret = ast.LiteralNull{}
		p.expect(token.NULL)
		return ret
	}
	p.ThrowError("expected literal")
	return
}

func (p *Parser) parseAtom() ast.Atom {
	ret := ast.Atom{Name: p.lit}
	p.expect(token.IDENTIFIER)
	return ret
}

func (p *Parser) parseShortQuote() ast.Quote {
	p.expect(token.SHORT_QUOTE)
	return ast.Quote{Element: p.parseElement()}
}

func (p *Parser) parseQuote() ast.Quote {
	p.expect(token.QUOTE)
	return ast.Quote{Element: p.parseElement()}
}

func (p *Parser) parseSetq() ast.Setq {
	p.expect(token.SETQ)
	return ast.Setq{Atom: p.parseAtom(), Element: p.parseElement()}
}

func (p *Parser) parseFunc() ast.Func {
	p.expect(token.FUNC)
	return ast.Func{Atom: p.parseAtom(), List: p.parseList(), SubProg: p.ParseSubProgram()}
}

func (p *Parser) parseLambda() ast.Lambda {
	p.expect(token.LAMBDA)
	return ast.Lambda{List: p.parseList(), SubProg: p.ParseSubProgram()}
}

func (p *Parser) parseProg() ast.Prog {
	p.expect(token.PROG)
	return ast.Prog{List: p.parseList(), SubProg: p.ParseSubProgram()}
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
	return ast.While{Element1: p.parseElement(), Element2: p.ParseSubProgram()}
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
	case token.SHORT_QUOTE:
		return p.parseShortQuote()
	case token.LPAREN:
		return p.parseList()
	}

	p.ThrowError("expected element")
	return nil
}

func (p *Parser) ParseSubProgram() ast.Program {
	var elements []ast.Element
	for p.tok != token.RPAREN {
		elements = append(elements, p.parseElement())
	}
	return ast.Program{Elements: elements}
}

func (p *Parser) ParseProgram() ast.Program {
	var elements []ast.Element
	for p.tok != token.EOF {
		elements = append(elements, p.parseElement())
	}
	return ast.Program{Elements: elements}
}

func (p *Parser) ThrowError(msg string) {
	panic(fmt.Sprintf("Syntax error: %s %d:%d", msg, p.pos.Line, p.pos.Column))
}
