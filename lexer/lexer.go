package lexer

import "unicode"

import "slices"

type Line string

var (
	Symbols = []string{
		`;`,
		`.`,
		`:`,
		`(`,
		`)`,
		`{`,
		`}`,
		`"`,
		`'`,
	}
)

const (
	BRACKET = "/bracket/"
	KEYWORD = "/keyword/"
	WORD = "/word/"
)

type LexObj struct {
	Class string
	Type string
	Value string
}

func Parse(line Line) []string {
	var (
		Final []string
		WIPString string
	)

	Clean := func() {
		if WIPString != "" {
			Final = append(Final, WIPString)
			WIPString = ""
		}
	}

	for _, wk := range line {
		if unicode.IsLetter(wk) || unicode.IsNumber(wk) {
			WIPString += string(wk)
			continue
		}
		Clean()
		if unicode.IsSpace(wk) {
			Final = append(Final, string(wk))
		}
		if slices.Contains(Symbols, string(wk)) {
			Final = append(Final, string(wk))
		}
	}

	return Final
}