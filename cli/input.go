package cli

func filterInput(input <-chan string, stdin chan<- rune, state *tokenizerState) {
	for {
		s := <- input
		for _, r := range s {
			if r == 10 && *state == cmdinputState {
				stdin <- ' '
				stdin <- ''
				stdin <- '☃'
				stdin <- ''
				stdin <- ''
				stdin <- '☃'
				stdin <- ''
			}
			stdin <- r
		}
	}
}
