package channel

import (
	. "github.com/blackspace/gotask"
)

type RunnablePoolChannel struct {
	_channel chan RunnableChannelPoolItem
}

type RunnableChannelPoolItem struct {
	C chan interface{}
	T Task
}

func NewRunnablePoolChannel() *RunnablePoolChannel {
	return &RunnablePoolChannel{_channel:make(chan RunnableChannelPoolItem,1<<8)}
}

func (tp *RunnablePoolChannel)AddTask(t Task) chan interface{} {
	c:=make(chan interface{},1)
	tp._channel <- RunnableChannelPoolItem{c,t}
	return c
}

func (tp *RunnablePoolChannel)Run() {
	go func(){
		for {
			i:= <-tp._channel

			go func() {
				r:=i.T.Exec()
				i.C<-r
			}()

		}
	}()
}
