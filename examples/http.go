package main

import (
	"io"
	"net/http"
	"log"
	"github.com/blackspace/gotask"
)


var task_pool *gotask.TaskPool

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	t:=gotask.NewHelloWorld()
	task_pool.AddTask(t)
	r:=t.ReceiveResult()
	io.WriteString(w, r.(string))
}

func init() {
	task_pool =gotask.NewTaskPool()

	task_pool.Run()
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}