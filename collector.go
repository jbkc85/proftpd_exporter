package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type proftpdCollector struct {
	proftpdConnections *prometheus.Desc
	proftpdServer      *prometheus.Desc
}

func newProftpdCollector() *proftpdCollector {
	return &proftpdCollector{
		proftpdServer: prometheus.NewDesc(
			prometheus.BuildFQName(EXPORTER_NAMESPACE, "process", "uptime"),
			"Process Uptime Metrics for ProFTPd",
			[]string{"server_type"}, nil,
		),
		// user: username
		// port: port (21|22|etc)
		// protocol: protocol used to connect (ftp|sftp|etc)
		// state: idling or 'active' at the time of scrape
		// value: 'connected_since_ms'
		proftpdConnections: prometheus.NewDesc(
			prometheus.BuildFQName(EXPORTER_NAMESPACE, "user", "connections"),
			"Current User Connections with ProFTPd",
			[]string{"user", "port", "protocol", "remote_address", "state"}, nil,
		),
	}
}

func (collector *proftpdCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.proftpdServer
	collector.proftpdConnections = prometheus.NewDesc(
		prometheus.BuildFQName(EXPORTER_NAMESPACE, "user", "connections"),
		"Current User Connections with ProFTPd",
		[]string{"user", "port", "protocol", "remote_address", "state"}, nil,
	)
	ch <- collector.proftpdConnections
}

func (collector *proftpdCollector) Collect(ch chan<- prometheus.Metric) {
	results := parseFtpwho()

	ch <- prometheus.MustNewConstMetric(
		collector.proftpdServer,
		prometheus.GaugeValue,
		float64(results.Server.StartedMs),
		results.Server.ServerType,
	)

	uniqueConnections := make(map[string]ftpwhoUniqueConnection)

	// parse connections for unqiue
	for _, connection := range results.Connections {
		if conn, exists := uniqueConnections[connection.Username]; exists && conn.Connection.ConnectedSinceMs == connection.ConnectedSinceMs {
			conn.Count++
			uniqueConnections[connection.Username] = conn
		} else {
			uniqueConnections[connection.Username] = ftpwhoUniqueConnection{
				Count:      1,
				Connection: connection,
			}
		}
	}

	for username, uniqueConnection := range uniqueConnections {
		// []string{"user", "port", "protocol", "remote_address", "state"}
		ch <- prometheus.MustNewConstMetric(
			collector.proftpdConnections,
			prometheus.GaugeValue,
			float64(uniqueConnection.Count),
			username, fmt.Sprintf("%d", uniqueConnection.Connection.LocalPort), uniqueConnection.Connection.Protocol, uniqueConnection.Connection.RemoteAddress, uniqueConnection.Connection.DetermineState(),
		)
	}
}
