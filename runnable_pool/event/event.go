package event

import (
	"github.com/blackspace/goevent"
	. "github.com/blackspace/gotask"
)

type RunnablePoolEvent struct {
	TaskDoneEvent *goevent.Event
	_channel      chan Task
}


func NewRunnablePoolEvent() *RunnablePoolEvent {
	return &RunnablePoolEvent{TaskDoneEvent:goevent.NewEvent(), _channel:make(chan Task,1<<8)}
}

func (tp *RunnablePoolEvent)AddTask(t Task) {
	tp._channel<-t

}

func (tp *RunnablePoolEvent)Run() {
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

