package parser

import(
	"testing"
)

func equalTest(t *testing.T, s1, s2 interface{}) {
	if s1 != s2 {
		t.Fatalf("Expected %q, got %q.", s2, s1)
	}
}

func lexTest(t *testing.T, l lexer, expectedTyp int, expectedVal string) {
	var val NutshSymType
	typ := l.Lex(&val)
	equalTest(t, typ, expectedTyp)
	equalTest(t, val.val, expectedVal)
}

func TestSimple(t *testing.T) {
	pos = 0
	first = 0
	l := lexer{text: `say("hello")`}
	lexTest(t, l, IDENTIFIER, "say")
	lexTest(t, l, '(', "(")
	lexTest(t, l, STRING, `"hello"`)
	lexTest(t, l, ')', ")")
	//lexTest(t, l, 0, "")
}

func TestAvanced(t *testing.T) {
	pos = 0
	first = 0
	l := lexer{text: `
def greet(name) {
	say("Hi, "+name)
}

prompt {
	if command =~ "hi" {
		greet("Seb")
	}
}
`}

	lexTest(t, l, DEF, "def")
	lexTest(t, l, IDENTIFIER, "greet")
	lexTest(t, l, '(', "(")
	lexTest(t, l, IDENTIFIER, "name")
	lexTest(t, l, ')', ")")
	lexTest(t, l, '{', "{")
	lexTest(t, l, IDENTIFIER, "say")
	lexTest(t, l, '(', "(")
	lexTest(t, l, STRING, `"Hi, "`)
	lexTest(t, l, '+', "+")
	lexTest(t, l, IDENTIFIER, "name")
	lexTest(t, l, ')', ")")
	lexTest(t, l, '}', "}")

	lexTest(t, l, PROMPT, "prompt")
	lexTest(t, l, '{', "{")
	lexTest(t, l, IF, "if")
	lexTest(t, l, IDENTIFIER, "command")
	lexTest(t, l, MATCH, "=~")
	lexTest(t, l, STRING, `"hi"`)
	lexTest(t, l, '{', "{")
	lexTest(t, l, IDENTIFIER, "greet")
	lexTest(t, l, '(', "(")
	lexTest(t, l, STRING, `"Seb"`)
	lexTest(t, l, ')', ")")
	lexTest(t, l, '}', "}")
	lexTest(t, l, '}', "}")
}
