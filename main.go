package main

import (
	"io"
	"net/http"
	"log"
	"github.com/blackspace/gotasks/task"
)


var task_channel chan task.Task

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	t:=factory.TaskFromString("HelloWorld")
	task_channel <-t
	r:=t.ReceiveResult()
	io.WriteString(w, r.(string))
}

var factory * task.Factory

func init() {
	factory = task.NewFactory()
	factory.AddCreatable("HelloWorld",task.HelloWorldCreatable)


	task_channel =make(chan task.Task,1<<8)

	go func(){
		for {
			t:= <-task_channel
			r:=t.Exec()
			t.SendResult(r)
		}
	}()
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}