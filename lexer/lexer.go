package lexer

/*
-------------------------------------------------------
	THIS IS THE LEXER. RETURNS TOKENS FROM THE CODE
-------------------------------------------------------
*/

import "unicode"

import "slices"

import "logm"

import "os"

var log = lmd.Logger{
	File: os.Stdout,
	FileErr: os.Stderr,
}

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
	PUNCT = "/punct/"
	OPERATOR = "/operator/"
	CONFRONT = "/confront"
	NAME = "/name/"
)

// END

/*
CLASSES' OBJECTS AND NAMES
*/

var (
	Keyword = map[string]string{
		"func" : `FUNC`,
		"type" : `TYPE`,
		"var" : `VAR`,
		"if" : `IF`,
		"else": `ELSE`,
		"switch": `SWITCH`,
		"return": `RETURN`,
		"pragma": `PRAGMA`,
		"define": `DEFINE`,
		"typedef": `TYPEDEF`,
		"import": `IMPORT`,
		"include": `INCLUDE`,
		"const": `CONST`,
		"continue": `CONTINUE`,
		"goto": `GOTO`,
		"break": `BREAK`,
		"for": `FOR`,
		"while": `WHILE`,
		"when": `WHEN`,
		"thread": `THREAD`,
	}
	Brackets = map[string]string{
		"(" : `PA`,
		")" : `PA_C`,
		"[" : `SQ`,
		"]" : `SQ_C`,
		"{" : `CU`,
		"}" : `CU_C`,
	}
	Punct = map[string]string{
		"." : `DOT`,
		":" : `COLON`,
		";" : `SEMICOLON`,
	}
	Operators = map[string]string{
		"+" : `ADD`,
		"-" : `SUB`,
		"/" : `SPLIT`,
		"*" : `MULT`,
		"=" : `ASSIGN`,
		"++": `INCREMENT`,
		"--": `DECREMENT`,
		"+=": `ADDTO`,
		"-=": `SUBTO`,
		"/=": `SPLITTO`,
		"*=": `MULTTO`,
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
	Confront = map[string]string{
		"==": `EQUAL`,
		"!=": `NOT EQUAL`,
		"||": `OR`,
		"&&": `AND`,
		">": `ADDC`,
		"<": `MINC`,
		">=": `ADDEQUALC`,
		"<=": `MINEQUALC`,
	}
	AbsTypes = map[string]string{
		"STRING": "STRING",
		"RAWSTRING": "RAWSTRING",
		"NUMBER": "NUMBER",
	}
)

type DT string

type ERROR error

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

type Intm struct {
	Type DT
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
	Intm Intm
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

/* var (
	LexSPACE = LexObj{
		Class: SPACE,
		Type: `SPACE`,
	}
	LexSTRING = LexObj{
		Class: STRING
	}
	LexNUM
	LexKEYWORD
	LexWORD
) */

func Check[T comparable, TT any](a1f map[T]TT, a2s T) bool {
	_, err := a1f[a2s]

	return err
}

func CheckAll(a2s string) map[string]map[string]string {
	if Check(Keyword, a2s) {return map[string]map[string]string{
		"Type": map[string]string{"." : KEYWORD},
		"Obj": Keyword,
	}}

	if Check(Operators, a2s) {return map[string]map[string]string{
		"Type": map[string]string{"." : OPERATOR},
		"Obj": Operators,
	}}

	if Check(Confront, a2s) {return map[string]map[string]string{
		"Type": map[string]string{"." : CONFRONT},
		"Obj": Confront,
	}}

	if Check(Punct, a2s) {return map[string]map[string]string{
		"Type": map[string]string{"." : PUNCT},
		"Obj": Punct,
	}}

	if Check(Brackets, a2s) {return map[string]map[string]string{
		"Type": map[string]string{"." : BRACKET},
		"Obj": Brackets,
	}}

	return map[string]map[string]string{
		"Type": map[string]string{"." : NAME},
		"Obj": nil,
	}
}

func DictToSlice[TYPE_I comparable, TYPE any](Dict map[TYPE_I]TYPE) []TYPE_I {
	var Final []TYPE_I

	for kyks := range Dict {
		Final = append(Final, kyks)
	}

	return Final
}

func Parse(line Line) []LexObj {
	var (
		Final []Intm
		Outr []LexObj
		// WIPString string
		WIPString WIP[string] = WIP[string]{
			Intm: Intm{
				Type: DT("$"),
			},
		}
		WIPNumber WIP[string] = WIP[string]{
			Intm: Intm{
				Type: DT("£"),
			},
		}
		WIPKeyword WIP[string] = WIP[string]{
			Intm: Intm{
				Type: DT("%"),
			},
		}
	)

	Clean := func(obj *WIP[string]) {
		if obj.Value != "" {
			obj.Intm.Value = obj.Value
			Final = append(Final, obj.Intm)
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
				WIPString.Vr = string(wk)
				WIPString.Start()
			} else if WIPString.Vr == string(wk) {
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
		If the rune is a letter, a number or a symbol and WIPString is working, push it into them.
		*/

		if (unicode.IsLetter(wk) || unicode.IsNumber(wk) || unicode.IsPunct(wk)) && WIPString.Working {
			WIPString.Value += string(wk)
			continue
		}

		/*
		If the rune is a symbol and WIPString isn't working
		*/

		if unicode.IsPunct(wk) && !WIPString.Working {
			WIPKeyword.Stop()
			Clean(&WIPKeyword)
			Final = append(Final, Intm{
				Type: DT("!"),
				Value: string(wk),
			})
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

		/* Final = append(Final, Intm{
			Type: DT("!"),
			Value: string(wk),
		}) */

		WIPNumber.Stop()
		Clean(&WIPNumber)
		WIPKeyword.Stop()
		Clean(&WIPKeyword)

		// END
	}

	/*
	log.LogInline(false, Final)
	*/

	for _, rgk := range Final {
		/*
		log.Log(false, DictToSlice(Keyword))
		*/
		switch {
		case rgk.Type == "%":
			var Type = CheckAll(rgk.Value)

			if Type["Type"]["."] == NAME {
				Outr = append(Outr, LexObj{
					Class: Type["Type"]["."],
					Type: "NAME",
					Value: rgk.Value,
				})
			} else {
				Outr = append(Outr, LexObj{
					Class: Type["Type"]["."],
					Type: Type["Obj"][rgk.Value],
					Value: "NONE",
				})
			}
		case rgk.Type == "!" && slices.Contains(DictToSlice(Brackets), rgk.Value):
			Outr = append(Outr, LexObj{
				Class: BRACKET,
				Type: Brackets[rgk.Value],
				Value: rgk.Value,
			})
		case rgk.Type == "$":
			Outr = append(Outr, LexObj{
				Class: WORD,
				Type: `STRING`,
				Value: rgk.Value,
			})
		case rgk.Type == "£":
			Outr = append(Outr, LexObj{
				Class: WORD,
				Type: `INT`,
				Value: rgk.Value,
			})
		default:
			Outr = append(Outr, LexObj{
				Class: WORD,
				Type: `ANY`,
				Value: rgk.Value,
			})
		}
	}

	return Outr
}

// END