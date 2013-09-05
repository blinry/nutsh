
//line parser.l:2

package parser
import __yyfmt__ "fmt"
//line parser.l:3
		
import (
	"strings"
	"strconv"
)

var lesson Node

type Node struct {
	typ string
	children []Node
}

func node(typ string, n ...Node) Node {
	return Node{typ: typ, children: n}
}

func (n Node) String() string {
	r := n.typ
	r += ":"
	for _, child := range(n.children) {
		lines := strings.Split(child.String(), "\n")
		for _, line := range(lines) {
			r += "\n    "+line
		}
	}
	return r
}

//line parser.l:34
type NutshSymType struct {
	yys int
    val string
	node Node
}

const DEF = 57346
const IDENTIFIER = 57347
const LINE = 57348
const LINESEP = 57349
const IF = 57350
const ELSE = 57351
const PROMPT = 57352
const STRING = 57353
const AND = 57354
const OR = 57355
const NOT = 57356
const MATCH = 57357
const EQ = 57358

var NutshToknames = []string{
	"DEF",
	"IDENTIFIER",
	"LINE",
	"LINESEP",
	"IF",
	"ELSE",
	"PROMPT",
	"STRING",
	"AND",
	"OR",
	"NOT",
	"MATCH",
	"EQ",
}
var NutshStatenames = []string{}

const NutshEofCode = 1
const NutshErrCode = 2
const NutshMaxDepth = 200

//line parser.l:100


func Parse(text string) Node {
	pos = 0
	first = 0
	NutshParse(lexer{text: text})
	return lesson
}

//line yacctab:1
var NutshExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 10,
	19, 24,
	21, 24,
	-2, 15,
}

const NutshNprod = 40
const NutshPrivate = 57344

var NutshTokenNames []string
var NutshStates []string

const NutshLast = 81

