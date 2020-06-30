package events

type ListenerContract interface {
	Handle(event string, args ...interface{}) error
}

type DispatcherContract interface {
	/**
	 * Register an event listener with the dispatcher.
	 */
	Listen(event string, listener ListenerContract)

	/**
	 * Dispatch an event and call the listeners.
	 */
	Dispatch(event string, args ...interface{}) error
}


type Dispatcher struct {
	listeners map[string][]ListenerContract
}

func New() *Dispatcher {
	listeners := make(map[string][]ListenerContract)
	return &Dispatcher{
		listeners: listeners,
	}
}

// Register an event listener with the dispatcher.
func (dispatcher *Dispatcher) Listen(event string, listeners []ListenerContract) {
	for _, listener := range listeners { //nolint:gosimple
		dispatcher.listeners[event] = append(dispatcher.listeners[event], listener)
	}
}

// Dispatch an event and call the listeners.
func (dispatcher *Dispatcher) Dispatch(event string, args ...interface{}) error {
	if listeners, ok := dispatcher.listeners[event]; ok {
		for _, listener := range listeners {
			err := listener.Handle(event, args...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
