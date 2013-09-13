package cli

import (
	"bufio"
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
			if *state == cmdinputState {
				if r == 10 {
					stdin <- ''
					stdin <- ' '
					stdin <- ''
					stdin <- '☀'
					stdin <- ''
					stdin <- ''
					stdin <- '☀'
					stdin <- ''
					stdin <- ''
				} else if r == 4 {
					continue
				}
			}
			stdin <- r
		}
	}
}

func runeToString(input <-chan rune, output chan<- string, quit <-chan bool) {
	outer: for {
		select {
		case r, ok := <-input:
			if ! ok {
				break outer
			}
			output <- string(r)
		case <-quit:
			break outer
		}
	}
	close(output)
}

func readStdin(input chan<- rune) {
	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		input <- r
	}
	close(input)
}
