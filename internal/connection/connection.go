package connection

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sanglantes/go-ntfy-me/internal/config"
	"github.com/sanglantes/go-ntfy-me/pkg/action"
	"github.com/sanglantes/go-ntfy-me/pkg/events"
	"github.com/sanglantes/go-ntfy-me/pkg/ntfy"
)

func Start(cfg *config.Config, eb events.EventBus, r *action.Registry) {
	for true {
		err := listen(cfg, eb)

		if err != nil {
			log.Printf("connection error: %v. retrying in %d seconds...", err, cfg.RetryTime)
			eb.Publish(events.Event{
				Type: "self.conn_error",
				Data: err,
			})
		}

		time.Sleep(time.Second * time.Duration(cfg.RetryTime))
	}
}

func listen(cfg *config.Config, eb events.EventBus) error {
	req, err := http.NewRequest(http.MethodGet, cfg.Endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+cfg.AuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if cfg.Polling {
		var respBody []byte
		if _, err := resp.Body.Read(respBody); err != nil {
			return err
		}

		msg := ntfy.NtfyMessage{}
		if err := json.Unmarshal(respBody, &msg); err != nil {
			return err
		}

		eb.Publish(events.Event{
			Type: "ntfy.msg",
			Data: respBody,
		})
	} else {
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			msg := ntfy.NtfyMessage{}
			if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
				return err
			}

			eb.Publish(events.Event{
				Type: "ntfy.msg",
				Data: msg,
			})
		}

		return scanner.Err()
	}

	return nil
}
