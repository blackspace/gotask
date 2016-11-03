package gotask


type RunablePool struct {
	Channel chan _PoolItem
}

type _PoolItem struct {
	C chan interface{}
	T Task
}

func NewRunablePool() *RunablePool {
	return &RunablePool{Channel:make(chan _PoolItem,1<<8)}
}

func (tp *RunablePool)AddTask(t Task) chan interface{} {
	c:=make(chan interface{})
	tp.Channel<- _PoolItem{c,t}
	return c
}

func (tp *RunablePool)Run() {
	go func(){
		for {
			pi:= <-tp.Channel
			r:=pi.T.Exec()
			pi.C<-r
		}
	}()
}
