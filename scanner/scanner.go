package scanner

import (
	"github.com/flychario/flylang/token"
	"unicode/utf8"
)

type Scanner struct {
	src    []byte // source
	ch     rune   // current character
	offset int    // current offset
	prev   rune   // previous character

	lineOffset int // offset of current line
	line       int // current line
}

func (s *Scanner) Init(src []byte) {
	s.src = src
	s.ch = ' '
	s.offset = 0
	s.lineOffset = 0
	s.line = 1
	s.next()
}

func (s *Scanner) next() {
	if s.offset >= len(s.src) {
		s.prev = s.ch
		s.ch = -1 // eof
	} else {
		r, w := rune(s.src[s.offset]), 1
		switch {
		case r == 0:
			panic("illegal character NUL")
		case r >= utf8.RuneSelf:
			r, w = utf8.DecodeRune(s.src[s.offset:])
			if r == utf8.RuneError && w == 1 {
				panic("illegal UTF-8 encoding")
			}
		}
		s.offset += w
		s.lineOffset += 1
		if r == '\n' {
			s.line++
			s.lineOffset = 0
		}

		s.prev = s.ch
		s.ch = r
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (s *Scanner) scanIdentifier() string {
	var buf [64]byte
	n := 0
	for isLetter(s.ch) || isDigit(s.ch) {
		if n < len(buf) {
			buf[n] = byte(s.ch)
		}
		n++
		s.next()
	}
	return string(buf[0:n])
}

func (s *Scanner) scanNumber() (token.Token, string) {
	var buf [64]byte
	i := 0
	tok := token.INTEGER
	for isDigit(s.ch) || s.ch == '.' {
		if s.ch == '.' && tok == token.INTEGER {
			tok = token.REAL
		} else if s.ch == '.' && tok == token.REAL {
			tok = token.ILLEGAL
		}
		if i < len(buf) {
			buf[i] = byte(s.ch)
		}
		i++
		s.next()
	}

	if isLetter(s.ch) || s.prev == '.' {
		tok = token.ILLEGAL
	}

	return tok, string(buf[0:i])
}

func (s *Scanner) Scan() (pos token.Position, tok token.Token, lit string) {
	// skip white space
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.next()
	}

	// current position
	pos = token.Position{
		Offset: s.offset,
		Line:   s.line,
		Column: s.lineOffset,
	}

	// identifier or keyword
	if isLetter(s.ch) {
		lit = s.scanIdentifier()
		tok = token.Lookup(lit)

		return
	}

	// number
	if isDigit(s.ch) {
		tok, lit = s.scanNumber()
		return
	}

	// special character
	switch s.ch {
	case -1:
		tok = token.EOF
	case '(':
		tok = token.LPAREN
	case ')':
		tok = token.RPAREN
	case '\'':
		tok = token.SHORT_QUOTE
	case '+':
		s.next()
		tok, lit = s.scanNumber()
		return
	case '-':
		s.next()
		tok, lit = s.scanNumber()
		lit = "-" + lit
		return
	default:
		tok = token.ILLEGAL
	}
	s.next()

	return
}
