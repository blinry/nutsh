package main

import (
	. "morr.cc/nutsh.git/model"
)

func main() {
	m := Model{}
	l := Lesson{}
	s := State{}

	initBlock := Block{}
	greetCommand := Command{OutputCommandType, "Hi!"}
	initBlock.Statements = append(initBlock.Statements, greetCommand)
	
	loopBlock := Block{}
	trueBlock := Block{}
	trueBlock.Statements = append(trueBlock.Statements, greetCommand)
	ifStatement := IfStatement{CommandIfType, "hello", "", &trueBlock, &Block{}}
	loopBlock.Statements = append(loopBlock.Statements, ifStatement)

	s.InitBlock = initBlock
	s.LoopBlock = loopBlock

	l.States = make(map[string]State)
	l.States["hi"] = s
	m.Lessons = make(map[string]Lesson)
	m.Lessons["example"] = l

	m.Interpret()
}
