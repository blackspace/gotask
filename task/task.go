package task

import (
	"time"
)

type Task interface {
	Exec() interface{}
	SetStartAt()
	SetEndAt()
	String() string
}


type TaskBase struct {
	CreatedAt time.Time
	StartAt   time.Time
	EndAt     time.Time
}

func (t *TaskBase)SetStartAt() {
	t.StartAt=time.Now()
}

func (t *TaskBase)SetEndAt() {
	t.EndAt=time.Now()
}







