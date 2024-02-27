package domain

import (
	"log/slog"
	"srenity/entities/pkl"
)

// Start the server
func (d *Domain) Server(config *pkl.Configuration) error {
	slog.Info("Starting the server")

	// Open input and output connections
	err := d.LoadInputs(config.Inputs)
	if err != nil {
		return err
	}

	// Open input and output connections
	err = d.LoadOutputs(config.Outputs)
	if err != nil {
		return err
	}

	// Start monitoring the SLI's in go routines
	for _, slo := range config.SLOs {
		// Start the monitoring
		err := d.MonitorSLO(slo)
		if err != nil {
			return err
		}
	}

	return nil
}
