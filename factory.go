package gotask

import (
	"strings"
	"reflect"
)

type CreateFunc func (s string) Task

type Creatable struct{
	_type string
	_func CreateFunc
}

type Factory struct {
	_data map[string]Creatable
}

func NewFactory() *Factory {
	return &Factory{_data:make(map[string]Creatable)}
}

func (f *Factory)AddCreatable(n string,c Creatable) {
	f._data[n]= c
}

func (f *Factory)GetCreatableByName(s string) (string, Creatable,bool) {
	if b,ok:= f._data[s];ok {
		return s,b,true
	} else {
		return "", Creatable{},false
	}
}

func (f *Factory)GetCreatableByType(s string)  (string, Creatable,bool) {
	for lk,lv:=range f._data {
		if lv._type ==s {
			return lk,lv,true
		}
	}

	return "", Creatable{},false
}

func (f *Factory)TaskFromString(s string) Task {
	i:=strings.IndexRune(s,' ')

	if i>0 {
		n:=s[:i]
		a:=s[i+1:]

		if _,b,ok:= f.GetCreatableByName(n);ok {
			return b._func(a)
		} else {
			panic("Can't build task from "+s)
		}
	} else {
		if _,b,ok:= f.GetCreatableByName(s);ok {
			return b._func("")
		} else {
			panic("Can't build task from "+s)
		}
	}


}

func (f *Factory)TaskToString(t Task) string {
	if n,_,ok:= f.GetCreatableByType(reflect.TypeOf(t).String());ok {
		return n+" "+t.String()
	} else {
		panic("Can't find type "+n)
	}
}










