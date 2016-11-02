package test

import (
	"github.com/blackspace/gotasks"
)

func ExamplePrintln() {
	tl:=gotasks.NewTaskList()
	tl.AddPrepareTask(gotasks.NewTaskPrintln("Hello World"))
	t:=tl.GetPrepareTask()
	t.Exec()
	//Output: Hello World
}
