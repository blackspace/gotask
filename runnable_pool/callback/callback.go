package callback

import (
	. "github.com/blackspace/gotask/task"
)


type CallbackFun func(result interface{})

type RunnablePoolWithCallbackItem struct {
	Task
	CallbackFun
}

type RunnablePoolWithCallback struct {
	_channel   chan RunnablePoolWithCallbackItem
	_callbacks chan func()
}


func NewRunnablePoolWithCallback() *RunnablePoolWithCallback {
	return &RunnablePoolWithCallback{_channel:make(chan RunnablePoolWithCallbackItem,1<<8),_callbacks:make(chan func(),1<<8)}
}

func (tp *RunnablePoolWithCallback)AddTask(t Task,f CallbackFun){
	tp._channel<- RunnablePoolWithCallbackItem{t,f}
}

func (tp *RunnablePoolWithCallback)Run() {
	go func() {
		for {
			f := <-tp._callbacks
			f()
		}
	}()


	go func(){
		for {
			i:= <-tp._channel

			r:=i.Exec()

			tp._callbacks <-func(result interface{},item RunnablePoolWithCallbackItem) func() {
				return func() {
					item.CallbackFun(result)
				}

			}(r,i)


		}
	}()
}









