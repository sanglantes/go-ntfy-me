//go:build windows

package config

import (
	"fmt"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func AddToStartUp(cfg *Config) error {
	hd, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	startupFolder := path.Join(hd, "AppData", "Roaming", "Microsoft",
		"Windows", "Start Menu", "Programs", "Startup")
	if err := os.MkdirAll(startupFolder, 0); err != nil {
		return err
	}

	e, err := os.Executable()
	if err != nil {
		return err
	}
	eBytes, err := os.ReadFile(e)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path.Join(startupFolder, "go-ntfy-me.exe"), eBytes, 0); err != nil {
		return err
	}

	return godotenv.Write(map[string]string{
		"NTFY_ENDPOINT":         cfg.Endpoint,
		"NTFY_AUTH_TOKEN":       cfg.AuthToken,
		"ADD_TO_START_UP":       fmt.Sprintf("%t", cfg.AddToStartUp),
		"CONNECTION_RETRY_TIME": fmt.Sprintf("%d", cfg.RetryTime),
	}, path.Join(startupFolder, ".env"))
}
