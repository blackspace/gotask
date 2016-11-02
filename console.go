package gotasks

import (
	"fmt"
	"time"
)

type Println struct {
	TaskBase
	Content string
}

func (pl *Println)Exec() bool{
	fmt.Println(pl.Content)
	return true
}

func (pl *Println)String() string {
	return pl.Content
}

func BuildTaskPrintlnFromString(s string) Task {
	t:=NewTaskPrintln()
	t.CreatedAt=time.Now()
	t.Content=s
	return t
}


func NewTaskPrintln() *Println {
	return &Println{}
}

