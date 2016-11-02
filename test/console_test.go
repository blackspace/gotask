package test

import (
	"github.com/blackspace/gotasks"
)

func ExampleTaskPrintln() {
	tl:=gotasks.NewTaskList()
	t0:=gotasks.NewTaskPrintln()
	t0.Content="Hello World"
	tl.AddPrepareTask(t0)
	t1:=tl.GetPrepareTask()
	t1.Exec()
	//Output: Hello World
}


func ExampleTaskPrintlnFromString() {
	task:=gotasks.BuildTaskFromString("*gotasks.TaskPrintln"+" "+"Hello World")
	task.Exec()
	//Output: Hello World
}