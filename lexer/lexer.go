package lexer

import "unicode"

import "slices"

import "strconv"

type Line string

const (
	BRACKET = "/bracket/" // '{', '}', '(', ')', '[', ']'
	KEYWORD = "/keyword/"
	WORD = "/word/"
	SPACE = "/space/"
	SYMBOL = "/symbol/"
)

var (
	Keywords = []string{
		"func",
		"type",
		"var",
	}
	KeywordType = map[string]string{
		"func" : `FUNC`,
		"type" : `TYPE`,
		"var" : `VAR`,
	}
	Brackets = []string{
		"(",
		")",
		"[",
		"]",
		"{",
		"}",
	}
	BracketsType = map[string]string{
		"(" : `PA`,
		")" : `PA_C`,
		"[" : `SQ`,
		"]" : `SQ_C`,
		"{" : `CU`,
		"}" : `CU_C`,
	}
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
	Vrs = []string{
		`"`,
		`'`,
	}
	VrsType = map[string]string{
		`"` : `DOUBLE`,
		`'` : `SINGLE`,
	}
)

type LexObj struct {
	Class string
	Type string
	Value string
}

type WIP[TYPE any] struct {
	Working bool
	Value TYPE
}

func (wip *WIP[any]) Start() {
	wip.Working = true
}

func (wip *WIP[any]) Stop() {
	wip.Working = false
}

func (wip *WIP[any]) Comm() {
	wip.Working = !wip.Working
}

func In[TYPE comparable](obj []TYPE, item TYPE) bool {
	return slices.Contains(obj, item)
}

func Parse(line Line) []LexObj {
	var (
		Final []string
		Outr []LexObj
		WIPString string
		WIPTString WIP[string]
	)

	Clean := func() {
		if WIPString != "" {
			Final = append(Final, WIPString)
			WIPString = ""
		}
	}

	for _, wk := range line {
		/* if unicode.IsLetter(wk) || unicode.IsNumber(wk) {
			WIPString += string(wk)
			continue
		}
		Clean()
		if unicode.IsSpace(wk) {
			Final = append(Final, string(wk))
		}
		if slices.Contains(Symbols, string(wk)) {
			Final = append(Final, string(wk))
		} */
		
		if In(Vrs, string(wk)) {
			WIPTString.Comm()
		}
	}

	for _, rgk := range Final {
		switch {
		case slices.Contains(Keywords, rgk):
			Outr = append(Outr, LexObj{
				Class: KEYWORD,
				Type: KeywordType[rgk],
				Value: rgk,
			})
		case slices.Contains(Brackets, rgk):
			Outr = append(Outr, LexObj{
				Class: BRACKET,
				Type: BracketsType[rgk],
				Value: rgk,
			})
		case func() bool {_, err := strconv.Atoi(rgk); return err == nil}():
			continue
		case rgk == " ":
			Outr = append(Outr, LexObj{
				Class: SPACE,
				Type: `SPACE`,
				Value: rgk,
			})
		default:
			Outr = append(Outr, LexObj{
				Class: WORD,
				Type: `NAME`,
				Value: rgk,
			})
		}
	}

	return Outr
}