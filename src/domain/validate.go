package domain

import (
	"errors"
	"srenity/entities/pkl"
)

// Validate will validate the configuration is correct
func (d *Domain) Validate(config *pkl.Configuration) error {

	if config == nil {
		return errors.New("config is nil")
	}

	// Check if there are any inputs
	if len(config.Inputs) == 0 {
		return errors.New("no inputs found")
	}

	// Check if there are any outputs
	if len(config.Outputs) == 0 {
		return errors.New("no outputs found")
	}

	// Check if there are any SLO's
	if len(config.SLOs) == 0 {
		return errors.New("no SLO's found")
	}

	// Check if the SLO's have any SLI's
	for _, slo := range config.SLOs {
		if len(slo.SLIs) == 0 {
			return errors.New("No SLI's found for " + slo.Name)
		}

		// Check if the SLI's have any input
		for _, sli := range slo.SLIs {
			if sli.Input == "" {
				return errors.New("No input found for " + sli.Name)
			} else {
				// Check if the input exists
				found := false
				for _, input := range config.Inputs {
					if sli.Input == input.Name {
						found = true
					}
				}

				// check the SLI over
				if sli.Interval == nil {
					return errors.New("No interval found for " + sli.Name)
				}

				// check the SLI over

				if !found {
					return errors.New("Input " + sli.Input + " not found for " + sli.Name)
				}
			}
		}

		// Check if the SLO has an output
		if slo.Output == "" {
			return errors.New("No output found for " + slo.Name)
		} else {
			// Check if the output exists
			found := false
			for _, output := range config.Outputs {
				if slo.Output == output.Name {
					found = true
				}
			}

			if !found {
				return errors.New("Output " + slo.Output + " not found for " + slo.Name)
			}
		}
	}

	return nil
}
