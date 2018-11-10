package main

import (
	"fmt"
	"net/http"

	"github.com/drazen-todorovic/netstatex/util"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	go util.RunMetricWorker(1)

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Prometheus netstat exproter is running on http://0.0.0.0:2112")

	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		panic(err)
	}

}
