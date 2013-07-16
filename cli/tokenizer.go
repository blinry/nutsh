package cli

type tokenType int

const (
	outputType tokenType = iota
	promptType
	partialCommandType
	finalCommandType
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
	queue := make([]rune, 0)

	for {

		var r rune
		if len(queue) > 0 {
			r = queue[0]
			queue = queue[1:len(queue)]
		} else {
			r = <-input
		}

		if r == '☃' {
			switch *state {
			case cmdinputState:
				*state++
				<-input
				<-input
				<-input
				<-input
			case cmdechoState:
				<-input
				<-input
				<-input
				<-input
				<-input
				<-input
				<-input
				<-input
				<-input
				<-input

				//queue = append(queue, <-input)
				r2 := <-input
				queue = append(queue, r2)

				if r2 == '★' {
					tokens <- token{partialCommandType, buffer[0 : len(buffer)-1]+"\n"}
					queue = make([]rune, 0)
					*state = promptState
				} else {
					tokens <- token{finalCommandType, buffer[0 : len(buffer)-1]+"\n"}
					*state++
				}
			case outputState:
				*state++
				tokens <- token{outputType, buffer}
			case promptState:
				*state = cmdinputState
				tokens <- token{promptType, buffer}
			}
			buffer = ""
		} else {
			buffer = buffer + string(r)
			if *state == cmdinputState || *state == cmdechoState {
				runes <- r
			}
		}
	}
}
