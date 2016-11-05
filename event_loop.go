package gotask

import (
	"reflect"
	"fmt"
)

type Event interface{}
type Handler func(Event)

type Delegate struct {
	_object   interface{}
	_handlers []Handler
}

func NewDelegate(o interface{}) *Delegate {
	return &Delegate{o,make([]Handler,0,1<<8)}
}

func (d *Delegate)AddHandler(h Handler) {
	d._handlers =append(d._handlers,h)
}

func (d *Delegate)Exec(e Event) {
	for _,h:=range d._handlers {
		h(e)
	}
}

type EventLoop struct {
	_channel  chan Event
	_delegates []*Delegate
}

func NewEventLoop() *EventLoop {
	return &EventLoop{_channel:make(chan Event,1<<10), _delegates:make([]*Delegate,0,1<<8)}
}


func (el *EventLoop)GetDelegate(o interface{}) *Delegate {
	var d *Delegate

	for i:=0;i<len(el._delegates);i++ {
		if reflect.TypeOf(el._delegates[i]._object) == reflect.TypeOf(o) {
			d=el._delegates[i]
			break
		}
	}


	if d==nil {
		return nil
	} else {
		return d
	}
}

func (el *EventLoop)CreateDelegate(o interface{}) *Delegate {
	d:=el.GetDelegate(o)

	if d==nil {
		d=NewDelegate(o)
		el._delegates=append(el._delegates,d)
		return d
	} else {
		panic(fmt.Sprintf("This delegate of %v is existing",o))
	}
}

func (el *EventLoop)GetOrCreateDelegate(o interface{}) *Delegate {
	d:=el.GetDelegate(o)

	if d==nil {
		d=NewDelegate(o)
		el._delegates=append(el._delegates,d)
		return d
	} else {
		return  d
	}
}


func (el *EventLoop)AddEvent(e Event) {
	el._channel<-e
}

func (el *EventLoop)Run() {
	for {
		e:=<-el._channel

		var d *Delegate =el.GetDelegate(e)

		if d!=nil {
			d.Exec(e)
		} else {
			panic("Can't find a handler for the event of "+reflect.TypeOf(e).String()+".")
		}
	}
}


