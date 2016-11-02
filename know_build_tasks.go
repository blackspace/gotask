package gotasks

import (
	"strings"
	"reflect"
)

type BuildTaskFunc func (s string) Task

type BuildFunc struct{
	ReturenType string
	BuildFunc BuildTaskFunc
}

type TaskSet struct {
	_data map[string]BuildFunc
}

func NewKnowBuildTasks() *TaskSet {
	return &TaskSet{_data:make(map[string]BuildFunc)}
}

func (k *TaskSet)Add(n string,t string,b BuildTaskFunc) {
	k._data[n]=BuildFunc{t,b}
}

func (k *TaskSet)GetByName(s string)  BuildTaskFunc {
	if b,ok:=k._data[s];ok {
		return b.BuildFunc
	} else {
		return 	nil
	}
}

func (k *TaskSet)GetByType(s string)  (string,BuildFunc,bool) {
	for lk,lv:=range k._data {
		if lv.ReturenType==s {
			return lk,lv,true
		}
	}

	return "",BuildFunc{},false
}

var buildTaskFuncSet *TaskSet

func init() {
	buildTaskFuncSet =NewKnowBuildTasks()

	buildTaskFuncSet.Add("Println",reflect.TypeOf((*Println)(nil)).String(),BuildTaskPrintlnFromString)
}


func BuildTask(s string) Task {
	i:=strings.IndexRune(s,' ')
	n:=s[:i]
	a:=s[i+1:]

	if f:= buildTaskFuncSet.GetByName(n);f!=nil {
		return f(a)
	} else {
		panic("Can't build task from "+s)
	}
}

func TaskToString(t Task) string {
	if n,_,ok:=buildTaskFuncSet.GetByType(reflect.TypeOf(t).String());ok {
		return n+" "+t.String()
	} else {
		panic("Can't find type "+n)
	}
}






