package main

import (
	"github.com/blackspace/gotask/runnable_pool/channel"
	"github.com/blackspace/gotask/examples/tasks"
	"log"
)

var runnable_pool= channel.NewRunnablePoolWithChannel()

func init() {
	runnable_pool.Run()
}

func main() {
	for {
		c:=runnable_pool.AddTask(tasks.NewHelloWorld())
		log.Println((<-c).(string))
	}
}

