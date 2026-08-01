package actions

import (
	"github.com/sanglantes/go-ntfy-me/actions/custom"
	"github.com/sanglantes/go-ntfy-me/pkg/action"
)

func Install(r *action.Registry) {
	// Extend me.

	r.Register(custom.ExportDefaultAction)
}
