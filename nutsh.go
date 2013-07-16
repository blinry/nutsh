package main

import (
	"fmt"
	"morr.cc/nutsh.git/cli"
)

func main() {
	c := cli.Spawn("bash")

	for {
		cmd := c.ReadCommand()
		fmt.Printf("[35m\n\n    you entered: %q\n\n[0m", cmd)

		output, wasInteractive := c.ReadOutput()

		if (! wasInteractive) {
			fmt.Print("[36m    command was non-interactive\n\n[0m")
			fmt.Print(output)
		} else {
			fmt.Print("[36m\n    command was interactive\n[0m")
		}

		fmt.Printf("[32m\n    the output was: %q\n\n[0m", output)
	}
}
