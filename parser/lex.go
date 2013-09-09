package parser

type lexer struct {
	text string
	pos  int
}

var first int
var pos int

func (l lexer) next() rune {
	if pos >= len([]rune(l.text)) {
		pos += 1
		return 0
	}
	c := []rune(l.text)[pos]
	pos += 1
	return c
}

func (l lexer) emit(lval *NutshSymType) {
	lval.val = string([]rune(l.text)[first : pos-1])
	pos -= 1
	first = pos
}

func (l lexer) skip() {
	first = pos
}

func (l lexer) Lex(lval *NutshSymType) int {
	var c rune

start:
	c = l.next()
	switch {
	case whitespace(c):
		l.skip()
		goto start
	case c == '"':
		//l.skip()
		c = l.next()
		for c != '"' {
			if c == '\\' {
				c = l.next()
				if c != '"' && c != '\\' {
					panic("Syntax error: Expected \" or \\ after \\.")
				}
			}
			c = l.next()
		}
		l.next()
		l.emit(lval)
		return STRING
	case c == '/':
		c = l.next()
		if c == '/' {
			c = l.next()
			for c != '\r' && c != '\n' {
				c = l.next()
			}
			l.skip()
			goto start
		} else if c == '*' {
			c = l.next()
		in_comment:
			for c != '*' {
				c = l.next()
			}
			c = l.next()
			if c != '/' {
				c = l.next()
				goto in_comment
			}
			c = l.next()
			l.skip()
			goto start
		} else {
			panic("Syntax error: Expected / or * after /.")
		}

	case c == '=':
		c = l.next()
		if c == '~' {
			c = l.next()
			l.emit(lval)
			return MATCH
		} else if c == '=' {
			c = l.next()
			l.emit(lval)
			return EQ
		} else {
			panic("Syntax error: Expected ~ or = after =.")
		}
	case c == '&':
		c = l.next()
		if c == '&' {
			c = l.next()
			l.emit(lval)
			return AND
		} else {
			panic("Syntax error: Expected & after &.")
		}
	case c == '|':
		c = l.next()
		if c == '|' {
			c = l.next()
			l.emit(lval)
			return OR
		} else {
			panic("Syntax error: Expected | after |.")
		}
	case c == '!':
		c = l.next()
		l.emit(lval)
		return NOT
	case alnum(c):
		for alnum(c) {
			c = l.next()
		}
		l.emit(lval)
		switch lval.val {
		case "if":
			return IF
		case "def":
			return DEF
		case "prompt":
			return PROMPT
		case "else":
			return ELSE
		default:
			return IDENTIFIER
		}
	}
	if c != 0 {
		l.next()
		l.emit(lval)
	}
	return int(c)
}

func alnum(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_'
}

func whitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func (l lexer) Error(e string) {
	panic(e)
}
