package parser

import(
	"testing"
)

func TestSimpleParse(t *testing.T) {
	pos = 0
	first = 0
	l := lexer{text: `say`}
	NutshParse(l)
}

func TestAvancedParse(t *testing.T) {
	pos = 0
	first = 0
	l := lexer{text: `
def greet(name) {
	say("Hi, "+name)
}
`}
	NutshParse(l)
}
