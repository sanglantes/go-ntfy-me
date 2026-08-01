package custom

import (
	"fmt"

	"github.com/sanglantes/go-ntfy-me/pkg/action"
	"github.com/sanglantes/go-ntfy-me/pkg/event"
	"github.com/sanglantes/go-ntfy-me/pkg/ntfy"
)

// This action echoes a received message.
// The event type must be of `message`. See https://docs.ntfy.sh/subscribe/api/#list-of-all-parameters

var ExportDefaultAction = action.Registration{
	Name:       "default",
	Action:     defaultAct,
	Priority:   1,
	IsBlocking: false,
}

func defaultAct(e event.Event, eb event.EventBus) {
	switch msgx := e.Data.(type) {
	case ntfy.NtfyMessage:
		printMsg(msgx)

	case []ntfy.NtfyMessage:
		for _, msg := range msgx {
			printMsg(msg)
		}
	}
}

func printMsg(msg ntfy.NtfyMessage) {
	if msg.Event == "message" {
		fmt.Println(msg.Message)
	}
}
