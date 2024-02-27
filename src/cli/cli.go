package cli

import (
	"log/slog"
	"os"
	"sort"
	"srenity/domain"

	"github.com/urfave/cli/v2"
)

// CLI is the command line interface
type CLI struct {
	Domain     *domain.Domain
	AppVersion string
}

// NewCLI creates a new CLI
func NewCLI(domain *domain.Domain, version string) *CLI {
	return &CLI{
		Domain:     domain,
		AppVersion: version,
	}
}

// Run the CLI
func (c *CLI) Run() {
	app := &cli.App{
		Name:    "srenity",
		Usage:   "SREniy the SLO Tool",
		Version: c.AppVersion,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Value:    "config.json",
				Usage:    "Load configuration from `FILE`",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"r"},
				Usage:   "Run the SLO tool",
				Action: func(ctx *cli.Context) error {
					// Start the server
					// with the configuration file
					file := ctx.String("config")
					err := c.Start(file)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "validate",
				Aliases: []string{"v"},
				Usage:   "Validate the configuration",
				Action: func(ctx *cli.Context) error {
					// Validate the configuration
					file := ctx.String("config")
					err := c.Validate(file)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Test the configuration",
				Action: func(ctx *cli.Context) error {
					// Test the configuration
					file := ctx.String("config")
					err := c.Test(file)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		slog.Error(err.Error())
	}
}
