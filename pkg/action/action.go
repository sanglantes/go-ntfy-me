package action

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/sanglantes/go-ntfy-me/pkg/event"
)

type Action func(event.Event, event.EventBus)

type Registration struct {
	Name       string
	Action     Action
	Priority   int
	IsBlocking bool
}

type Registry struct {
	registryPriority map[int][]Registration
	registryName     map[string]Registration
	priorities       []int
	eventName        string
	eventBus         event.EventBus
}

func NewRegistry(eventName string, eb event.EventBus) *Registry {
	r := Registry{
		registryPriority: make(map[int][]Registration),
		registryName:     make(map[string]Registration),
		eventName:        eventName,
	}

	eb.Subscribe(eventName, r.RunAll)

	return &r
}

// Register registers actions to be fired when the ntfy.msg event is triggered.
func (r *Registry) Register(in Registration) {
	r.registryPriority[in.Priority] = append(r.registryPriority[in.Priority], in)
	r.registryName[in.Name] = in

	r.priorities = append(r.priorities, in.Priority)
	slices.SortFunc(r.priorities, func(a, b int) int { return cmp.Compare(b, a) })
}

func (r *Registry) Run(name string, e event.Event) error {
	handler, ok := r.registryName[name]
	if !ok {
		return fmt.Errorf("no such registry entry: %s", name)
	}
	if handler.IsBlocking {
		handler.Action(e, r.eventBus)
	} else {
		go handler.Action(e, r.eventBus)
	}

	return nil
}

func (r *Registry) RunAll(e event.Event) {
	for _, p := range r.priorities {
		for _, v := range r.registryPriority[p] {
			if v.IsBlocking {
				v.Action(e, r.eventBus)
			} else {
				go v.Action(e, r.eventBus)
			}
		}
		break
	}
}
