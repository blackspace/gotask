package test


import (
	"github.com/blackspace/gotasks/task"
	"testing"
)

func TestHelloWorldFromString(t *testing.T) {
	f := task.NewFactory()
	f.AddCreatable("HelloWorld",task.HelloWorldCreatable)

	task:=f.TaskFromString("HelloWorld")
	r:=task.Exec()

	if r.(string)!="Hello World" {
		t.Fail()
	}

}
