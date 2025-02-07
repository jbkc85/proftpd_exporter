package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

const EXPORTER_NAMESPACE = "proftpd"

var exporterServerCmd = &cobra.Command{
	Use:   "server",
	Short: fmt.Sprintf("Start %s_exporter server.", EXPORTER_NAMESPACE),
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {
	initializeConfig()
	collector := newProftpdCollector()
	prometheus.MustRegister(collector)
	router := mux.NewRouter()
	router.HandleFunc("/", healthHandler)
	router.Handle("/metrics", promhttp.Handler())

	log.Fatal(
		http.ListenAndServe(
			":"+flagPort,
			router,
		),
	)
}

type healthReturn struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(healthReturn{
		Status:  "ok",
		Message: "",
	})
}
