package main

import (
	"fmt"
	"github.com/flychario/flylang/ast"
	"github.com/flychario/flylang/parser"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestSamples(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf(fmt.Sprintf("%v", r))
		}
	}()

	for _, sample := range []struct {
		programFile string
		evalResult  ast.Element
	}{
		{"samples/break.fly", ast.LiteralInteger{Value: 3}},
		{"samples/break_return.fly", ast.LiteralInteger{Value: 3}},
		{"samples/eval.fly", ast.LiteralInteger{Value: 15}},
		{"samples/lists.fly", ast.LiteralInteger{Value: 2}},
		{"samples/logical_operators.fly", ast.LiteralBoolean{true}},
		{"samples/quote.fly", ast.LiteralInteger{11}},
		{"samples/return.fly", ast.LiteralInteger{5}},
		{"samples/while.fly", ast.LiteralInteger{0}},

		{"tests/fib.fly", ast.LiteralInteger{55}},
		{"tests/lambda.fly", ast.LiteralInteger{-3}},
		{"tests/logical-operators.fly", ast.LiteralBoolean{false}},
		{"tests/lists.fly", ast.ListElement{Elements: []ast.Element{ast.LiteralInteger{1}}}},
	} {
		elem := runProgram(sample.programFile)
		if !reflect.DeepEqual(elem, sample.evalResult) {
			t.Errorf("sample %s: expected %v, got %v", sample.programFile, sample.evalResult, elem)
		}
	}
}

func runProgram(programFile string) ast.Element {
	file, err := os.Open(programFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read bytes from file
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var p parser.Parser
	p.Init(programFile, content)
	program := p.ParseProgram()

	c := ast.GetGlobalContext()
	return program.Eval(c)
}
