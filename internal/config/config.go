package config

import (
	"fmt"
	"os"
	"time"
)

// Config holds all runtime configuration for Gauntlet.
type Config struct {
	HTTP HTTPConfig

	NVIDIA NVIDIAConfig
}

type HTTPConfig struct {
	Address string
	Timeout time.Duration
}

type NVIDIAConfig struct {
	APIKey  string
	BaseURL string
}

// Load reads configuration from the environment.
func Load() (*Config, error) {
	cfg := &Config{
		HTTP: HTTPConfig{
			Address: ":8080",
			Timeout: 120 * time.Second,
		},
		NVIDIA: NVIDIAConfig{
			BaseURL: "https://integrate.api.nvidia.com/v1",
			APIKey:  os.Getenv("NVIDIA_API_KEY"),
		},
	}

	if cfg.NVIDIA.APIKey == "" {
		return nil, fmt.Errorf("NVIDIA_API_KEY is not set")
	}

	return cfg, nil
}