var NutshAct = []int{

	18, 23, 20, 19, 5, 10, 11, 34, 35, 15,
	32, 33, 12, 51, 13, 16, 25, 31, 56, 29,
	30, 57, 36, 37, 52, 24, 24, 26, 41, 39,
	42, 27, 46, 47, 48, 49, 50, 54, 45, 60,
	55, 10, 11, 53, 44, 34, 35, 58, 32, 33,
	34, 35, 24, 32, 33, 31, 59, 15, 6, 15,
	31, 15, 12, 16, 13, 16, 22, 17, 14, 21,
	9, 8, 7, 38, 43, 28, 4, 40, 3, 2,
	1,
}
var NutshPact = []int{

	-1000, -1000, 54, -1000, -1000, -1000, 62, -1000, -1000, -1000,
	-1000, -1000, 52, 7, 6, 14, -1000, 2, 33, -1000,
	-1000, 52, 52, -1000, -1000, -1000, 56, 52, 7, 39,
	29, 52, 52, 52, 52, 52, -5, 38, 4, -1000,
	19, 38, -1000, 0, -1000, 7, 38, 38, 38, 38,
	38, -1000, -1000, -1000, -1000, 52, -1000, 34, -1000, 38,
	-1000,
}
var NutshPgo = []int{

	0, 80, 79, 78, 4, 2, 77, 0, 76, 75,
	74, 1, 73, 72, 71, 70, 68, 3,
}
var NutshR1 = []int{

	0, 1, 8, 9, 9, 11, 2, 2, 3, 3,
	12, 12, 4, 4, 4, 4, 4, 13, 13, 14,
	15, 10, 10, 10, 16, 16, 5, 5, 6, 6,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 17,
}
var NutshR2 = []int{

	0, 1, 4, 0, 3, 3, 0, 2, 1, 1,
	0, 2, 1, 1, 1, 1, 1, 3, 5, 2,
	2, 0, 1, 3, 1, 3, 1, 4, 1, 3,
	1, 1, 3, 3, 3, 3, 2, 3, 3, 1,
}
var NutshChk = []int{

	-1000, -1, -2, -3, -8, -4, 4, -13, -14, -15,
	-5, -17, 8, 10, -16, 5, 11, 5, -7, -17,
	-5, 17, 14, -11, 19, -11, 21, 17, -9, 17,
	-11, 22, 15, 16, 12, 13, -7, -7, -12, -5,
	-6, -7, -11, -10, 5, 9, -7, -7, -7, -7,
	-7, 18, 20, -4, 18, 21, 18, 21, -11, -7,
	5,
}
var NutshDef = []int{

	6, -2, 1, 7, 8, 9, 0, 12, 13, 14,
	-2, 16, 0, 0, 0, 26, 39, 3, 0, 30,
	31, 0, 0, 19, 10, 20, 0, 0, 0, 21,
	17, 0, 0, 0, 0, 0, 0, 36, 0, 25,
	0, 28, 2, 0, 22, 0, 32, 33, 34, 37,
	38, 35, 5, 11, 27, 0, 4, 0, 18, 29,
	23,
}
var NutshTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	17, 18, 3, 22, 21, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 19, 3, 20,
}
var NutshTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
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

	case 1:
		//line parser.l:44
		{ NutshVAL.node = NutshS[Nutshpt-0].node; lesson = NutshVAL.node }
	case 2:
		//line parser.l:46
		{ NutshVAL.node = node("def", node(NutshS[Nutshpt-2].val), NutshS[Nutshpt-1].node, NutshS[Nutshpt-0].node) }
	case 3:
		//line parser.l:48
		{ NutshVAL.node = node("arguments") }
	case 4:
		//line parser.l:49
		{ NutshVAL.node = node("arguments", NutshS[Nutshpt-1].node.children...) }
	case 5:
		//line parser.l:51
		{ NutshVAL.node = NutshS[Nutshpt-1].node }
	case 6:
		//line parser.l:53
		{ NutshVAL.node = node("block") }
	case 7:
		//line parser.l:54
		{ NutshVAL.node = node("block", append(NutshS[Nutshpt-1].node.children, NutshS[Nutshpt-0].node)...) }
	case 8:
		//line parser.l:56
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 9:
		//line parser.l:57
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 10:
		//line parser.l:59
		{ NutshVAL.node = node("block") }
	case 11:
		//line parser.l:60
		{ NutshVAL.node = node("block", append(NutshS[Nutshpt-1].node.children, NutshS[Nutshpt-0].node)...) }
	case 12:
		//line parser.l:62
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 13:
		//line parser.l:63
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 14:
		//line parser.l:64
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 15:
		//line parser.l:65
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 16:
		//line parser.l:66
		{ NutshVAL.node = node("call", node("say"), node("stringexpressions", NutshS[Nutshpt-0].node)) }
	case 17:
		//line parser.l:68
		{ NutshVAL.node = node("if", NutshS[Nutshpt-1].node, NutshS[Nutshpt-0].node, node("block")) }
	case 18:
		//line parser.l:69
		{ NutshVAL.node = node("if", NutshS[Nutshpt-3].node, NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node) }
	case 19:
		//line parser.l:71
		{ NutshVAL.node = node("prompt", NutshS[Nutshpt-0].node) }
	case 20:
		//line parser.l:73
		{ NutshVAL.node = node("state", NutshS[Nutshpt-1].node, NutshS[Nutshpt-0].node) }
	case 21:
		//line parser.l:75
		{ NutshVAL.node = node("identifiers") }
	case 22:
		//line parser.l:76
		{ NutshVAL.node = node("identifiers", node("id", node(NutshS[Nutshpt-0].val))) }
	case 23:
		//line parser.l:77
		{ NutshVAL.node = node("identifiers", append(NutshS[Nutshpt-2].node.children, node(NutshS[Nutshpt-0].val))...) }
	case 24:
		//line parser.l:79
		{ NutshVAL.node = node("block", NutshS[Nutshpt-0].node) }
	case 25:
		//line parser.l:80
		{ NutshVAL.node = node("block", append(NutshS[Nutshpt-2].node.children, NutshS[Nutshpt-0].node)...) }
	case 26:
		//line parser.l:82
		{ NutshVAL.node = node("call", node(NutshS[Nutshpt-0].val), node("stringexpressions")) }
	case 27:
		//line parser.l:83
		{ NutshVAL.node = node("call", node(NutshS[Nutshpt-3].val), NutshS[Nutshpt-1].node) }
	case 28:
		//line parser.l:85
		{ NutshVAL.node = node("stringexpressions", NutshS[Nutshpt-0].node) }
	case 29:
		//line parser.l:86
		{ NutshVAL.node = node("stringexpressions", append(NutshS[Nutshpt-2].node.children, NutshS[Nutshpt-0].node)...) }
	case 30:
		//line parser.l:88
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 31:
		//line parser.l:89
		{ NutshVAL.node = NutshS[Nutshpt-0].node }
	case 32:
		//line parser.l:90
		{ NutshVAL.node = node("+", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node) }
	case 33:
		//line parser.l:91
		{ NutshVAL.node = node("call", node("match"), node("stringexpressions", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node)) }
	case 34:
		//line parser.l:92
		{ NutshVAL.node = node("call", node("equal"), node("stringexpressions", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node)) }
	case 35:
		//line parser.l:93
		{ NutshVAL.node = NutshS[Nutshpt-1].node }
	case 36:
		//line parser.l:94
		{ NutshVAL.node = node("not", NutshS[Nutshpt-0].node) }
	case 37:
		//line parser.l:95
		{ NutshVAL.node = node("and", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node) }
	case 38:
		//line parser.l:96
		{ NutshVAL.node = node("or", NutshS[Nutshpt-2].node, NutshS[Nutshpt-0].node) }
	case 39:
		//line parser.l:98
		{ s, _ := strconv.Unquote(strings.Replace(NutshS[Nutshpt-0].val, "\n", " ", -1)); NutshVAL.node = node("string", node(s)) }
	}
	goto Nutshstack /* stack new state and value */
}
