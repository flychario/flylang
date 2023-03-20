package scanner

import (
	"github.com/flychario/flylang/token"
	"testing"
)

func TestNumbers(t *testing.T) {
	for _, test := range []struct {
		input string
		tok   token.Token
		lit   string
	}{
		{"0", token.INTEGER, "0"},
		{"1", token.INTEGER, "1"},
		{"123", token.INTEGER, "123"},
		{"0.0", token.REAL, "0.0"},
		{"1.0", token.REAL, "1.0"},
		{"1.1", token.REAL, "1.1"},
		{"1.1.1", token.ILLEGAL, "1.1.1"},
	} {
		var s Scanner
		s.Init([]byte(test.input))
		tok, lit := s.Scan()
		if tok != test.tok {
			t.Errorf("expected token %s, got %s", test.tok, tok)
		}
		if lit != test.lit {
			t.Errorf("expected literal %s, got %s", test.lit, lit)
		}
	}
}

func TestIdentifiers(t *testing.T) {
	for _, test := range []struct {
		input string
		tok   token.Token
		lit   string
	}{
		{"a", token.IDENTIFIER, "a"},
		{"abc", token.IDENTIFIER, "abc"},
		{"a1", token.IDENTIFIER, "a1"},
		{"a1b2c3", token.IDENTIFIER, "a1b2c3"},

		{"setq", token.SETQ, "setq"},
		{"func", token.FUNC, "func"},
		{"lambda", token.LAMBDA, "lambda"},
		{"prog", token.PROG, "prog"},
		{"cond", token.COND, "cond"},
		{"while", token.WHILE, "while"},
		{"return", token.RETURN, "return"},
		{"break", token.BREAK, "break"},
	} {
		var s Scanner
		s.Init([]byte(test.input))
		tok, _ := s.Scan()
		if tok != test.tok {
			t.Errorf("expected token %s, got %s", test.tok, tok)
		}
	}
}

func TestProgram(t *testing.T) {
	src := []byte(`
(setq a 2)
(lambda (x) (plus x 1))
(prog (a 1) (b 2))
(cond ((greater a 0) 2 true))
(return 1)
	`)
	want := []struct {
		tok token.Token
		lit string
	}{
		{token.LPAREN, ""},
		{token.SETQ, "setq"},
		{token.IDENTIFIER, "a"},
		{token.INTEGER, "2"},
		{token.RPAREN, ""},

		{token.LPAREN, ""},
		{token.LAMBDA, "lambda"},
		{token.LPAREN, ""},
		{token.IDENTIFIER, "x"},
		{token.RPAREN, ""},
		{token.LPAREN, ""},
		{token.IDENTIFIER, "plus"},
		{token.IDENTIFIER, "x"},
		{token.INTEGER, "1"},
		{token.RPAREN, ""},
		{token.RPAREN, ""},

		{token.LPAREN, ""},
		{token.PROG, "prog"},
		{token.LPAREN, ""},
		{token.IDENTIFIER, "a"},
		{token.INTEGER, "1"},
		{token.RPAREN, ""},
		{token.LPAREN, ""},
		{token.IDENTIFIER, "b"},
		{token.INTEGER, "2"},
		{token.RPAREN, ""},
		{token.RPAREN, ""},

		{token.LPAREN, ""},
		{token.COND, "cond"},
		{token.LPAREN, ""},
		{token.LPAREN, ""},
		{token.IDENTIFIER, "greater"},
		{token.IDENTIFIER, "a"},
		{token.INTEGER, "0"},
		{token.RPAREN, ""},
		{token.INTEGER, "2"},
		{token.IDENTIFIER, "true"},
		{token.RPAREN, ""},
		{token.RPAREN, ""},

		{token.LPAREN, ""},
		{token.RETURN, "return"},
		{token.INTEGER, "1"},
		{token.RPAREN, ""},
	}
	var s Scanner
	s.Init(src)
	for _, w := range want {
		tok, lit := s.Scan()
		if tok != w.tok {
			t.Errorf("expected token %s, got %s", w.tok, tok)
		}
		if lit != w.lit {
			t.Errorf("expected literal %s, got %s", w.lit, lit)
		}
	}

}
