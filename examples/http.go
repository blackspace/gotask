package main

import (
	"io"
	"net/http"
	"log"
	"github.com/blackspace/gotask"
	"github.com/blackspace/gotask/examples/task_http"
)


var runnable_pool *gotask.RunnablePool =gotask.NewRunnablePool()

func Handler(w http.ResponseWriter, req *http.Request) {
	t:= task_http.NewHttpRequest()
	t.Request=req
	io.WriteString(w, (<-runnable_pool.AddTask(t)).(string))
}

func init() {
	runnable_pool.Run()
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}