package main

import "ast"

import (
	"os"

	"fmt"
)

func main() {
	if len(os.Args) > 1 {
		var (
			element ast.Element = ast.Element(os.Args[1])
		)
		fmt.Println(ast.Match(element, ast.Comment))
		fmt.Println(ast.Match(element, ast.Function))
	}
}