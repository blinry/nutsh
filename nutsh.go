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

	low := -1
	high := 999

	// dirty hack for the Vorkurs:
	if _, err := os.Stat(dir+"/all"); err != nil {
		_, _, day := time.Now().Date()
		switch day {
			case 16: // Mo
			high = 4
			case 17: // Di
			high = 9
			case 18: // Mi
			high = 14
			case 19: // Do
			high = 19
			case 20: // Fr
			high = 21
			case 21: // Sa
			high = 21
			case 22: // So
			high = 21
			case 23: // Mo
			high = 25
			case 24: // Di
			high = 26
			case 25: // Mi
			high = 26
			/*
			case 26: // Do
			case 27: // Fr
			*/
		}
	}

	if lesson_name != "" {
		low = model.NameToNumber(lesson_name)
		high = model.NameToNumber(lesson_name)
	}

	tut := model.Init(dir, low, high)
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
			lesson_name, ok = tut.SelectLesson(false)
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
			lesson_name, ok = tut.SelectLesson(false)
			if ! ok {
				log("exit", "")
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
	logfile.Close()
}
