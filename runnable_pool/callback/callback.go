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

			_callbacks <-func(result interface{},item RunnablePoolWithCallbackItem) func() {
				return func() {
					item.CallbackFun(result)
				}

			}(r,i)


		}
	}()
}

var _callbacks =make(chan func(),1<<8)

func init() {
	go func() {
		for {
			f := <-_callbacks
			f()
		}
	}()

}








