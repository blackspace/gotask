package main

import (
	"net/http"
	"log"
	"github.com/blackspace/gotask/runnable_pool/channel"
	"github.com/blackspace/gotask/examples/tasks"
)

var runnable_pool *channel.RunnablePoolWithChannel =channel.NewRunnablePoolWithChannel()

func init() {
	runnable_pool.Run()
}

func HelloWorldHandler(w http.ResponseWriter, req *http.Request) {
	s:=(<-runnable_pool.AddTask(tasks.NewHelloWorld())).(string)
	w.Write([]byte(s))
}

func ImageHandler(w http.ResponseWriter, req *http.Request) {
	buf:=(<-runnable_pool.AddTask(tasks.NewImage())).([]byte)
	w.Write(buf)
}

func main() {
	http.HandleFunc("/hello", HelloWorldHandler)
	http.HandleFunc("/image", ImageHandler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}