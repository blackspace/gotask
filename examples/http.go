package main

import (
	"net/http"
	"log"
	"github.com/blackspace/gotask"
)


type HelloWorld struct {}

func (r *HelloWorld)Exec() interface{} {
	return "Hello World"
}

func NewHelloWorld() *HelloWorld {
	return &HelloWorld{}
}

var runnable_pool *gotask.RunnablePool =gotask.NewRunnablePool()

func Handler(w http.ResponseWriter, req *http.Request) {
	s:=(<-runnable_pool.AddTask(NewHelloWorld())).(string)
	w.Write([]byte(s))
}

func init() {
	runnable_pool.Run()
}

func main() {
	http.HandleFunc("/hello", Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}