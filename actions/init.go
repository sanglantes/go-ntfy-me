package actions

import (
	"github.com/sanglantes/go-ntfy-me/actions/custom"
	"github.com/sanglantes/go-ntfy-me/pkg/action"
	"github.com/sanglantes/go-ntfy-me/pkg/events"
)

func Install(r *action.Registry, eb events.EventBus) {
	// Extend me.

	r.Register(custom.ExportDefaultAction, nil)
}
