package connection

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sanglantes/go-ntfy-me/internal/config"
	"github.com/sanglantes/go-ntfy-me/pkg/action"
	"github.com/sanglantes/go-ntfy-me/pkg/event"
	"github.com/sanglantes/go-ntfy-me/pkg/ntfy"
)

func Start(cfg *config.Config, eb event.EventBus, r *action.Registry) {
	for true {
		err := listen(cfg, eb)

		if err != nil {
			log.Printf("connection error: %v. retrying in %d seconds...", err, cfg.RetryTime)
			eb.Publish(event.Event{
				Type: "self.conn_error",
				Data: err,
			})
		}

		time.Sleep(time.Second * time.Duration(cfg.RetryTime))
	}
}

func listen(cfg *config.Config, eb event.EventBus) error {
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received HTTP status %s", resp.Status)
	}

	if cfg.Polling {
		var msgs []ntfy.NtfyMessage

		decoder := json.NewDecoder(resp.Body)
		for decoder.More() {
			var msg ntfy.NtfyMessage
			if err := decoder.Decode(&msg); err != nil {
				return err
			}
			msgs = append(msgs, msg)
		}

		eb.Publish(event.Event{
			Type: "ntfy.msg",
			Data: msgs,
		})
	} else {
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			msg := ntfy.NtfyMessage{}
			if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
				return err
			}

			eb.Publish(event.Event{
				Type: "ntfy.msg",
				Data: msg,
			})
		}

		return scanner.Err()
	}

	return nil
}
