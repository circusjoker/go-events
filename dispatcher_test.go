package events

import (
	"fmt"
	"testing"
)

func TestDispatcher_listen(t *testing.T) {
	dispatcher := New()
	listener1 := &Listener1{}
	listener2 := &Listener2{}
	var listeners []ListenerContract
	listeners = append(listeners, listener1)
	listeners = append(listeners, listener2)
	dispatcher.Listen("test", listeners)

	var payload interface{}
	payload = "event test"
	_ = dispatcher.Dispatch("test", payload)

}

type Listener1 struct{}

func (listener1 *Listener1) Handle(payload interface{}) error {
	fmt.Println("listener1 got value", payload)
	return nil
}

type Listener2 struct{}

func (listener2 *Listener2) Handle(payload interface{}) error {
	fmt.Println("listener2 got value", payload)
	return nil
}
