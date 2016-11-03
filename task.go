package gotask

import (
	"time"
)

type Task interface {
	Exec() interface{}
	SetStartAt()
	SetEndAt()
}


type TaskBase struct {
	CreatedAt time.Time
	StartAt   time.Time
	EndAt     time.Time
}

func NewTaskBase() *TaskBase{
	return &TaskBase{CreatedAt:time.Now()}
}

func (t *TaskBase)SetStartAt() {
	t.StartAt=time.Now()
}

func (t *TaskBase)SetEndAt() {
	t.EndAt=time.Now()
}








