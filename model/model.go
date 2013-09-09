package model

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"morr.cc/nutsh.git/parser"
)

type Tutorial struct {
	Name    string
	Target  string
	Version int
	Basedir string
	Lessons map[string]Lesson
}

type Lesson struct {
	Root parser.Node
}

func Init(dir string) Tutorial {
	info, _ := ioutil.ReadFile(dir + "/info.yaml")
	var tut Tutorial
	goyaml.Unmarshal(info, &tut)
	tut.Basedir = dir
	tut.Lessons = make(map[string]Lesson)

	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.Name()[len(file.Name())-6:len(file.Name())] == ".nutsh" {
			content, _ := ioutil.ReadFile(dir + "/" + file.Name())
			rootnode := parser.Parse(string(content))
			tut.Lessons[file.Name()[0:len(file.Name())-6]] = Lesson{rootnode}
		}
	}

	return tut
}

func (t Tutorial) SelectLesson() Lesson {
	i := 0
	lessons := make([]Lesson, 0)
	for name, l := range t.Lessons {
		fmt.Printf("%d: %s\n", i, name)
		lessons = append(lessons, l)
		i += 1
	}

	sel := 0
tryagain:
	fmt.Print("Bitte wählen Sie eine Lektion: ")
	_, err := fmt.Scanf("%d", &sel)
	if err != nil {
		goto tryagain
	}

	if sel < 0 || sel > len(lessons)-1 {
		goto tryagain
	}
	return lessons[sel]
}
