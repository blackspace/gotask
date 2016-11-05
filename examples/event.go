package main

import (
	"github.com/blackspace/gotask"
	"time"
	"log"
)

var  event_loop *gotask.EventLoop


type TickEvent struct {
	gotask.Event
	Time time.Time
}

func (t TickEvent)Fire() {
	event_loop.AddEvent(TickEvent{Time:time.Now()})
}


type MyTick struct {
	Tick *time.Ticker
	TickEvent TickEvent
}

func NewMyTick() *MyTick {
	return &MyTick{Tick:time.NewTicker(time.Second)}
}

func init() {
	event_loop=gotask.NewEventLoop()
}


func main() {
	t:=NewMyTick()
	go func() {
		for {
			_=<-t.Tick.C
			t.TickEvent.Fire()
		}
	}()

	event_loop.GetOrCreateDelegate(TickEvent{}).AddHandler(func(e gotask.Event){
		log.Println(e.(TickEvent).Time)
	})

	event_loop.Run()
}
