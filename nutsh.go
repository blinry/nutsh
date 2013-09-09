package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	"morr.cc/nutsh.git/model"
	"morr.cc/nutsh.git/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nutsh (run|test) <tutorial dir>")
	}

	command := os.Args[1]
	dir := os.Args[2]

	tut := model.Init(dir)

	switch command {
	case "run":
		for {
			l := tut.SelectLesson()
			parser.Interpret(l.Root)
		}
	case "test":
		for _, l := range tut.Lessons {
			fmt.Println(l)
			//parser.test(l)
		}
	}

	/*
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
	*/
}
