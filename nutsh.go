package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	"morr.cc/nutsh.git/model"
	"morr.cc/nutsh.git/parser"
	"strconv"
	"time"
)

var (
	logfile *os.File
)

func log(typ string, text string) {
	s := strconv.Quote(text)
	logfile.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)+"\t"+typ+"\t"+s+"\n"))
	logfile.Sync()
}

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
	logfile, _ = os.OpenFile(dir+"/log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	switch command {
	case "run":
		log("begin", "")
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
			lesson_name, ok = tut.SelectLesson(true)
			if ! ok {
				break
			}
		}
		for {
			l, _ = tut.Lessons[lesson_name]
			log("start", lesson_name)
			lesson_name, done = parser.Interpret(l.Root, tut.Common)
			if done {
				log("done", lesson_name)
				l.Done = true
				tut.SaveProgress()
			} else {
				log("quit", lesson_name)
			}
			if lesson_name != "" {
				l, exists = tut.Lessons[lesson_name]
				if exists {
					continue
				}
			}
			lesson_name, ok = tut.SelectLesson(done)
			if ! ok {
				log("exit", "")
				break
			}
		}
	case "test":
		if lesson_name != "" {
			l, _ := tut.Lessons[lesson_name]
			parser.Test(l.Root, tut.Common)
			fmt.Println(parser.GetName(l.Root)+" passed.")
		} else {
			for _, l := range tut.Lessons {
				parser.Test(l.Root, tut.Common)
			}
		}
	}
	logfile.Close()
}
