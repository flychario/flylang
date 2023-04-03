package token

/*
digit         [0-9]
letter        [a-zA-Z]
boolean       "true"|"false"
integer       {digit}+
real          {integer} . {integer}
identifier    {letter}({letter}|{digit})*
literal       {boolean} | {integer} | {real}
atom          {identifier}
element       {literal} | {atom} | {element}
list          ( {element} {element}* )
program       {element} {element}*

%%
"("   { printf("LPAREN "); }
")"   { printf("RPAREN "); }
"quote" { printf("QUOTE "); }
"'("  { printf("QUOTE "); }
"setq" { printf("SETQ "); }
"func" { printf("FUNC "); }
"lambda" { printf("LAMBDA "); }
"prog" { printf("PROG "); }
"cond" { printf("COND "); }
"while" { printf("WHILE "); }
"return" { printf("RETURN "); }
"break" { printf("BREAK "); }
{identifier} { printf("IDENTIFIER "); }
{literal} { printf("LITERAL "); }
"\n" { printf("\n"); }
%%
*/

type Token int

const (
	ILLEGAL Token = iota
	EOF

	IDENTIFIER
	INTEGER
	REAL
	BOOLEAN
	NULL

	LPAREN
	RPAREN
	QUOTE
	PLUS
	MINUS

	SETQ
	FUNC
	LAMBDA
	PROG
	COND
	WHILE
	RETURN
	BREAK
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENTIFIER: "IDENTIFIER",
	INTEGER:    "INTEGER",
	REAL:       "REAL",
	BOOLEAN:    "BOOLEAN",

	LPAREN: "(",
	RPAREN: ")",
	QUOTE:  "'",
	PLUS:   "+",
	MINUS:  "-",

	SETQ:   "setq",
	FUNC:   "func",
	LAMBDA: "lambda",
	PROG:   "prog",
	COND:   "cond",
	WHILE:  "while",
	RETURN: "return",
	BREAK:  "break",
}

var keywords = map[string]Token{
	"setq":   SETQ,
	"func":   FUNC,
	"lambda": LAMBDA,
	"prog":   PROG,
	"cond":   COND,
	"while":  WHILE,
	"return": RETURN,
	"break":  BREAK,
	"quote":  QUOTE,
}

func (tok Token) String() string {
	return tokens[tok]
}

func isBoolean(ident string) bool {
	return ident == "true" || ident == "false"
}

func isNull(ident string) bool {
	return ident == "null"
}

func Lookup(ident string) Token {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	if isBoolean(ident) {
		return BOOLEAN
	}
	if isNull(ident) {
		return NULL
	}
	return IDENTIFIER
}

type Position struct {
	//Filename string // filename, if any
	Offset int // offset, starting at 0
	Line   int // line number, starting at 1
	Column int // column number, starting at 1 (byte count)
}
