package model

import (
	"morr.cc/nutsh.git/dsl"
)

type Model struct {
	Lessons map[string]Lesson
}

type Lesson struct {
	States map[string]State
}

type State struct {
	InitBlock Block;
	LoopBlock Block;
}

type Block struct {
	Statements []Statement
}

type Statement interface {
}

type Command struct {
	CommandType
	String string
}

type CommandType int
const (
	ExecuteCommandType CommandType = iota
	OutputCommandType
)

type IfStatement struct {
	IfType
	String1 string
	String2 string
	TrueBlock *Block
	FalseBlock *Block
}

type IfType int
const (
	QueryIfType IfType = iota
	CommandIfType
)

func (m *Model) Interpret() {
	dsl.Spawn("bash")

	nextLesson := "example"
	for {
		nextState := "hi"
		for {
			state := m.Lessons[nextLesson].States[nextState]
			interpretBlock(&state.InitBlock)
			for {
				dsl.Prompt()
				dsl.Output()
				interpretBlock(&state.LoopBlock)
			}
		}
	}
}

func interpretBlock(b *Block) {
	for _, s := range b.Statements {
		switch s.(type) {
		case Command:
			c := s.(Command)
			switch c.CommandType {
			case OutputCommandType:
				dsl.Say(c.String)
			}
		case IfStatement:
			c := s.(IfStatement)
			switch c.IfType {
			case CommandIfType:
				if dsl.OutputMatch(c.String1) {
					interpretBlock(c.TrueBlock)
				} else {
					interpretBlock(c.FalseBlock)
				}
			}
		}
	}
}
