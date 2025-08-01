package lexer

/*
-------------------------------------------------------
	THIS IS THE LEXER. RETURNS TOKENS FROM THE CODE
-------------------------------------------------------
*/

import "unicode"

import "slices"

import "strconv"

type Line string

/*
CLASSES
*/

const (
	BRACKET = "/bracket/" // '{', '}', '(', ')', '[', ']'
	KEYWORD = "/keyword/"
	WORD = "/word/"
	SPACE = "/space/"
	SYMBOL = "/symbol/"
	TYPE = "/type/"
)

// END

/*
CLASSES' OBJECTS AND NAMES
*/

var (
	Keywords = []string{
		"func",
		"type",
		"var",
		"class",
		"struct",
		"string",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"float",
		"float8",
		"float16",
		"float32",
		"float64",
		"map",
	}
	KeywordType = map[string]string{
		"func" : `FUNC`,
		"type" : `TYPE`,
		"var" : `VAR`,
		"class" : `CLASS`,
		"struct" : `STRUCT`,
		"string" : `STRING`,
		"int" : `INT`,
		"int8" : `INT8`,
		"int16" : `INT16`,
		"int32" : `INT32`,
		"int64" : `INT64`,
		"float" : `FLOAT`,
		"float8" : `FLOAT8`,
		"float16" : `FLOAT16`,
		"float32" : `FLOAT32`,
		"float64" : `FLOAT64`,
		"map" : `MAP`,
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
		"`",
	}
	VrsType = map[string]string{
		`"` : `DOUBLE`,
		`'` : `SINGLE`,
		"`" : `INVERSE`,
	}
	Types = []string{
		`STRING`,
		`NUMBER`,
	}
)

// END

/*
LexObj is the token generated from the lexer.
Containes:
	- Class
	- Type
	- Value
*/

type LexObj struct {
	Class string
	Type string
	Value string
}

// END

/*
WIP struct
*/

type WIP[TYPE any] struct {
	Working bool
	Value TYPE
	Vr string
	Del string
	Oth []string
}

/*
Methods
*/

func (wip *WIP[any]) Start() {
	wip.Working = true
}

func (wip *WIP[any]) Stop() {
	wip.Working = false
}

func (wip *WIP[any]) Comm() {
	wip.Working = !wip.Working
}

// END

/*
In util.
A more powerful alias for slices.Contains.
*/

func In[TYPE comparable](obj []TYPE, item TYPE) bool {
	return slices.Contains(obj, item)
}

// END

/*
PARSER
*/

func Parse(line Line) []LexObj {
	var (
		Final []string
		Outr []LexObj
		// WIPString string
		WIPString WIP[string] = WIP[string]{
			Del: "$",
		}
		WIPNumber WIP[string] = WIP[string]{
			Del: "£",
		}
		WIPKeyword WIP[string] = WIP[string]{
			Del: "%",
		}
	)

	Clean := func(obj *WIP[string]) {
		if obj.Value != "" {
			Final = append(Final, obj.Value) // obj.Del + obj.Value)
			obj.Value = ""
		}
	}

	for _, wk := range line {
		/*
		First, if the rune is empty, continue.
		*/

		if string(wk) == "" {
			continue
		}

		/*
		If the rune isn't empty, check if it's a string delimiter.
		*/

		if In(Vrs, string(wk)) {

			/*
			If WIPString isn't running, start it.
			Else, stop it and run `Clean(WIPString)` to Clean it.
			Then continue.
			*/

			if !WIPString.Working {
				WIPString.Start()
			} else {
				WIPString.Stop()
				Clean(&WIPString)
			}
			continue
		}

		/*
		If the rune isn't a string delimiter and isn't empty, check if it's a number or a character and push it into the string if it's running.
		If no one of the precedent condition passed, check if the rune is a number but the string isn't running.
		Then, in that case, start WIPNumber and push the rune into it, after, continue.
		Finally, if no one of the precedent called a continue, stop WIPNumber and Clean it.
		*/

		/*
		If the rune is a letter or a number and WIPString is working, push it into them.
		*/

		if (unicode.IsLetter(wk) || unicode.IsNumber(wk)) && WIPString.Working {
			WIPString.Value += string(wk)
			continue
		}

		/*
		Else, if the rune is a number but the keyword isn't working, add it to WIPNumber
		*/

		if unicode.IsNumber(wk) && !WIPKeyword.Working {
			WIPNumber.Start()
			WIPNumber.Value += string(wk)
			continue
		}

		/*
		Finally, if 
		*/

		if unicode.IsLetter(wk) {
			WIPKeyword.Start()
		}

		if (unicode.IsLetter(wk) || unicode.IsNumber(wk)) {
			WIPKeyword.Value += string(wk)
			continue
		}

		WIPNumber.Stop()
		Clean(&WIPNumber)
		WIPKeyword.Stop()
		Clean(&WIPKeyword)

		// END
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

// END