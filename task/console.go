package task

import (
	"fmt"
)

type Println struct {
	*TaskBase
	Content string
}

func (pl *Println)Exec() interface{} {
	fmt.Println(pl.Content)
	return nil
}

func (pl *Println)String() string {
	return pl.Content
}

func BuildTaskPrintlnFromString(s string) Task {
	t:=NewTaskPrintln()
	t.Content=s
	return t
}


func NewTaskPrintln() *Println {
	return &Println{TaskBase:NewTaskBase()}
}

