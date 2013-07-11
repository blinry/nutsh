package cli

import (
	"fmt"
)

type CLI struct {
	tokens chan token
	runes chan rune
	input
}

func Spawn(t target) CLI {
	stdin := make(chan rune)
	stdout := make(chan rune)

	c := CLI{make(chan token), make(chan rune), input{stdin}}

	go startProcess(t.spawnCmd, stdin, stdout)
	go tokenize(stdout, c.tokens, c.runes)
	go inputStdin(c.input)

	c.send(t.initCmd)
	
	return c
}

func (c CLI) ReadOutput() string {
	return c.read(outputType)
}

func (c CLI) WasInteractive() bool {
	return false
}

func (c CLI) ReadCommand() string {
	fmt.Print(c.read(promptType))
	return c.read(commandType)
}

func (c CLI) Query(cmd string) string {
	c.send(cmd)
	o := c.ReadOutput()
	return o
}

func (c CLI) send(s string) {
	c.input.write(s)
}

func (c CLI) read(k tokenType) string {
	for {
		select {
		case t := <- c.tokens:
			if t.tokenType == outputType {
				return t.string
			}
		case r := <- c.runes:
			fmt.Printf(string(r))
		}
	}
}
