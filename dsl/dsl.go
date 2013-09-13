package dsl

import (
	"fmt"
	"morr.cc/nutsh.git/cli"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	cmdline                 cli.CLI
	lastCommand, lastOutput string
	wasInteractive          bool
	didOutput               bool
	running                 bool
)

func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			<-c
			if running {
				cmdline.Interrupt()
				select {
				case <-time.After(time.Second):
					break
				case <-c:
					cmdline.Interrupt()
					Say("Geben Sie zum Beenden der Nut Shell `exit` ein.")
				}
			}
		}
	}()
}

func Spawn(target string) {
	cmdline = cli.Spawn(target)
	running = true
}

func Query(query string) (string, bool) {
	s, ok := cmdline.Query(" " + query)
	if ! ok {
		return "", false
	}
	return strings.TrimSpace(s), true
}

func SimulatePrompt(query string) bool {
	lastCommand = query
	fmt.Println("$ " + query)
	lastOutput, ok := cmdline.Query(query)
	if ! ok {
		return false
	}
	fmt.Print(lastOutput)
	return true
}

func QueryOutput(query string, expression string) (bool, bool) {
	output, ok := cmdline.Query(query)
	if ! ok {
		return false, false
	}
	return regexp.MustCompile(expression).MatchString(output), true
}

func Say(text string) {
	text = regexp.MustCompile("`([^`]+)`").ReplaceAllString(text, "[32m$1[36m")
	text = regexp.MustCompile("\\s+").ReplaceAllString(text, " ")
	_, c := getsize()
	fmt.Printf("[36m\n%s\n\n[0m", indent(wrap(text, c-4), 4))
}

func wrap(text string, width int) string {
	ret := ""
	line_len := 0
	for _, w := range strings.Split(text, " ") {
		l := len(w)
		if line_len + l + 1 > width {
			ret += "\n"
			line_len = 0
		}
		ret += w+" "
		line_len += l + 1
	}
	return ret
}

func indent(text string, spaces int) string {
	iden := ""
	for i := 0; i < spaces; i++ {
		iden += " "
	}
	text = strings.Replace(text, "\n", "\n"+iden, -1)
	return iden+text
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
	rows, columns := getsize()
	cmdline.Query(" stty rows " + strconv.Itoa(rows))
	cmdline.Query(" stty columns " + strconv.Itoa(columns))

	didOutput = false
	exec.Command("stty", "-F", "/dev/tty", "-echo", "-icanon", "min", "1").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "sane").Run()
	var ok bool
	lastCommand, ok = cmdline.ReadCommand()
	if ! ok {
		// cli terminated
		return false
	}
	lastOutput, wasInteractive, ok = cmdline.ReadOutput()
	if ! ok {
		return false
	}
	Output()
	time.Sleep(time.Second)

	return true
}

func Quit() {
	cmdline.Quit()
	running = false
}
