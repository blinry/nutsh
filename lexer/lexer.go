package lexer

import (
	"bufio"
	"os"
)

var (
	src     = bufio.NewReader(os.Stdin)
	buf     []byte
	current byte
)

func getc() byte {
	if current != 0 {
		buf = append(buf, current)
	}
	current = 0
	if b, err := src.ReadByte(); err == nil {
		current = b
	}
	return current
}

func Lex(text string) chan Token {
	bytetext := []byte(text)
	bytes := make(chan byte)
	buffer := []byte("")

	go func() {
		for _, b := range bytetext {
			buffer = append(buffer, b)
			println(string(buffer))
			bytes <- b
		}
	}()

	tokens := make(chan Token)

	go func() {
		for {
			typ := nextTokenType(bytes)
			tokens <- Token{typ, buffer}
			buffer = []byte("")
		}
	}()

	return tokens
}

func nextTokenType(bytes <-chan byte) tokenType {
	var c byte

yystate0:

	goto yystart1

	goto yystate1 // silence unused label error
yystate1:
	c = <-bytes
yystart1:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c == '!' || c >= '#' && c <= '\'' || c == '*' || c >= '-' && c <= '<' || c >= '>' && c <= '@' || c >= '[' && c <= '^' || c == '`' || c == '|' || c >= '~' && c <= 'ÿ'
	case c == '"':
		goto yystate6
	case c == '(' || c == ')' || c == '+' || c == ',' || c == '{' || c == '}':
		goto yystate9
	case c == '=':
		goto yystate10
	case c == '\n':
		goto yystate5
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == 'b':
		goto yystate14
	case c == 'd':
		goto yystate19
	case c == 'i':
		goto yystate21
	case c == 'r':
		goto yystate22
	case c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'c' || c >= 'e' && c <= 'h' || c >= 'j' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate12
	}

yystate2:
	c = <-bytes
	goto yyrule6

yystate3:
	c = <-bytes
	goto yyrule7

yystate4:
	c = <-bytes
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate5:
	c = <-bytes
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate6:
	c = <-bytes
	switch {
	default:
		goto yyrule7
	case c == '"':
		goto yystate8
	case c >= '\x01' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate7
	}

yystate7:
	c = <-bytes
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate8
	case c >= '\x01' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate7
	}

yystate8:
	c = <-bytes
	goto yyrule5

yystate9:
	c = <-bytes
	goto yyrule4

yystate10:
	c = <-bytes
	switch {
	default:
		goto yyrule7
	case c == '=' || c == '~':
		goto yystate11
	}

yystate11:
	c = <-bytes
	goto yyrule4

yystate12:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate13
	}

yystate13:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate13
	}

yystate14:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'r':
		goto yystate15
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate13
	}

yystate15:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'e':
		goto yystate16
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate13
	}

yystate16:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'a':
		goto yystate17
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate13
	}

yystate17:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'k':
		goto yystate18
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate13
	}

yystate18:
	c = <-bytes
	switch {
	default:
		goto yyrule2
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate13
	}

yystate19:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'e':
		goto yystate20
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate13
	}

yystate20:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'f':
		goto yystate18
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate13
	}

yystate21:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'f':
		goto yystate18
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate13
	}

yystate22:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'e':
		goto yystate23
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate13
	}

yystate23:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 't':
		goto yystate24
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate13
	}

yystate24:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'u':
		goto yystate25
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate13
	}

yystate25:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'r':
		goto yystate26
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate13
	}

yystate26:
	c = <-bytes
	switch {
	default:
		goto yyrule3
	case c == 'n':
		goto yystate18
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate13
	}

yyrule1: // [ \t\n\r]+

	goto yystate0
yyrule2: // def|return|if|break
	{
		return typeKeyword
	}
yyrule3: // {alpha}+
	{
		return typeIdentifier
	}
yyrule4: // "("|")"|","|"+"|"=~"|"=="|"{"|"}"
	{
		return typePunct
	}
yyrule5: // "\""[^"\""]*"\""
	{
		return typeString
	}
yyrule6: // \0
	{
		return typeEOF
	}
yyrule7: // .
	{
		panic("Unexpected " + string(c))
		goto yystate0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	//println("")
	//println(c)
	//panic("Unexpected")
	return typeEOF
}
