package main

import "ast"

import "lexer"

import (
	"os"

	"fmt"
)

func main() {
	if len(os.Args) > 2 {
		var (
			element ast.Element = ast.Element(os.Args[1])
		)
		fmt.Println(ast.Match(element, ast.Comment))
		fmt.Println(ast.Match(element, ast.FunctionInline))
		for _, gwxksk := range lexer.Parse(lexer.Line(os.Args[2])) {
			fmt.Println(gwxksk)
		}
	}
}