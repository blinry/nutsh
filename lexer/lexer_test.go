package lexer

import (
	"testing"
)

func equalTest(t *testing.T, s1, s2 interface{}) {
	if s1 != s2 {
		t.Fatalf("Expected %q, got %q.", s2, s1)
	}
}

func TestCreate(t *testing.T) {
	tokens := Lex("say")
	var token Token
	token = <-tokens

	equalTest(t, token.val, "say")
	equalTest(t, token.typ, typeIdentifier)
}
