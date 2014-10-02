package cli

import (
	"fmt"
	"time"
)

var (
	input chan rune
)

func init() {
	input = make(chan rune)
}

func UseStdin() {
	go readStdin(input)
}

// CLI represents an command line interface instance.
type CLI struct {
	tokens             chan token
	runes              chan rune
	input              chan string
	allowInteractivity bool
	quit               chan bool
}

// Spawn starts a new instance of the target.
func Spawn(target string) CLI {

	t := targets[target]

	stdin := make(chan rune)
	stdout := make(chan rune)
	tokens := make(chan token)
	runes := make(chan rune)
	stringInput := make(chan string)
	quit := make(chan bool)

	state := outputState

	c := CLI{tokens, runes, stringInput, true, quit}

	go startProcess(t.spawnCmd, stdin, stdout)
	go tokenize(stdout, tokens, runes, &state)
	go runeToString(input, stringInput, quit)
	go filterInput(stringInput, stdin, &state)

	c.send(t.initCmd)
	c.read(outputType)

	return c
}

// ReadOutput waits for the next output token and returns it.
func (c CLI) ReadOutput() (string, bool, bool) {
	return c.read(outputType)
}

// ReadCommand waits for the next command token and returns it.
func (c CLI) ReadCommand() (string, bool) {
	command := ""

	prompt, _, ok := c.read(promptType)
	if ! ok {
		return "", false
	}
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
				return command, true
			case endType:
				println("endtype")
				return "", false
			}
		case r := <-c.runes:
			if c.allowInteractivity {
				fmt.Print(string(r))
			}
		}
	}
}

// Query executes cmd and returns the output.
func (c CLI) Query(cmd string) (string, bool) {
	return c.QueryInteractive(cmd,"")
}

func (c CLI) QueryInteractive(cmd string, interaction string) (string, bool) {
	_, _, ok := c.read(promptType)
	if ! ok {
		return "", false
	}
	c.send(cmd)
	c.send("\n")
	c.allowInteractivity = false
	if interaction != "" {
		c.allowInteractivity = true
		time.Sleep(500*time.Millisecond)
		c.send(interaction)
	}
	o, _, ok := c.ReadOutput()
	if ! ok {
		return "", false
	}
	c.allowInteractivity = true

	return o, true
}

func (c CLI) Interrupt() {
	c.send("")
}

func (c CLI) send(s string) {
	c.input <- s
}

func (c CLI) read(k tokenType) (data string, wasInteractive bool, ok bool) {
	wasInteractive = false
	ok = true
	for {
		select {
		case t := <-c.tokens:
			if t.tokenType == k {
				data = t.string
				return
			} else if t.tokenType == endType {
				ok = false
				return
			}
		case r := <-c.runes:
			if c.allowInteractivity {
				wasInteractive = true
				fmt.Print(string(r))
			}
		}
	}
}

func (c CLI) Quit() {
	c.quit <- true
}

func GetInput() chan rune {
	return input
}
