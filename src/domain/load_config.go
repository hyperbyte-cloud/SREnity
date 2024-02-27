package domain

import (
	"context"
	"log/slog"
	"srenity/entities/pkl"
)

// LoadConfig is to load the configuration
func (d *Domain) LoadConfig(file string) (*pkl.Configuration, error) {
	slog.Info("Loading the configuration " + file)

	ctx := context.Background()

	pklConfig, err := pkl.LoadFromPath(ctx, file)
	if err != nil {
		return &pkl.Configuration{}, err
	}

	// Validate the configuration
	err = d.Validate(pklConfig)
	if err != nil {
		return &pkl.Configuration{}, err
	}

	return pklConfig, nil
}
