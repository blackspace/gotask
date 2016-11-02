package test

import (
	"github.com/blackspace/gotasks"
)

func ExampleTaskPrintln() {
	tl:=gotasks.NewTaskList()
	t0:=gotasks.NewTaskPrintln()
	t0.Content="Hello World"
	tl.AddTask(t0)
	t1:=tl.GetTask()
	t1.Exec()
	//Output: Hello World
}


func ExampleTaskPrintlnFromString() {
	task:=gotasks.TaskFromString("Println"+" "+"Hello World")
	task.Exec()
	//Output: Hello World
}

func ExampleTaskPrintlnToString() {
	t0:=gotasks.NewTaskPrintln()
	t0.Content="Hello World"

	s:=gotasks.TaskToString(t0)

	task:=gotasks.TaskFromString(s)
	task.Exec()
	//Output: Hello World
}