package main

import (
	"github.com/blackspace/goevent"
	"github.com/blackspace/gotask/runnable_pool/event"
	"github.com/blackspace/gotask/examples/tasks"
	"log"
)

var runnable_pool=event.NewRunnablePoolEvent()

func init() {
	runnable_pool.Run()
}

func main() {
	runnable_pool.TaskDoneEvent.AddHandler(func(s goevent.Source,a goevent.EventArg){
		switch s.(type) {
		case *tasks.HelloWorld:
			log.Println(a.(string))
		}

	})

	go func() {
		for {
			runnable_pool.AddTask(tasks.NewHelloWorld())
		}
	}()


	goevent.Run()
}
