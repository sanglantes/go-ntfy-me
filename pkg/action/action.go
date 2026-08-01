package action

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/sanglantes/go-ntfy-me/pkg/events"
)

type Action func(events.Event)

type Registration struct {
	Name       string
	Action     Action
	Priority   int
	IsBlocking bool
	EventBus   events.EventBus
}

type Registry struct {
	registryPriority map[int][]Registration
	registryName     map[string]Registration
	priorities       []int
	eventName        string
}

func NewRegistry(eventName string, eb events.EventBus) *Registry {
	r := Registry{
		registryPriority: make(map[int][]Registration),
		registryName:     make(map[string]Registration),
		eventName:        eventName,
	}

	eb.Subscribe(eventName, r.RunAll)

	return &r
}

// Register registers actions to be fired when the ntfy.msg event is triggered.
func (r *Registry) Register(in Registration, eb events.EventBus) {
	r.registryPriority[in.Priority] = append(r.registryPriority[in.Priority], in)
	r.registryName[in.Name] = in

	r.priorities = append(r.priorities, in.Priority)
	slices.SortFunc(r.priorities, func(a, b int) int { return cmp.Compare(b, a) })

	in.EventBus = eb
}

func (r *Registry) Run(name string, event events.Event) error {
	handler, ok := r.registryName[name]
	if !ok {
		return fmt.Errorf("no such registry entry: %s", name)
	}
	if handler.IsBlocking {
		handler.Action(event)
	} else {
		go handler.Action(event)
	}

	return nil
}

func (r *Registry) RunAll(event events.Event) {
	for _, p := range r.priorities {
		for _, v := range r.registryPriority[p] {
			if v.IsBlocking {
				v.Action(event)
			} else {
				go v.Action(event)
			}
		}
	}
}
