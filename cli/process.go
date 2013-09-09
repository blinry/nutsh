package cli

import (
	"bufio"
	"github.com/kr/pty"
	"os"
	"os/exec"
)

func Quit() {
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	os.Exit(0)
}

func startProcess(command string, stdin <-chan rune, stdout chan<- rune) {
	tty, _ := pty.Start(stringToCmd(command))

	input, _ := os.Create("/tmp/nutsh-input")
	output, _ := os.Create("/tmp/nutsh-output")

	go func() {
		for {
			r := <-stdin
			/*rows, cols, err := pty.Getsize(tty)
			if err != nil {
				panic(err)
			}
			print(rows, cols)
			*/
			input.Write([]byte(string(r)))
			tty.Write([]byte(string(r)))
		}
	}()

	go func() {
		reader := bufio.NewReader(tty)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				Quit()
			}
			output.Write([]byte(string(r)))
			stdout <- r
		}
	}()
}
