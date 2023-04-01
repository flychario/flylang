package parser

import (
	"fmt"
	"testing"
)

func TestBoolean(t *testing.T) {
	for _, test := range []struct {
		input string
		//ast   ast.Program
	}{
		{"(true false)"},
	} {
		var p Parser
		p.Init("test", []byte(test.input))
		res := p.ParseProgram()
		fmt.Printf("%#v\n", res)
		t.Log(res)
	}
}
