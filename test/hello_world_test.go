package test


import (
	. "github.com/blackspace/gotasks/task"
	"testing"
)

func TestHelloWorldFromString(t *testing.T) {
	task:=TaskFromString("HelloWorld")
	r:=task.Exec()

	if r.(string)!="Hello World" {
		t.Fail()
	}

}
