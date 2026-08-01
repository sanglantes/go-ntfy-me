//go:build !windows

package config

import (
	"fmt"
)

func AddToStartUp(cfg *Config) error {
	return fmt.Errorf("unimplemented")
}
