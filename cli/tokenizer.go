package cli

type tokenType int
const (
	outputType tokenType = iota
	promptType
	commandType
)

type token struct {
	tokenType
	string
}

type tokenizerState int
const (
    cmdinputState tokenizerState = iota
    cmdechoState
    outputState
    promptState
)

func tokenize(input <-chan rune, tokens chan<- token, runes chan<- rune, state *tokenizerState) {
	buffer := ""

	for {
		r := <- input

        if r == '☃' {
			switch *state {
			case cmdinputState:
				*state++
				<- input
				<- input
				<- input
				<- input
			case cmdechoState:
				*state++
				<- input
				<- input
				<- input
				<- input
				<- input
				<- input
				tokens <- token{commandType, buffer[0:len(buffer)-1]}
			case outputState:
				*state++
				tokens <- token{outputType, buffer}
			case promptState:
				*state = cmdinputState
				tokens <- token{promptType, buffer}
			}
			buffer = ""
		} else {
			buffer = buffer+string(r)
			if *state == cmdinputState {
				runes <- r
			}
		}
	}
}
