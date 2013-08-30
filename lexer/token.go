package lexer

type Token struct {
	typ tokenType
	val []byte
}

type tokenType int

const (
	typeIdentifier tokenType = iota
	typeEOF
	typeKeyword
	typePunct
	typeString
)
