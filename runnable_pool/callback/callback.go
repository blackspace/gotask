package callback

import (
	"github.com/blackspace/goevent"
	. "github.com/blackspace/gotask"
)


type RunnablePoolWithCallbackItem struct {
	Task
	*goevent.Event
}

type RunnablePoolWithCallback struct {
	_channel      chan RunnablePoolWithCallbackItem
}


func NewRunnablePoolWithCallback() *RunnablePoolWithCallback {
	return &RunnablePoolWithCallback{_channel:make(chan RunnablePoolWithCallbackItem,1<<8)}
}

func (tp *RunnablePoolWithCallback)AddTaskWithCallback(t Task,h goevent.Handler){
	e:=goevent.NewEvent()
	e.AddHandler(h)
	tp._channel<- RunnablePoolWithCallbackItem{t,e}
}

func (tp *RunnablePoolWithCallback)Run() {
	go func(){
		for {
			i:= <-tp._channel
			r:=i.Exec()
			i.Event.Fire(i,r)
		}
	}()
}

