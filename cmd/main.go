package main

import (
	"log"

	"github.com/sanglantes/go-ntfy-me/actions"
	"github.com/sanglantes/go-ntfy-me/internal/config"
	"github.com/sanglantes/go-ntfy-me/internal/connection"
	"github.com/sanglantes/go-ntfy-me/pkg/action"
	"github.com/sanglantes/go-ntfy-me/pkg/events"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to read configuration: %v", err)
	}
	log.Println("Loaded configuration.")
	log.Printf("%+v", cfg)

	eb := events.Default()
	log.Println("Created event bus.")

	registry := action.NewRegistry("ntfy.msg", eb)
	actions.Install(registry, eb)
	log.Println("Created action registry.")

	if cfg.AddToStartUp {
		if err := config.AddToStartUp(&cfg); err != nil {
			log.Printf("failed to add program to auto-start: %v", err)
		}
		log.Println("Added program to start up.")
	}

	eb.Publish(events.Event{
		Type: "self.start",
		Data: &cfg,
	})

	log.Printf("Starting listener...")

	connection.Start(&cfg, eb, registry)
}
