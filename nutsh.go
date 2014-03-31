package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	"morr.cc/nutsh.git/model"
	"morr.cc/nutsh.git/parser"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: nutsh (run|test) <tutorial dir> [lesson name]")
	}

	command := os.Args[1]
	dir := os.Args[2]

	lesson_name := ""
	done := false

	if len(os.Args) > 3 {
		lesson_name = os.Args[3]
	}

	tut := model.Init(dir)

	switch command {
	case "run":
		var l *model.Lesson
		var exists bool
		var ok bool
		if lesson_name != "" {
			l, exists = tut.Lessons[lesson_name]
			if ! exists {
				lesson_name = ""
			}
		}
		if lesson_name == "" {
			lesson_name, ok = tut.SelectLesson(false)
			if ! ok {
				break
			}
		}
		for {
			l, _ = tut.Lessons[lesson_name]
			lesson_name, done = parser.Interpret(l.Root, tut.Common)
			if done {
				l.Done = true
				tut.SaveProgress()
			}
			if lesson_name != "" {
				l, exists = tut.Lessons[lesson_name]
				if exists {
					continue
				}
			}
			lesson_name, ok = tut.SelectLesson(false)
			if ! ok {
				break
			}
		}
	case "test":
		if lesson_name != "" {
			l, _ := tut.Lessons[lesson_name]
			parser.Test(l.Root, tut.Common)
			fmt.Println("\""+parser.GetName(l.Root)+"\""+" passed.")
		} else {
			for _, l := range tut.Lessons {
				parser.Test(l.Root, tut.Common)
			}
			fmt.Println("All lessons passed!")
		}
	}
}
