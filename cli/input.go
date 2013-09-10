package cli

import (
	"bufio"
	"io"
	"os"
)

func filterInput(input <-chan string, stdin chan<- rune, state *tokenizerState) {
	for {
		s, ok := <-input
		if ! ok {
			close(stdin)
			return
		}
		for _, r := range s {
			if r == 10 && *state == cmdinputState {
				stdin <- ''
				stdin <- ' '
				stdin <- ''
				stdin <- '☀'
				stdin <- ''
				stdin <- ''
				stdin <- '☀'
				stdin <- ''
				stdin <- ''
			}
			stdin <- r
		}
	}
}

func inputStdin(input chan<- string, state *tokenizerState) {
	reader := bufio.NewReader(os.Stdin)
	for *state != quitState {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		input <- string(r)
	}
	close(input)
	println("quit")
}
