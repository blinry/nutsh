package cli

import (
	"bufio"
	"io"
	"os"
)

func filterInput(input <-chan string, stdin chan<- rune, state *tokenizerState) {
	for {
		s := <-input
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

func inputStdin(input chan<- string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}
			input <- string(r)
		}
	}
}
