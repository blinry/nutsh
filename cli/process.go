package cli

import (
	"bufio"
	"github.com/kr/pty"
	"os"
)

func startProcess(command string, stdin <-chan rune, stdout chan<- rune) {
	tty, _ := pty.Start(stringToCmd(command))

	input, _ := os.Create("/tmp/nutsh-input")
	output, _ := os.Create("/tmp/nutsh-output")

	go func() {
		for {
			r, ok := <-stdin
			if ! ok {
				return
			}
			input.Write([]byte(string(r)))
			tty.Write([]byte(string(r)))
		}
	}()

	go func() {
		reader := bufio.NewReader(tty)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				close(stdout)
				return
			}
			output.Write([]byte(string(r)))
			stdout <- r
		}
	}()
}
