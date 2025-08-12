package main

import "lexer"

import (
	"os"

	"fmt"

	"logm"
)

var log = lmd.Logger{
	File: os.Stdout,
	FileErr: os.Stderr,
}

func main() {
	if len(os.Args) > 1 {
		content, err := os.ReadFile(os.Args[1])
		log.ErrLog(err)
		cont := string(content)
		fmt.Println(`---- LEXER ----`)
		fmt.Println(lexer.Parse(lexer.Line(cont)))
	}
}