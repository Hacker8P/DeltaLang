module delta

go 1.24.5

replace ast => ./ast

replace lexer => ./lexer

replace logm => ../goinit/lmd

require lexer v0.0.0-00010101000000-000000000000

require (
	github.com/fatih/color v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
	logm v0.0.0-00010101000000-000000000000 // direct
)
