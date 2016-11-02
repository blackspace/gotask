package task


type TaskPool struct {
	_data []Task
}

func NewTaskPool() *TaskPool {
	return &TaskPool{_data:make([]Task,0,1<<8)}
}


func (tl *TaskPool)AddTask(t Task) {
	tl._data =append(tl._data,t)
}

func (tl *TaskPool)GetTask() Task {
	if len(tl._data)>0 {
		t:=tl._data[0]
		tl._data=tl._data[1:]
		return t
	} else {
		return nil
	}
}


