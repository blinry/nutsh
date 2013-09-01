package main

import (
	"fmt"
	"morr.cc/nutsh.git/parser"
)

func main() {
	text := `
prompt {
	if output =~ "help" {
		say("Sehr lustig, "+command)
		break
	}
}
`
	l := parser.Parse(text)
	fmt.Println(l)
	parser.Interpret(l)
}
