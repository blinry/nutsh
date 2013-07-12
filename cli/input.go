package cli

type input struct {
	stdin chan<- rune
	state *tokenizerState
}

func (i input) write(s string) {
	for _, r := range s {
		if r == 10 && *i.state == cmdinputState {
			i.stdin <- ' '
			i.stdin <- ''
			i.stdin <- '☃'
			i.stdin <- ''
			i.stdin <- ''
			i.stdin <- '☃'
			i.stdin <- ''
		}
		i.stdin <- r
	}
}
