%{
package lexer

import (
        "bufio"
        "os"
       )

var (
        src      = bufio.NewReader(os.Stdin)
        buf      []byte
        current  byte
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
		for _, b := range(bytetext) {
			buffer = append(buffer, b)
			println(string(buffer))
			bytes <- b
		}
	}()

	tokens := make(chan Token)

	go func() {
		for {
			typ := nextTokenType(bytes)
			tokens<- Token{typ, buffer}
			buffer = []byte("")
		}
	}()

	return tokens
}

func nextTokenType(bytes <-chan byte) tokenType {
	var c byte
%}

%yyc c
%yyn c = <-bytes

D [0-9]+
C [a-z][a-z0-9]+
alpha [a-zA-Z_]

%%

[ \t\n\r]+

def|return|if|break return typeKeyword

{alpha}+ return typeIdentifier

"("|")"|","|"+"|"=~"|"=="|"{"|"}" return typePunct

"\""[^"\""]*"\"" return typeString

\0 return typeEOF

. panic("Unexpected "+string(c))

%%
	//println("")
	//println(c)
	//panic("Unexpected")
	return typeEOF
}
