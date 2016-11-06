package event

import (
	"github.com/blackspace/goevent"
	. "github.com/blackspace/gotask"
)

type RunnableEventPool struct {
	TaskDoneEvent *goevent.Event
	_channel      chan Task
}


func NewRunnableEventPool() *RunnableEventPool {
	return &RunnableEventPool{TaskDoneEvent:goevent.NewEvent(), _channel:make(chan Task,1<<8)}
}

func (tp *RunnableEventPool)AddTask(t Task) {
	tp._channel<-t

}

func (tp *RunnableEventPool)Run() {
	go func(){
		for {
			i:= <-tp._channel

			go func() {
				r:=i.Exec()
				tp.TaskDoneEvent.Fire(i,r)
			}()

		}
	}()
}

