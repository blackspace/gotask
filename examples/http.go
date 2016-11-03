package main

import (
	"io"
	"net/http"
	"log"
	"github.com/blackspace/gotask"
	"github.com/blackspace/gotask/dohttp"
)


var run_pool *gotask.TaskPool

func Handler(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	t:=dohttp.NewDoRequest()
	t.Request=req
	io.WriteString(w, (<-run_pool.AddTask(t)).(string))
}

func init() {
	run_pool =gotask.NewTaskPool()
	run_pool.Run()
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}