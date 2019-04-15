package store

import (
	"os"
)

// Config is the store configuration.
type Config struct {
	Path string `json:"path"`
}

// DefaultConfig is the default store configuration.
func DefaultConfig() *Config {
	return &Config{
		Path: os.TempDir(), // Default path in system tmp dir. TODO: create new directory?
	}
}
