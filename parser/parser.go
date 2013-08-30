
//line parser.l:2

package parser
import __yyfmt__ "fmt"
//line parser.l:3
		

//line parser.l:7
type NutshSymType struct {
	yys int
    val string
}

const DEF = 57346
const IDENTIFIER = 57347
const LINE = 57348
const LINESEP = 57349
const IF = 57350
const PROMPT = 57351
const STRING = 57352
const AND = 57353
const OR = 57354
const NOT = 57355
const MATCH = 57356

var NutshToknames = []string{
	"DEF",
	"IDENTIFIER",
	"LINE",
	"LINESEP",
	"IF",
	"PROMPT",
	"STRING",
	"AND",
	"OR",
	"NOT",
	"MATCH",
}
var NutshStatenames = []string{}

const NutshEofCode = 1
const NutshErrCode = 2
const NutshMaxDepth = 200

//line parser.l:57


func main() {
    text := "ab"
    NutshParse(lexer{text: text})
}

//line yacctab:1
var NutshExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 20,
	14, 22,
	20, 22,
	-2, 24,
}

const NutshNprod = 30
const NutshPrivate = 57344

var NutshTokenNames []string
var NutshStates []string

const NutshLast = 73

var NutshAct = []int{

	20, 8, 21, 39, 8, 12, 49, 35, 19, 35,
	48, 27, 33, 34, 17, 42, 8, 25, 30, 26,
	47, 14, 15, 14, 31, 32, 24, 36, 41, 23,
	40, 37, 38, 35, 22, 18, 41, 41, 45, 46,
	9, 14, 43, 44, 10, 11, 13, 7, 33, 34,
	41, 15, 40, 50, 33, 34, 2, 14, 15, 6,
	16, 5, 24, 3, 29, 4, 1, 0, 0, 0,
	0, 0, 28,
}
var NutshPact = []int{

	36, -1000, -1000, 36, -1000, -1000, -1000, -1000, -3, 30,
	16, 7, -1000, 7, -7, 36, -1000, 18, 7, 43,
	-1000, 13, 16, 16, -1000, -1000, -1000, 52, -1, -1000,
	-3, -1000, -1000, 16, 16, 52, 52, 1, 37, -9,
	-11, -1000, -1000, 37, 37, -13, -13, -1000, -1000, 52,
	-1000,
}
var NutshPgo = []int{

	0, 66, 56, 65, 5, 63, 61, 59, 47, 0,
	8, 46, 3, 2,
}
var NutshR1 = []int{

	0, 1, 3, 4, 2, 2, 5, 5, 5, 5,
	5, 6, 7, 8, 8, 11, 11, 9, 9, 12,
	12, 13, 13, 13, 10, 10, 10, 10, 10, 10,
}
var NutshR2 = []int{

	0, 1, 3, 3, 0, 2, 1, 1, 1, 1,
	1, 3, 2, 1, 2, 1, 3, 1, 4, 1,
	3, 1, 1, 3, 1, 3, 3, 2, 3, 3,
}
var NutshChk = []int{

	-1000, -1, -2, -5, -3, -6, -7, -8, -9, 4,
	8, 9, -4, -11, 5, 15, -2, 17, 5, -10,
	-9, -13, 18, 13, 10, -4, -4, 18, -2, -11,
	-9, -4, -4, 11, 12, 20, 14, -10, -10, -12,
	-13, -9, 16, -10, -10, -13, -13, 19, 19, 17,
	-12,
}
var NutshDef = []int{

	4, -2, 1, 4, 6, 7, 8, 9, 10, 0,
	0, 0, 13, 0, 17, 4, 5, 0, 0, 0,
	-2, 0, 0, 0, 21, 12, 14, 0, 0, 16,
	15, 2, 11, 0, 0, 0, 0, 0, 27, 0,
	19, 22, 3, 28, 29, 23, 25, 26, 18, 0,
	20,
}
var NutshTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	18, 19, 3, 20, 17, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 15, 3, 16,
}
var NutshTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
}
var NutshTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var NutshDebug = 0

type NutshLexer interface {
	Lex(lval *NutshSymType) int
	Error(s string)
}

const NutshFlag = -1000

func NutshTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(NutshToknames) {
		if NutshToknames[c-4] != "" {
			return NutshToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func NutshStatname(s int) string {
	if s >= 0 && s < len(NutshStatenames) {
		if NutshStatenames[s] != "" {
			return NutshStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Nutshlex1(lex NutshLexer, lval *NutshSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = NutshTok1[0]
		goto out
	}
	if char < len(NutshTok1) {
		c = NutshTok1[char]
		goto out
	}
	if char >= NutshPrivate {
		if char < NutshPrivate+len(NutshTok2) {
			c = NutshTok2[char-NutshPrivate]
			goto out
		}
	}
	for i := 0; i < len(NutshTok3); i += 2 {
		c = NutshTok3[i+0]
		if c == char {
			c = NutshTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = NutshTok2[1] /* unknown char */
	}
	if NutshDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), NutshTokname(c))
	}
	return c
}

func NutshParse(Nutshlex NutshLexer) int {
	var Nutshn int
	var Nutshlval NutshSymType
	var NutshVAL NutshSymType
	NutshS := make([]NutshSymType, NutshMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Nutshstate := 0
	Nutshchar := -1
	Nutshp := -1
	goto Nutshstack

ret0:
	return 0

ret1:
	return 1

Nutshstack:
	/* put a state and value onto the stack */
	if NutshDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", NutshTokname(Nutshchar), NutshStatname(Nutshstate))
	}

	Nutshp++
	if Nutshp >= len(NutshS) {
		nyys := make([]NutshSymType, len(NutshS)*2)
		copy(nyys, NutshS)
		NutshS = nyys
	}
	NutshS[Nutshp] = NutshVAL
	NutshS[Nutshp].yys = Nutshstate

Nutshnewstate:
	Nutshn = NutshPact[Nutshstate]
	if Nutshn <= NutshFlag {
		goto Nutshdefault /* simple state */
	}
	if Nutshchar < 0 {
		Nutshchar = Nutshlex1(Nutshlex, &Nutshlval)
	}
	Nutshn += Nutshchar
	if Nutshn < 0 || Nutshn >= NutshLast {
		goto Nutshdefault
	}
	Nutshn = NutshAct[Nutshn]
	if NutshChk[Nutshn] == Nutshchar { /* valid shift */
		Nutshchar = -1
		NutshVAL = Nutshlval
		Nutshstate = Nutshn
		if Errflag > 0 {
			Errflag--
		}
		goto Nutshstack
	}

Nutshdefault:
	/* default state action */
	Nutshn = NutshDef[Nutshstate]
	if Nutshn == -2 {
		if Nutshchar < 0 {
			Nutshchar = Nutshlex1(Nutshlex, &Nutshlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if NutshExca[xi+0] == -1 && NutshExca[xi+1] == Nutshstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Nutshn = NutshExca[xi+0]
			if Nutshn < 0 || Nutshn == Nutshchar {
				break
			}
		}
		Nutshn = NutshExca[xi+1]
		if Nutshn < 0 {
			goto ret0
		}
	}
	if Nutshn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Nutshlex.Error("syntax error")
			Nerrs++
			if NutshDebug >= 1 {
				__yyfmt__.Printf("%s", NutshStatname(Nutshstate))
				__yyfmt__.Printf("saw %s\n", NutshTokname(Nutshchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Nutshp >= 0 {
				Nutshn = NutshPact[NutshS[Nutshp].yys] + NutshErrCode
				if Nutshn >= 0 && Nutshn < NutshLast {
					Nutshstate = NutshAct[Nutshn] /* simulate a shift of "error" */
					if NutshChk[Nutshstate] == NutshErrCode {
						goto Nutshstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if NutshDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", NutshS[Nutshp].yys)
				}
				Nutshp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if NutshDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", NutshTokname(Nutshchar))
			}
			if Nutshchar == NutshEofCode {
				goto ret1
			}
			Nutshchar = -1
			goto Nutshnewstate /* try again in the same state */
		}
	}

	/* reduction by production Nutshn */
	if NutshDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Nutshn, NutshStatname(Nutshstate))
	}

	Nutshnt := Nutshn
	Nutshpt := Nutshp
	_ = Nutshpt // guard against "declared and not used"

	Nutshp -= NutshR2[Nutshn]
	NutshVAL = NutshS[Nutshp+1]

	/* consult goto table to find next state */
	Nutshn = NutshR1[Nutshn]
	Nutshg := NutshPgo[Nutshn]
	Nutshj := Nutshg + NutshS[Nutshp].yys + 1

	if Nutshj >= NutshLast {
		Nutshstate = NutshAct[Nutshg]
	} else {
		Nutshstate = NutshAct[Nutshj]
		if NutshChk[Nutshstate] != -Nutshn {
			Nutshstate = NutshAct[Nutshg]
		}
	}
	// dummy call; replaced with literal code
	switch Nutshnt {

	}
	goto Nutshstack /* stack new state and value */
}
