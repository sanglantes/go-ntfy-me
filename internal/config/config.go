package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Endpoint     string
	AuthToken    string
	Polling      bool
	AddToStartUp bool
	RetryTime    uint32
}

func New() (Config, error) {
	c := Config{}

	// Endpoint
	endpoint := os.Getenv("NTFY_ENDPOINT")
	if endpoint == "" {
		return c, fmt.Errorf("no endpoint specified")
	}
	c.Endpoint = endpoint

	// ListenType
	if !strings.HasPrefix(endpoint, "https") {
		return c, fmt.Errorf("invalid endpoint %s. endpoints must begin with http(s) and be a valid URI", endpoint)
	}
	if strings.Contains(endpoint, "poll=1") {
		c.Polling = true
	}

	// Auth token
	c.AuthToken = os.Getenv("NTFY_AUTH_TOKEN")

	// Add program to Windows start up.
	c.AddToStartUp = (os.Getenv("ADD_TO_START_UP") == "true")

	// Seconds before retrying a lost connection.
	rt, err := strconv.ParseUint(os.Getenv("CONNECTION_RETRY_TIME"), 10, 32)
	if err != nil {
		rt = 10
	}
	c.RetryTime = uint32(rt)

	return c, nil
}
