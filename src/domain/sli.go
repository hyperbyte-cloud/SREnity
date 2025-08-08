package domain

import (
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"srenity/entities"
	"srenity/entities/pkl"
	"time"
)

// StartMonitoring is to start the monitoring
func (d *Domain) MonitorSLI(sli *pkl.SLI) error {
	slog.Info("Starting to monitor: " + sli.Name)

	// Check the d.inputDatasources are there
	if _, ok := d.inputDatasources[sli.Input]; !ok {
		return errors.New("input " + sli.Input + " not found")
	}

	go func() {
		// Start the ticker
		ticker := time.NewTicker(sli.Interval.GoDuration())

		for range ticker.C {
			// Calculate the SLI
			SLOmet, SLIMax, SLIMean, SLIMin, _, SLIError, err := d.CalculateSLI(sli)
			if err != nil {
				slog.Error(err.Error())
				continue
			}

			if SLIError != "" {
				slog.Error(SLIError)
				continue
			}

			errorBudget := new(big.Float).Sub(big.NewFloat(100.0), big.NewFloat(sli.Goal))
			errorBudgetFloat, _ := errorBudget.Float64()

			// Create a new metric
			metric := entities.WriteMetric{
				Timestamp: time.Now(),
				Name:      "slo_output",
				Tags: map[string]string{
					"sli_name": sli.Name,
				},
				Values: map[string]interface{}{
					"state":      SLOmet,
					"goal":       sli.Goal,
					"err_budget": errorBudgetFloat,
					"sli_max":    SLIMax,
					"sli_min":    SLIMin,
					"sli_mean":   SLIMean,
				},
			}

			d.metricChan <- metric
		}
	}()

	// Start the monitoring
	return nil
}

// CalculateSLI is to calculate the SLI
// This function returns the SLI met, the max, mean, min, data, query error and go-error
func (d *Domain) CalculateSLI(sli *pkl.SLI) (bool, float64, float64, float64, []float64, string, error) {
	slog.Debug("Querying the database for: " + sli.Name)
	// Ensure the input exists
	repo, ok := d.inputDatasources[sli.Input]
	if !ok {
		return false, 0, 0, 0, nil, "", errors.New("input " + sli.Input + " not found")
	}
	// Query the database
	QueryMetric, err := repo.Query(sli.Query)
	if err != nil {
		return false, 0, 0, 0, nil, "", err
	}

	var met bool
	var max, mean, min float64
	var data []float64
	var queryError string
	if len(QueryMetric.Values) == 0 {
		queryError = "No values returned from the query"
		met = false
	} else {
		// initialize min to first value to compute proper minimum
		min = QueryMetric.Values[0]
		for _, value := range QueryMetric.Values {
			// Check to see if the SLI goal is met
			// If this value exceeds the goal, this sample fails
			if value > sli.Goal {
				met = false
			}
			if value > max {
				max = value
			}
			if value < min {
				min = value
			}
			mean += value
			data = append(data, value)
		}

		// Calculate the mean
		mean = mean / float64(len(QueryMetric.Values))
		// does this meet the SLO?
		met = mean <= sli.Goal
	}

	// log the SLOmet
	if !met {
		slog.Warn(fmt.Sprintf("SLO not met for %s", sli.Name))
	}

	return met, max, mean, min, data, queryError, nil
}
