package cli

import (
	"fmt"
)

// CLI represents an command line interface instance.
type CLI struct {
	tokens chan token
	runes  chan rune
	input  chan string
	allowInteractivity bool
}

// Spawn starts a new instance of the target.
func Spawn(target string) CLI {

	t := targets[target]

	stdin := make(chan rune)
	stdout := make(chan rune)
	tokens := make(chan token)
	runes := make(chan rune)
	input := make(chan string)
	state := outputState

	c := CLI{tokens, runes, input, true}

	go startProcess(t.spawnCmd, stdin, stdout)
	go tokenize(stdout, tokens, runes, &state)
	go inputStdin(input)
	go filterInput(input, stdin, &state)

	c.send(t.initCmd)
	c.read(outputType)

	return c
}

// ReadOutput waits for the next output token and returns it.
func (c CLI) ReadOutput() (string, bool) {
	return c.read(outputType)
}

// ReadCommand waits for the next command token and returns it.
func (c CLI) ReadCommand() string {
	command := ""

	prompt, _ := c.read(promptType)
	fmt.Print(prompt)
	for {
		select {
		case t := <-c.tokens:
			switch t.tokenType {
			case promptType:
				fmt.Print("\r\n")
				fmt.Print(t.string)
			case partialCommandType:
				command += t.string
			case finalCommandType:
				command += t.string
				fmt.Print("\r\n")
				return command
			}
		case r := <-c.runes:
			if c.allowInteractivity {
				fmt.Printf(string(r))
			}
		}
	}
}

// Query executes cmd and returns the output.
func (c CLI) Query(cmd string) string {
	c.read(promptType)
	c.send(cmd)
	c.send("\n")
	c.allowInteractivity = false
	o, _ := c.ReadOutput()
	c.allowInteractivity = true
	return o
}

func (c CLI) send(s string) {
	c.input <- s
}

func (c CLI) read(k tokenType) (data string, wasInteractive bool) {
	wasInteractive = false
	for {
		select {
		case t := <-c.tokens:
			if t.tokenType == k {
				data = t.string
				return
			}
		case r := <-c.runes:
			if c.allowInteractivity {
				wasInteractive = true
				fmt.Printf(string(r))
			}
		}
	}
}
