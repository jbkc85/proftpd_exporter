package main

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Verbose bool `json:"verbose" yaml:"verbose"`
}

var (
	exporterConfig config

	flagPort    string
	flagVerbose bool
)

func init() {
	proftpdExporter.PersistentFlags().StringVarP(&flagPort, "port", "p", "41337", "Port for vuls_exporter to bind to and expose openmetrics")
	proftpdExporter.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "Enable Verbose mode for log output")
	proftpdExporter.AddCommand(exporterServerCmd)
}

func initializeConfig() {
	//registerCVEMetrics()
	//registerResultsMetrics()
	// configure viper to look into the following directories for
	// a 'config.yaml' configuration file.
	// 1. $HOME/.kafkactl, 2. /etc/kafkactl, 3. $PWD
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.proftpd_exporter")
	viper.AddConfigPath("/opt/proftpd_exporter")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		log.Printf("[DEBUG] Unable to find configuration file. (message: %s)", err)
	} else {
		err = viper.Unmarshal(&exporterConfig)
		if err != nil {
			if flagVerbose {
				log.Fatalf("[FATAL] Configuration file found but unable to parse (err :%s)", err)
			}
		}
	}
}
