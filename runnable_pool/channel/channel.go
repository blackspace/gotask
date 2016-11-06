package channel

import (
	. "github.com/blackspace/gotask"
)

type RunnableChannelPool struct {
	_channel chan RunnableChannelPoolItem
}

type RunnableChannelPoolItem struct {
	C chan interface{}
	T Task
}

func NewRunnableChannelPool() *RunnableChannelPool {
	return &RunnableChannelPool{_channel:make(chan RunnableChannelPoolItem,1<<8)}
}

func (tp *RunnableChannelPool)AddTask(t Task) chan interface{} {
	c:=make(chan interface{},1)
	tp._channel <- RunnableChannelPoolItem{c,t}
	return c
}

func (tp *RunnableChannelPool)Run() {
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
