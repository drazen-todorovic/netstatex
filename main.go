package main

import (
	"net/http"

	"netstatex/util"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	go util.RunMetricWorker(1)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
