package main

import (
	"go/ast"
	"go/parser"
	"log"
)

func main() {
	expr, err := parser.ParseExpr("v + 1")
	if err != nil {
		log.Fatal("Error:", err)
	}

	ast.Print(nil, expr)
}
