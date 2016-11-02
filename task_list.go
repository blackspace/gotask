package gotasks


type TaskList struct {
	_data []Task
}

func NewTaskList() *TaskList {
	return &TaskList{_data:make([]Task,0,1<<8)}
}


func (tl *TaskList)AddTask(t Task) {
	tl._data =append(tl._data,t)
}

func (tl *TaskList)GetTask() Task {
	if len(tl._data)>0 {
		t:=tl._data[0]
		tl._data=tl._data[1:]
		return t
	} else {
		return nil
	}
}


