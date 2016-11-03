package gotask


type RunnablePool struct {
	Channel chan RunnablePoolItem
}

type RunnablePoolItem struct {
	C chan interface{}
	T Task
}

func NewRunnablePool() *RunnablePool {
	return &RunnablePool{Channel:make(chan RunnablePoolItem,1<<8)}
}

func (tp *RunnablePool)AddTask(t Task) chan interface{} {
	c:=make(chan interface{})
	tp.Channel<- RunnablePoolItem{c,t}
	return c
}

func (tp *RunnablePool)Run() {
	go func(){
		for {
			i:= <-tp.Channel
			r:=i.T.Exec()
			i.C<-r
		}
	}()
}
