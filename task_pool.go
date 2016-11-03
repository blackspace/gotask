package gotask


type TaskPool struct {
	Channel chan PoolItem
}

type PoolItem struct {
	C chan interface{}
	T Task
}

func NewTaskPool() *TaskPool {
	return &TaskPool{Channel:make(chan PoolItem,1<<8)}
}

func (tp *TaskPool)AddTask(t Task) chan interface{} {
	c:=make(chan interface{})
	tp.Channel<-PoolItem{c,t}
	return c
}

func (tp *TaskPool)Run() {
	go func(){
		for {
			pi:= <-tp.Channel
			r:=pi.T.Exec()
			pi.C<-r
		}
	}()
}
