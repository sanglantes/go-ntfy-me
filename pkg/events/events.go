package events

type Event struct {
	Type string
	Data any
}

type Handler func(Event)

type EventBus interface {
	Subscribe(string, Handler)
	Publish(Event)
}

type DefaultEventBus struct {
	handlers map[string][]Handler
}

func (eb *DefaultEventBus) Subscribe(eventType string, handler Handler) {
	eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

func (eb *DefaultEventBus) Publish(event Event) {
	for _, handler := range eb.handlers[event.Type] {
		go handler(event)
	}
}

func Default() *DefaultEventBus {
	return &DefaultEventBus{
		handlers: make(map[string][]Handler),
	}
}
