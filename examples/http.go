package main

import (
	"io"
	"net/http"
	"log"
	"github.com/blackspace/gotask"
	"github.com/blackspace/gotask/hello_world"
)


var run_pool *gotask.TaskPool

// hello world, the web server
func Handler(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	t:=hello_world.NewHelloWorld()
	run_pool.AddTask(t)
	r:=t.ReceiveResult()
	io.WriteString(w, r.(string))
}

func init() {
	run_pool =gotask.NewTaskPool()
	run_pool.Run()
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}