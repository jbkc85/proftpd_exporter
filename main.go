package main

import (
	"log"

	"github.com/spf13/cobra"
)

var proftpdExporter = &cobra.Command{
	Use:           "proftpd_exporter",
	Short:         "exporter to provide an interface between ProFTPd Server Scoreboard and Prometheus.",
	Long:          `proftpd_exporter exposes generated 'reports' as metrics to Prometheus for aggregation and compact viewing purposes.`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func main() {
	if err := proftpdExporter.Execute(); err != nil {
		log.Fatal(err)
	}
}
