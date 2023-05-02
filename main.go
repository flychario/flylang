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
	// get file name from args
	fileName := os.Args[1]

	// open file
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

	//var s scanner.Scanner
	//s.Init(content)
	//for {
	//	tok, lit := s.Scan()
	//	if tok == token.EOF {
	//		break
	//	}
	//	println(tok.String(), lit)
	//}

	var p parser.Parser
	p.Init(fileName, content)
	program := p.ParseProgram()
	//fmt.Printf("%#v\n\n", program)

	c := ast.GetGlobalContext()
	execRes := program.Eval(c)
	fmt.Printf("\n%v", execRes[len(execRes)-1])
}
