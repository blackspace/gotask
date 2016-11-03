package test


import (
	"github.com/blackspace/gotasks"
	"testing"
)

func TestHelloWorldFromString(t *testing.T) {
	f := gotask.NewFactory()
	f.AddCreatable("HelloWorld",gotask.HelloWorldCreatable)

	task:=f.TaskFromString("HelloWorld")
	r:=task.Exec()

	if r.(string)!="Hello World" {
		t.Fail()
	}

}
