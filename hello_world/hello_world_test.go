package hello_world


import (
	"testing"
	"github.com/blackspace/gotask"
)

func TestHelloWorldFromString(t *testing.T) {
	f := gotask.NewFactory()
	f.AddCreatable("HelloWorld",HelloWorldCreatable)

	task:=f.TaskFromString("HelloWorld")
	r:=task.Exec()

	if r.(string)!="Hello World" {
		t.Fail()
	}

}
