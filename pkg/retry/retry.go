package retry

import (
	"fmt"
	"time"
)

type Config struct {
	MaxRetries int
	Delay      time.Duration
}

func DefaultConfig() Config {
	return Config{
		MaxRetries: 3,
		Delay:      1 * time.Second,
	}
}

func Retry(config Config, opr func() error) error {

	var lastErr error

	for i := 1; i <= config.MaxRetries; i++ {
		err := opr()
		if err == nil {
			return nil
		}

		lastErr = err

		if i == config.MaxRetries {
			break
		}

		fmt.Printf("Attempt %d failed: %v. Retrying in %v...\n", i, err, config.Delay)
		time.Sleep(config.Delay)
	}
	return fmt.Errorf("failed after %d attempts: %w", config.MaxRetries, lastErr)
}
