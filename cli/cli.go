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
func (c CLI) ReadOutput() string {
	return c.read(outputType)
}

// WasInteractive return true if the last ReadOutput invocation turned
// interactive. In that case, the output was already printed.
func (c CLI) WasInteractive() bool {
	return false
}

// ReadCommand waits for the next command token and returns it.
func (c CLI) ReadCommand() string {
	command := ""

	fmt.Print(c.read(promptType))
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
	c.allowInteractivity = false
	o := c.ReadOutput()
	c.allowInteractivity = true
	return o
}

func (c CLI) send(s string) {
	c.input <- s
}

func (c CLI) read(k tokenType) string {
	for {
		select {
		case t := <-c.tokens:
			if t.tokenType == k {
				return t.string
			}
		case r := <-c.runes:
			if c.allowInteractivity {
				fmt.Printf(string(r))
			}
		}
	}
}
