package dsl

import (
	"fmt"
	"regexp"
	"os"
	"os/signal"
	"strings"
	"morr.cc/nutsh.git/cli"
)

var (
	cmdline cli.CLI
	lastCommand, lastOutput string
	wasInteractive bool
	didOutput bool
)

func Spawn(target string) {
	cmdline = cli.Spawn(target)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			<-c
			Say("Press Ctrl-C again to quit.")
			<-c
			os.Exit(0)
		}
	}()
}

func Query(query string) string {
	return strings.TrimSpace(cmdline.Query(query))
}

func QueryOutput(query string, expression string) bool {
	output := cmdline.Query(query)
	return regexp.MustCompile(expression).MatchString(output)
}

func Say(text string) {
	text = regexp.MustCompile("`([^`]+)`").ReplaceAllString(text, "[32m$1[36m")
	text = regexp.MustCompile("\\s+").ReplaceAllString(text, " ")
	fmt.Printf("[36m\n    %s\n\n[0m", text)
}

func LastCommand() string {
	return strings.TrimSpace(lastCommand)
}

func LastOutput() string {
	return strings.TrimSpace(lastOutput)
}

func Command(expression string) bool {
	return regexp.MustCompile(expression).MatchString(lastCommand)
}

func OutputMatch(expression string) bool {
	return regexp.MustCompile(expression).MatchString(lastOutput)
}

func Output() {
	if !wasInteractive && !didOutput {
		fmt.Print(lastOutput)
		didOutput = true
	}
}

func Prompt() bool {
	didOutput = false
	lastCommand = cmdline.ReadCommand()
	lastOutput, wasInteractive = cmdline.ReadOutput()
	Output()
	return true
}
