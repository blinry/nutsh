package cli

type input struct {
	stdin chan<- rune
}

func (i input) write(s string) {
	for _, r := range s {
		i.stdin <- r
	}
}
