package callback

import (
	. "github.com/blackspace/gotask"
)


type CallbackFun func(result interface{})

type RunnablePoolWithCallbackItem struct {
	Task
	CallbackFun
}

type RunnablePoolWithCallback struct {
	_channel      chan RunnablePoolWithCallbackItem
}


func NewRunnablePoolWithCallback() *RunnablePoolWithCallback {
	return &RunnablePoolWithCallback{_channel:make(chan RunnablePoolWithCallbackItem,1<<8)}
}

func (tp *RunnablePoolWithCallback)AddTask(t Task,f CallbackFun){
	tp._channel<- RunnablePoolWithCallbackItem{t,f}
}

func (tp *RunnablePoolWithCallback)Run() {
	go func(){
		for {
			i:= <-tp._channel

			r:=i.Exec()
			i.CallbackFun(r)
		}
	}()
}

