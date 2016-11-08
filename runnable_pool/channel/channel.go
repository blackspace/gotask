package channel

import (
	. "github.com/blackspace/gotask"
)

type RunnablePoolWithChannel struct {
	_channel chan RunnablePoolWithChannelItem
}

type RunnablePoolWithChannelItem struct {
	C chan interface{}
	T Task
}

func NewRunnablePoolWithChannel() *RunnablePoolWithChannel {
	return &RunnablePoolWithChannel{_channel:make(chan RunnablePoolWithChannelItem,1<<8)}
}

func (tp *RunnablePoolWithChannel)AddTask(t Task) chan interface{} {
	c:=make(chan interface{},1)
	tp._channel <- RunnablePoolWithChannelItem{c,t}
	return c
}

func (tp *RunnablePoolWithChannel)Run() {
	go func(){
		for {
			i:= <-tp._channel

			go func(item RunnablePoolWithChannelItem) {
				r:=item.T.Exec()
				i.C<-r
			}(i)

		}
	}()
}
