package domain

import (
	"log/slog"
	"srenity/entities"
	"srenity/entities/pkl"
)

// Monitor is to monitor the SLI's
func (d *Domain) MonitorSLO(slo *pkl.SLO) error {
	slog.Info("Monitoring the SLO: " + slo.Name)

	// open the metrics channel
	d.metricChan = make(chan entities.WriteMetric)

	// Start the monitoring
	for _, sli := range slo.SLIs {
		// Start the monitoring
		err := d.MonitorSLI(sli)
		if err != nil {
			return err
		}
	}

	// Start the writer
	go d.StartOutputWriter(slo)

	return nil
}
