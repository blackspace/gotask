package main

import (
	"github.com/blackspace/gotask/runnable_pool/callback"
	"github.com/blackspace/gotask/examples/tasks"
	"log"
)

var runnable_pool= callback.NewRunnablePoolWithCallback()

func init() {
	runnable_pool.Run()
}

func main() {
	for {
		runnable_pool.AddTask(tasks.NewHelloWorld(),func(r interface{}){
			log.Println(r)
		})
	}
}
