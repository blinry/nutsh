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
		fmt.Println("Usage: nutsh (run|test) <tutorial dir> [lesson name]")
	}

	command := os.Args[1]
	dir := os.Args[2]

	lesson_name := ""
	if len(os.Args) > 2 {
		lesson_name = os.Args[3]
	}

	tut := model.Init(dir)

	switch command {
	case "run":
		var l model.Lesson
		var exists bool
		if lesson_name == "" {
			l = tut.SelectLesson()
		} else {
			l, exists = tut.Lessons[lesson_name]
			if ! exists {
				l = tut.SelectLesson()
			}
		}
		for {
			lesson_name = parser.Interpret(l.Root)
			println(lesson_name)
			if lesson_name != "" {
				l, exists = tut.Lessons[lesson_name]
				if exists {
					continue
				}
			}
			l = tut.SelectLesson()
		}
	case "test":
		for _, l := range tut.Lessons {
			fmt.Println(l)
			//parser.test(l)
		}
	}
}
