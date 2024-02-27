package cli

import (
	"log/slog"
)

// Validate the configuration
func (c *CLI) Validate(file string) error {
	// Validate the configuration
	_, err := c.Domain.LoadConfig(file)
	if err != nil {
		return err
	}

	slog.Info("Configuration is valid")

	return nil
}
