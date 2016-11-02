package test

import (
	. "github.com/blackspace/gotasks/task"
)

func ExamplePrintln() {
	t0:=NewTaskPrintln()
	t0.Content="Hello World"
	t0.Exec()
	//Output: Hello World
}


func ExamplePrintlnFromString() {
	task:=TaskFromString("Println"+" "+"Hello World")
	task.Exec()
	//Output: Hello World
}

func ExamplePrintlnToString() {
	t0:=NewTaskPrintln()
	t0.Content="Hello World"

	s:=TaskToString(t0)

	task:=TaskFromString(s)
	task.Exec()
	//Output: Hello World
}