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
	var last_lesson model.Lesson
	done := false

	if len(os.Args) > 3 {
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
			last_lesson = l
			lesson_name, done = parser.Interpret(l.Root)
			if done {
				last_lesson.Done = true
			}
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
