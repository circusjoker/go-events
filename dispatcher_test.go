package events

import (
	"fmt"
	"testing"
)

func TestDispatcher_listen(t *testing.T) {
	dispatcher := New()

	dispatcher.Listen("test", []ListenerContract{&Listener1{}, &Listener2{}})

	dispatcher.Dispatch("test", "event test")
}

type Listener1 struct{}

func (listener1 *Listener1) Handle(event string, payload interface{}) error {
	fmt.Println("event", event, ":", "listener1", payload)
	return nil
}

type Listener2 struct{}

func (listener2 *Listener2) Handle(event string, payload interface{}) error {
	fmt.Println("event", event, ":", "listener2", payload)
	return nil
}
