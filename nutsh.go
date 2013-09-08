package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"morr.cc/nutsh.git/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nutsh (run|test) <lesson file>")
	}

	command := os.Args[1]
	file := os.Args[2]

	text, _ := ioutil.ReadFile(file)
	l := parser.Parse(string(text))

	//println(string(text))

	switch command {
	case "run":
		parser.Interpret(l)
	case "test":
		parser.Test(l)
	}

	//fmt.Println(l)
}
