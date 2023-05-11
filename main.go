package main

import (
	"fmt"
	"github.com/flychario/flylang/ast"

	//"github.com/flychario/flylang/scanner"
	//"github.com/flychario/flylang/token"
	"github.com/flychario/flylang/parser"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	runRes := run(fileName)
	fmt.Println(runRes)
}

func run(fileName string) string {
	file, err := os.Open(fileName)
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
	p.Init(fileName, content)
	execRes := execWithErrorHandling(p)

	return fmt.Sprintf("\n%v", execRes)
}

func execWithErrorHandling(parser parser.Parser) ast.Element {
	defer func() {
		if r := recover(); r != nil { // Check if we get element from panic, otherwise it is error message
			fmt.Printf("%v", r)
		}
	}()
	program := parser.ParseProgram()

	c := ast.GetGlobalContext()
	return program.Eval(c)
}
