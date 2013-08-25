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
	GotoCommandType
)

type IfStatement struct {
	IfType
	String1 string
	String2 string
	TrueBlock Block
	FalseBlock Block
}

type IfType int
const (
	QueryOutputIfType IfType = iota
	CommandIfType
)

func (m *Model) Interpret() {
	dsl.Spawn("bash")

	nextLesson := "example"
	for {
		nextState := "hi"
		for {
			state := m.Lessons[nextLesson].States[nextState]
			interpretBlock(state.InitBlock)
			for {
				dsl.Prompt()
				dsl.Output()
				gotoState := interpretBlock(state.LoopBlock)
				if gotoState != "" {
					switch gotoState {
					default:
						nextState = gotoState
					}
					break
				}
			}
		}
	}
}

func interpretBlock(b Block) string {
	for _, s := range b.Statements {
		switch s.(type) {
		case Command:
			c := s.(Command)
			switch c.CommandType {
			case OutputCommandType:
				dsl.Say(c.String)
			case GotoCommandType:
				return c.String
			}
		case IfStatement:
			c := s.(IfStatement)
			var value bool
			switch c.IfType {
			case CommandIfType:
				value = dsl.OutputMatch(c.String1)
			case QueryOutputIfType:
				value = dsl.QueryOutput(c.String1, c.String2)
			}

			var s string
			if value {
				s = interpretBlock(c.TrueBlock)
			} else {
				s = interpretBlock(c.FalseBlock)
			}
			if s != "" {
				return s
			}
		}
	}
	return ""
}
