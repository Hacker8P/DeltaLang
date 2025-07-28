package lexer

import "unicode"

type Line string

func Parse(line Line) []string {
	var (
		Final []string
		WIPString string
	)

	for rk, wk := range line {
		if unicode.IsLetter(wk) {
			WIPString += string(wk)
		} else {
			Final = append(Final, WIPString)
			WIPString = ""
		}
	}
}