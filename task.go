package gotask

import (
	"time"
)

type Task interface {
	Exec() interface{}
	SetStartAt()
	SetEndAt()
	String() string
	SendResult(interface{})
	ReceiveResult() interface{}
}


type TaskBase struct {
	CreatedAt time.Time
	StartAt   time.Time
	EndAt     time.Time
	C      chan interface{}
}

func NewTaskBase() *TaskBase{
	return &TaskBase{CreatedAt:time.Now(),C:make(chan interface{})}
}

func (t *TaskBase)SetStartAt() {
	t.StartAt=time.Now()
}

func (t *TaskBase)SetEndAt() {
	t.EndAt=time.Now()
}

func (t *TaskBase)SendResult(r interface{}) {
	t.C<-r
}

func (t *TaskBase)ReceiveResult() interface{} {
	return <-t.C
}








