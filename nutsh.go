package main

import (
	//"fmt"
	"os"
	"io/ioutil"
	"morr.cc/nutsh.git/parser"
)

func main() {
	file := os.Args[1]
	text, _ := ioutil.ReadFile(file)
	//println(string(text))
	l := parser.Parse(string(text))
	//fmt.Println(l)
	parser.Interpret(l)
}
