%{

package parser

import (
	"strings"
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
%}

%union {
    val string
	node Node
}

%token <val> DEF IDENTIFIER LINE LINESEP IF PROMPT STRING AND OR NOT MATCH
%type <node> lesson toplevel_lines toplevel_line line call stringexpressions stringexpression definition arguments identifiers block line lines ifclause prompt state call boolexpression calls

%%

lesson: toplevel_lines { $$ = $1; lesson = $$ }

definition: DEF IDENTIFIER arguments block { $$ = node("def", node($2), $3, $4) }

arguments: /* empty */ { $$ = node("arguments") }
| '(' identifiers ')' { $$ = node("arguments", $2.children...) }

block: '{' lines '}' { $$ = $2 }

toplevel_lines: /* empty */ { $$ = node("block") }
| toplevel_lines toplevel_line { $$ = node("block", append($1.children, $2)...) }

toplevel_line: definition { $$ = $1 }
| line { $$ = $1 }

lines: /* empty */ { $$ = node("block") }
| lines line { $$ = node("block", append($1.children, $2)...) }

line: ifclause { $$ = $1 }
| prompt { $$ = $1 }
| state { $$ = $1 }
| call { $$ = $1 }

ifclause: IF boolexpression block { $$ = node("if", $2, $3) }

prompt: PROMPT block { $$ = node("prompt", $2) }

state: calls block { $$ = node("state", $1, $2) }

identifiers: /* empty */ { $$ = node("identifiers") }
| IDENTIFIER { $$ = node("identifiers", node("id", node($1))) }
| identifiers ',' IDENTIFIER { $$ = node("identifiers", append($1.children, node($3))...) }

calls: call { node("calls", $1) }
| calls ',' call { node("calls", append($1.children, $3)...) }

call: IDENTIFIER { $$ = node("call", node($1), node("stringexpressions")) }
| IDENTIFIER '(' stringexpressions ')' { $$ = node("call", node($1), $3) }

stringexpressions: stringexpression { $$ = node("stringexpressions", $1) }
| stringexpressions ',' stringexpression { $$ = node("stringexpressions", append($1.children, $3)...) }

stringexpression: STRING { $$ = node("string", node($1)) }
| call { $$ = $1 }
| stringexpression '+' stringexpression { $$ = node("+", $1, $3) }

boolexpression: call { $$ = $1 }
| stringexpression MATCH stringexpression { $$ = node("call", node("match"), node("stringexpressions", $1, $3)) }
| '(' boolexpression ')' { $$ = $2 }
| NOT boolexpression { $$ = node("not", $2) }
| boolexpression AND boolexpression { $$ = node("and", $1, $3) }
| boolexpression OR boolexpression { $$ = node("or", $1, $3) }

%%

func Parse(text string) Node {
	pos = 0
	first = 0
	NutshParse(lexer{text: text})
	return lesson
}
