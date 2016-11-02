package test

import (
	. "github.com/blackspace/gotasks/task"
)

func ExampleTaskPrintln() {
	t0:=NewTaskPrintln()
	t0.Content="Hello World"
	t0.Exec()
	//Output: Hello World
}


func ExampleTaskPrintlnFromString() {
	task:=TaskFromString("Println"+" "+"Hello World")
	task.Exec()
	//Output: Hello World
}

func ExampleTaskPrintlnToString() {
	t0:=NewTaskPrintln()
	t0.Content="Hello World"

	s:=TaskToString(t0)

	task:=TaskFromString(s)
	task.Exec()
	//Output: Hello World
}