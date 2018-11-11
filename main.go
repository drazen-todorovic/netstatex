package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/drazen-todorovic/netstatex/util"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	interval int
	port     int
)

func init() {
	flag.IntVar(&interval, "interval", 1, "Interval of collecting metrics (seconds)")
	flag.IntVar(&port, "port", 2112, "Listening port")
	flag.Parse()
}

func main() {

	go util.RunMetricWorker(interval)

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println(fmt.Sprintf("Collecting interval: %v seconds", interval))
	fmt.Println(fmt.Sprintf("Running on http://:%v", port))

	addr := fmt.Sprintf(":%v", port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}

}
