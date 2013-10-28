%{

package parser

import (
	"strings"
	"strconv"
)

var lesson *Node

type Node struct {
	typ string
	children []*Node
}

func node(typ string, n ...*Node) *Node {
	return &Node{typ: typ, children: n}
}

func (n *Node) String() string {
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
	node *Node
}

%token <val> DEF IDENTIFIER LINE LINESEP IF ELSE PROMPT STRING AND OR NOT MATCH EQ
%type <node> lesson toplevel_lines toplevel_line line call stringexpressions stringexpression definition arguments identifiers block line lines ifclause prompt state call calls string

%left OR
%left AND
%left NOT
%left EQ MATCH
%left '+'

%%

lesson: toplevel_lines { $$ = node("lesson", $1); lesson = $$ }

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
| string { $$ = node("call", node("say"), node("stringexpressions", $1)) }

ifclause: IF stringexpression block { $$ = node("if", $2, $3, node("block")) }
| IF stringexpression block ELSE block { $$ = node("if", $2, $3, $5) }

prompt: PROMPT block { $$ = node("prompt", $2) }

state: calls block { $$ = node("state", $1, $2) }

identifiers: /* empty */ { $$ = node("identifiers") }
| IDENTIFIER { $$ = node("identifiers", node("id", node($1))) }
| identifiers ',' IDENTIFIER { $$ = node("identifiers", append($1.children, node($3))...) }

calls: call { $$ = node("block", $1) }
| calls ',' call { $$ = node("block", append($1.children, $3)...) }

call: IDENTIFIER { $$ = node("call", node($1), node("stringexpressions")) }
| IDENTIFIER '(' stringexpressions ')' { $$ = node("call", node($1), $3) }

stringexpressions: stringexpression { $$ = node("stringexpressions", $1) }
| stringexpressions ',' stringexpression { $$ = node("stringexpressions", append($1.children, $3)...) }

stringexpression: string { $$ = $1 }
| call { $$ = $1 }
| stringexpression '+' stringexpression { $$ = node("+", $1, $3) }
| stringexpression MATCH stringexpression { $$ = node("call", node("match"), node("stringexpressions", $1, $3)) }
| stringexpression EQ stringexpression { $$ = node("call", node("equal"), node("stringexpressions", $1, $3)) }
| '(' stringexpression ')' { $$ = $2 }
| NOT stringexpression { $$ = node("not", $2) }
| stringexpression AND stringexpression { $$ = node("and", $1, $3) }
| stringexpression OR stringexpression { $$ = node("or", $1, $3) }

string: STRING { s := strings.Replace($1, "\n", " ", -1); s, _ = strconv.Unquote(s); $$ = node("string", node(s)) }

%%

func Parse(text string) *Node {
	pos = 0
	first = 0
	NutshParse(lexer{text: text})
	return lesson
}
