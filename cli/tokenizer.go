package cli

import (
	"time"
)

type tokenType int

const (
	outputType tokenType = iota
	promptType
	partialCommandType
	finalCommandType
	endType
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
	quitState
)

func tokenize(input <-chan rune, tokens chan<- token, runes chan<- rune, state *tokenizerState) {
	buffer := ""
	queue := make([]rune, 0)
	interactive := false
	timer := time.NewTimer(0)
	timer.Stop()
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <- quit:
				return
			case <-timer.C:
				interactive = true
				for _, r := range buffer {
					runes <- r
				}
			}
		}
	}()

	for {

		var r rune
		if len(queue) > 0 {
			r = queue[0]
			queue = queue[1:len(queue)]
		} else {
			var ok bool
			r, ok = <-input
			if ! ok {
				quit <- true
				tokens <- token{endType, ""}
				close(tokens)
				close(runes)
				return
			}
		}

		if r == '☃' {
			switch *state {
			case outputState:
				*state++
				tokens <- token{outputType, buffer}
				interactive = false
				timer.Stop()
			case promptState:
				*state = cmdinputState
				tokens <- token{promptType, buffer}
			}
			buffer = ""
		} else if r == '☀' {
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
				var r2 rune
				select {
				case r2 = <-input:
					queue = append(queue, r2)
				case <-time.After(10 * time.Millisecond):
				}

				if r2 == '★' {
					tokens <- token{partialCommandType, buffer[0:len(buffer)-1] + "\n"}
					queue = make([]rune, 0)
					*state = promptState
				} else {
					tokens <- token{finalCommandType, buffer[0:len(buffer)-1] + "\n"}
					*state++
					interactive = false
					timer.Reset(500 * time.Millisecond)
				}
			}
			buffer = ""
		} else {
			buffer = buffer + string(r)
			if interactive || *state == cmdinputState || *state == cmdechoState {
				runes <- r
			}
		}
	}
}
