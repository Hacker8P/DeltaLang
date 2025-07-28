package ast

import (
	"regexp"
)

type Element string

type Syntax = *regexp.Regexp

func MkSyntax(compile string) Syntax {
	resk, _ := regexp.Compile(compile)

	return resk
}

var (
	Comment Syntax = MkSyntax(`^//.*$`)
	Function Syntax = MkSyntax(`^(\s*func\s+)?[\p{L}_][\p{L}\p{N}_]*\s*\(\)\s*\{.*\}$`)
)

func Match(element Element, syntax Syntax) bool {
	return syntax.Match([]byte(element))
}