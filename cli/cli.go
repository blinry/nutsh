package cli

import (
	"fmt"
)

type CLI struct {
	tokens chan token
	runes chan rune
	input chan string
}

func Spawn(t target) CLI {
	stdin := make(chan rune)
	stdout := make(chan rune)
	tokens := make(chan token)
	runes := make(chan rune)
	input := make(chan string)
	state := outputState

	c := CLI{tokens, runes, input}

	go startProcess(t.spawnCmd, stdin, stdout)
	go tokenize(stdout, tokens, runes, &state)
	go inputStdin(input)
	go filterInput(input, stdin, &state)

	c.send(t.initCmd)
	c.read(promptType)
	//print("!!!")
	
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
	c.input <- s
}

func (c CLI) read(k tokenType) string {
	//print("expect")
	//print(int(k))
	for {
		select {
		case t := <- c.tokens:
			//fmt.Printf("token %v: %q\n", t.tokenType, t.string)
			if t.tokenType == k {
				return t.string
			}
		case r := <- c.runes:
			fmt.Printf(string(r))
		}
	}
}
