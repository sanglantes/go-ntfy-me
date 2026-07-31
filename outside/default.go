package outside

import (
	"fmt"

	"github.com/sanglantes/go-ntfy-me/pkg/actions"
	"github.com/sanglantes/go-ntfy-me/pkg/events"
	"github.com/sanglantes/go-ntfy-me/pkg/ntfy"
)

// This action echoes a received message.
// The event type must be of `message`. See https://docs.ntfy.sh/subscribe/api/#list-of-all-parameters

var defaultActionStruct = actions.Registration{
	Name:       "default",
	Action:     defaultAct,
	Priority:   1,
	IsBlocking: false,
}

func defaultAct(e events.Event) {
	if e.Data.(ntfy.NtfyMessage).Event == "message" {
		fmt.Println(e.Data.(ntfy.NtfyMessage).Message)
	}
}
