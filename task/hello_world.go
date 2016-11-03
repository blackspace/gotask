package task

import "reflect"

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

var HelloWorldCreatable = Creatable{reflect.TypeOf((*HelloWorld)(nil)).String(), HelloWorldFromString}