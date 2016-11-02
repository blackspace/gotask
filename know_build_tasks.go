package gotasks

import (
	"reflect"
	"strings"
)

type BuildTaskFunc func (s string) Task

type KnownBuildTasks struct {
	_data map[string]BuildTaskFunc
}

func NewKnowBuildTasks() *KnownBuildTasks {
	return &KnownBuildTasks{_data:make(map[string]BuildTaskFunc)}
}

func (k *KnownBuildTasks)Add(t reflect.Type,f BuildTaskFunc) {
	k._data[t.String()]=f
}

func (k *KnownBuildTasks)Get(s string)  BuildTaskFunc {
	if f,ok:=k._data[s];ok {
		return f
	} else {
		return 	nil
	}
}


var known *KnownBuildTasks

func init() {
	known=NewKnowBuildTasks()

	known.Add(reflect.TypeOf((*TaskPrintln)(nil)),BuildTaskPrintlnFromString)
}


func BuildTaskFromString(s string) Task {
	i:=strings.IndexRune(s,' ')
	n:=s[:i]
	a:=s[i+1:]

	if f:=known.Get(n);f!=nil {
		return f(a)
	} else {
		panic(s)
	}
}




