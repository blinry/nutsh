%{

package parser

%}

%union {
    val string
}

%token <val> DEF IDENTIFIER LINE LINESEP IF PROMPT STRING AND OR NOT MATCH

%%

lesson: lines

definition: DEF IDENTIFIER block

block: '{' lines '}'

lines: /* nothing */
| line lines

line: definition
| ifclause
| prompt
| state
| call

ifclause: IF boolexpression block

prompt: PROMPT block

state: block
| calls block

calls: call
| call ',' calls

call: IDENTIFIER
| IDENTIFIER '(' stringexpressions ')'

stringexpressions: stringexpression
| stringexpression ',' stringexpressions

stringexpression: STRING
| call
| stringexpression '+' stringexpression

boolexpression: call
| stringexpression MATCH stringexpression
| '(' boolexpression ')'
| NOT boolexpression
| boolexpression AND boolexpression
| boolexpression OR boolexpression

%%

func main() {
    text := "ab"
    NutshParse(lexer{text: text})
}
