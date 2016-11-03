package hello_world

import (
	"reflect"
	. "github.com/blackspace/gotask"
)

var HelloWorldCreatable = Creatable{reflect.TypeOf((*HelloWorld)(nil)).String(), HelloWorldFromString}

type HelloWorld struct {
	*TaskBase

}

func (pl *HelloWorld)Exec() interface{} {
	return "Hello World"
}

func (pl *HelloWorld)String() string {
	return "Hello World"
}

func HelloWorldFromString(s string) Task {
	t:=NewHelloWorld()
	return t
}


func NewHelloWorld() *HelloWorld  {
	return &HelloWorld {TaskBase:NewTaskBase()}
}

