package task

type HelloWorld struct {
	*TaskBase

}

func (pl *HelloWorld)Exec() interface{} {
	return "Hello World"
}

func (pl *HelloWorld)String() string {
	return "Hello World"
}

func BuildTaskHelloWorldFromString(s string) Task {
	t:=NewHelloWorld()
	return t
}


func NewHelloWorld() *HelloWorld  {
	return &HelloWorld {TaskBase:NewTaskBase()}
}