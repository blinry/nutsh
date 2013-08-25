package main

import (
	. "morr.cc/nutsh.git/model"
)

func main() {
	m := Model{
		map[string]Lesson{
			"example": Lesson{
				map[string]State{
					"hi": State{
						Block{
							[]Statement{
								Command{OutputCommandType, "Hey! Say `hello`!"},
							},
						},
						Block{
							[]Statement{
								IfStatement{CommandIfType, "hello", "",
									Block{
										[]Statement{
											Command{OutputCommandType, "Nice!"},
											Command{GotoCommandType, "task"},
										},
									},
									Block{},
								},
							},
						},
					},

					"task": State{
						Block{
							[]Statement{
								Command{OutputCommandType, "Now go to `/tmp`."},
							},
						},
						Block{
							[]Statement{
								IfStatement{QueryOutputIfType, "pwd", "/tmp",
									Block{
										[]Statement{
											Command{OutputCommandType, "Nice!"},
											Command{GotoCommandType, "hi"},
										},
									},
									Block{},
								},
							},
						},
					},

				},
			},
		},
	}

	m.Interpret()
}
