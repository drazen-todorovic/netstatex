package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"netstatex/util"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	tcpConnections = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "netstat_tcp_connection_total",
		Help: "Total number of tcp connections",
	}, []string{"protocol", "local", "foreign", "status"})
	udpConnections = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "netstat_udp_connection_total",
		Help: "Total number of udp connections",
	}, []string{"protocol", "local", "foreign"})
)

func runMetricWorker(intervalSeconds int) {
	for {
		netstatResult, _ := util.TryExecuteNetstat()
		tcp, udp := util.ParseNetstatOutput(netstatResult)

		tcpConnections.Reset()

		for _, item := range tcp {
			tcpConnections.WithLabelValues(
				item.Protocol,
				item.LocalAddress,
				item.ForeignAddress,
				item.State,
			).Inc()
		}

		udpConnections.Reset()

		for _, item := range udp {
			udpConnections.WithLabelValues(
				item.Protocol,
				item.LocalAddress,
				item.ForeignAddress,
			).Inc()
		}

		time.Sleep(time.Second * time.Duration(intervalSeconds))
	}
}

func main() {

	go runMetricWorker(1)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
