package main

import (
	"net/http"
	"log"
	"github.com/blackspace/gotask"
)


type HttpRequest struct {
	Request *http.Request
}

func (r *HttpRequest)Exec() interface{} {
	return "Hello World"
}

func NewHttpRequest() *HttpRequest {
	return &HttpRequest{}
}

var runnable_pool *gotask.RunnablePool =gotask.NewRunnablePool()

func Handler(w http.ResponseWriter, req *http.Request) {
	t:= NewHttpRequest()
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