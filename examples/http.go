package main

import (
	"net/http"
	"log"
	"github.com/blackspace/gotask"
)


type TaskHelloWorld struct {
	Request *http.Request
}

func (r *TaskHelloWorld)Exec() interface{} {
	return "Hello World"
}

func NewTaskHelloWorld() *TaskHelloWorld {
	return &TaskHelloWorld{}
}

var runnable_pool *gotask.RunnablePool =gotask.NewRunnablePool()

func Handler(w http.ResponseWriter, req *http.Request) {
	t:= NewTaskHelloWorld()
	t.Request=req
	s:=(<-runnable_pool.AddTask(t)).(string)
	w.Write([]byte(s))
}

func init() {
	runnable_pool.Run()
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}