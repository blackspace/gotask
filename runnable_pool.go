package gotask


type RunnablePool struct {
	_channel chan RunnablePoolItem
}

type RunnablePoolItem struct {
	C chan interface{}
	T Task
}

func NewRunnablePool() *RunnablePool {
	return &RunnablePool{_channel:make(chan RunnablePoolItem,1<<8)}
}

func (tp *RunnablePool)AddTask(t Task) chan interface{} {
	c:=make(chan interface{},1)
	tp._channel <- RunnablePoolItem{c,t}
	return c
}

func (tp *RunnablePool)Run() {
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
