package parser

import(
	"testing"
	"fmt"
)

func TestAvancedParse(t *testing.T) {
	l := lexer{text: `
say("hi")
say("fu")
def greet(name) {
	say("Hi, "+name)
}

fufu {
	prompt {
		if command =~ "fu" {
			greet("Seb")
		}
	}
}
`}
	Parse(l)
	fmt.Print(lesson)
}
