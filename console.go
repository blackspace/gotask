package gotasks

import (
	"fmt"
	"time"
)

type TaskPrintln struct {
	TaskBase
	Content string
}

func (pl *TaskPrintln)Exec() bool{
	fmt.Println(pl.Content)
	return true
}

func (pl *TaskPrintln)String() string {
	return pl.Content
}

func BuildTaskPrintlnFromString(s string) Task {
	t:=NewTaskPrintln()
	t.CreatedAt=time.Now()
	t.Content=s
	return t
}


func NewTaskPrintln() *TaskPrintln {
	return &TaskPrintln{}
}

