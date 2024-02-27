package domain

import (
	"math/big"
	"srenity/entities"
	"srenity/entities/pkl"
)

// Test is a test function of the server to output the SLI and SLO data
func (d *Domain) Test(config *pkl.Configuration) ([]entities.SLIResult, error) {

	var results []entities.SLIResult

	// in the test mode, we don't need to start the server or wait forever
	// we just need to output the SLI and SLO data
	for _, slo := range config.SLOs {
		for _, sli := range slo.SLIs {
			// Get the SLI data
			SLOmet, SLIMax, SLIMean, SLIMin, data, SLIError, err := d.CalculateSLI(sli)
			if err != nil {
				return []entities.SLIResult{}, err
			}

			// calculate the error budget
			errorBudget := new(big.Float).Sub(big.NewFloat(100.0), big.NewFloat(sli.Goal))

			// add the data to the results
			results = append(results, entities.SLIResult{
				SLOName:     slo.Name,
				Name:        sli.Name,
				Data:        data,
				Status:      d.SLOMettoString(SLOmet),
				Goal:        sli.Goal,
				ErrorBudget: errorBudget.String(),
				Max:         SLIMax,
				Mean:        SLIMean,
				Min:         SLIMin,
				Error:       SLIError,
			})
		}
	}

	return results, nil
}

func (d *Domain) SLOMettoString(SLOMet bool) string {
	if SLOMet {
		return "Passing"
	}
	return "Failing"
}
