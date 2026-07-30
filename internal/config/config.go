package config

import _ "github.com/joho/godotenv"

type Protocol int

const (
	WS Protocol = iota
	HTTP
	POLL
)

type Config struct {
	Endpoint     string
	Protocol     Protocol
	Timeout      uint32
	AuthToken    string
	PollTime     uint32
	AddToStartup bool
}
