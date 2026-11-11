package config

import (
	"fmt"
	"os"
	"strconv"
)

const defaultMaxLineBytes = 1024 * 1024

// Config contains runtime settings for the application.
type Config struct {
	MaxLineBytes int
}

// Load returns configuration from environment variables and applies defaults.
func Load() (Config, error) {
	maxLineBytes := defaultMaxLineBytes
	if value, ok := os.LookupEnv("SECRET_DETECTOR_MAX_LINE_BYTES"); ok {
		parsed, err := strconv.Atoi(value)
		if err != nil || parsed <= 0 {
			return Config{}, fmt.Errorf("SECRET_DETECTOR_MAX_LINE_BYTES must be a positive integer")
		}
		maxLineBytes = parsed
	}
	return Config{MaxLineBytes: maxLineBytes}, nil
}
