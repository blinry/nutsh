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
	var last_lesson *model.Lesson
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
			l, ok = tut.SelectLesson(true)
			if ! ok {
				break
			}
		}
		for {
			last_lesson = l
			lesson_name, done = parser.Interpret(l.Root)
			if done {
				last_lesson.Done = true
				tut.SaveProgress()
			}
			if lesson_name != "" {
				l, exists = tut.Lessons[lesson_name]
				if exists {
					continue
				}
			}
			l, ok = tut.SelectLesson(done)
			if ! ok {
				break
			}
		}
	case "test":
		for _, l := range tut.Lessons {
			//fmt.Println(l)
			parser.Test(l.Root)
		}
	}
}
