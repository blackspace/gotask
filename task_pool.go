package gotask


type TaskPool struct {
	Channel chan Task
}


func NewTaskPool() *TaskPool {
	return &TaskPool{Channel:make(chan Task,1<<8)}
}

func (tp *TaskPool)AddTask(t Task) {
	tp.Channel<-t
}


func (tp *TaskPool)Run() {
	go func(){
		for {
			t:= <-tp.Channel
			r:=t.Exec()
			t.SendResult(r)
		}
	}()
}
