package main

import (
	"github.com/blackspace/gotask"
	"time"
	"log"
)

var  event_loop *gotask.EventLoop


type TickEvent struct {
	Time time.Time
}

func init() {
	event_loop=gotask.NewEventLoop()
}


func main() {
	t:=time.NewTicker(time.Second)
	go func() {
		for {
			event_loop.AddEvent(TickEvent{<-t.C})
		}
	}()

	event_loop.GetOrCreateDelegate(TickEvent{}).AddHandler(func(e gotask.Event){
		log.Println(e.(TickEvent).Time)
	})

	event_loop.Run()
}
