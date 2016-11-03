package main

import (
	"io"
	"net/http"
	"log"
	"github.com/blackspace/gotask"
	"github.com/blackspace/gotask/dohttp"
)


var runable_pool *gotask.RunablePool=gotask.NewRunablePool()

func Handler(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	t:=dohttp.NewDoRequest()
	t.Request=req
	io.WriteString(w, (<-runable_pool.AddTask(t)).(string))
}

func init() {
	runable_pool.Run()
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}