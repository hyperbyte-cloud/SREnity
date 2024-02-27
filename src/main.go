package main

import (
	"log/slog"

	"srenity/cli"
	"srenity/domain"
)

var version string

func main() {
	slog.Info("Starting SREniy the SLO Tool")

	// init the domain
	domain := domain.NewDomain()

	// init the CLI
	cli := cli.NewCLI(
		domain,
		version,
	)

	// run the CLI
	cli.Run()
}
