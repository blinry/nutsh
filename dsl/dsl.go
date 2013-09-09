package dsl

import (
	"fmt"
	"regexp"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"strconv"
	"time"
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
			cmdline.Interrupt()
			select {
			case <-time.After(time.Second):
				break
			case <-c:
				cmdline.Interrupt()
				Say("Enter `exit` to quit the Nut Shell.")
			}
		}
	}()
}

func Query(query string) string {
	return strings.TrimSpace(cmdline.Query(" "+query))
}

func SimulatePrompt(query string) {
	lastCommand = query
	fmt.Println("$ "+query)
	lastOutput = cmdline.Query(query)
	fmt.Print(lastOutput)
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
	exec.Command("stty", "-F", "/dev/tty", "-echo", "-icanon", "min", "1").Run()
	lastCommand = cmdline.ReadCommand()
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	lastOutput, wasInteractive = cmdline.ReadOutput()
	Output()

	rows, columns := getsize()
	println(rows)
	println(columns)
	cmdline.Query("stty rows "+strconv.Itoa(rows))
	cmdline.Query("stty columns "+strconv.Itoa(columns))

	return true
}
