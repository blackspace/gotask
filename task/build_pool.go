package task

import (
	"strings"
	"reflect"
)

type BuildFunc func (s string) Task

type Build struct{
	_type string
	_func BuildFunc
}

type BuildPool struct {
	_data map[string]Build
}

func NewBuildPool() *BuildPool {
	return &BuildPool{_data:make(map[string]Build)}
}

func (k *BuildPool)Add(n string,t string,b BuildFunc) {
	k._data[n]= Build{t,b}
}

func (k *BuildPool)GetByName(s string) (string,Build,bool) {
	if b,ok:=k._data[s];ok {
		return s,b,true
	} else {
		return "", Build{},false
	}
}

func (k *BuildPool)GetByType(s string)  (string, Build,bool) {
	for lk,lv:=range k._data {
		if lv._type ==s {
			return lk,lv,true
		}
	}

	return "", Build{},false
}

var task_set *BuildPool

func init() {
	task_set = NewBuildPool()
	task_set.Add("HelloWorld",reflect.TypeOf((*HelloWorld)(nil)).String(),BuildTaskHelloWorldFromString)
}


func TaskFromString(s string) Task {
	i:=strings.IndexRune(s,' ')

	if i>0 {
		n:=s[:i]
		a:=s[i+1:]

		if _,b,ok:= task_set.GetByName(n);ok {
			return b._func(a)
		} else {
			panic("Can't build task from "+s)
		}
	} else {
		if _,b,ok:= task_set.GetByName(s);ok {
			return b._func("")
		} else {
			panic("Can't build task from "+s)
		}
	}


}

func TaskToString(t Task) string {
	if n,_,ok:= task_set.GetByType(reflect.TypeOf(t).String());ok {
		return n+" "+t.String()
	} else {
		panic("Can't find type "+n)
	}
}






