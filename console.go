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


func NewTaskPrintln(s string) *TaskPrintln {
	return &TaskPrintln{TaskBase:TaskBase{CreatedAt:time.Now()},Content:s}
}

