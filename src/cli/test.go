package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// CLI Start is to Start the cli in test mode
func (c *CLI) Test(file string) error {

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("SLO", "SLI", "Data", "SLO Status", "SLO Goal", "SLO Error Budget", "SLI Max", "SLI Mean", "SLI Min", "SLI Error")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Load the configuration
	config, err := c.Domain.LoadConfig(file)
	if err != nil {
		return err
	}

	// Open input and output connections
	err = c.Domain.LoadInputs(config.Inputs)
	if err != nil {
		return err
	}

	// Start the server
	results, err := c.Domain.Test(config)
	if err != nil {
		return err
	}

	// Add the results to the table
	for _, result := range results {
		tbl.AddRow(result.SLOName, result.Name, result.Data, result.Status, result.Goal, result.ErrorBudget, result.Max, result.Mean, result.Min, result.Error)
	}

	fmt.Println()
	tbl.Print()
	fmt.Println()
	return nil
}
