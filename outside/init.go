package outside

import (
	"github.com/sanglantes/go-ntfy-me/pkg/actions"
	"github.com/sanglantes/go-ntfy-me/pkg/events"
)

func Install(r *actions.Registry, eb events.EventBus) {
	// Extend me.

	r.Register(defaultActionStruct)
}
