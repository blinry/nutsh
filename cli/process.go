package cli

import (
	"os"
	"os/exec"
	"bufio"
    "github.com/kr/pty"
)

func init() {
    exec.Command("stty", "-F", "/dev/tty", "-echo", "-icanon", "min", "1").Run()
}

func startProcess(command string, stdin <-chan rune, stdout chan<- rune) {
	tty, _ := pty.Start(stringToCmd(command))

	input, _ := os.Create("/tmp/nutsh-input")
	output, _ := os.Create("/tmp/nutsh-output")

	go func() {
		for {
			r := <- stdin
			input.Write([]byte(string(r)))
			tty.Write([]byte(string(r)))
		}
	}()

	go func() {
		reader := bufio.NewReader(tty)
		for {
			r, _, _ := reader.ReadRune()
			output.Write([]byte(string(r)))
			stdout <- r
		}
	}()
}
