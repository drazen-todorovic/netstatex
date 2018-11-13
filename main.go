package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kardianos/service"

	"github.com/drazen-todorovic/netstatex/util"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var logger service.Logger

type program struct{}

var (
	interval int
	port     int
)

func init() {
	flag.IntVar(&interval, "interval", 1, "Interval of collecting metrics (seconds)")
	flag.IntVar(&port, "port", 2112, "Listening port")
	flag.Parse()
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func (p *program) run() {
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

func main() {
	svcConfig := &service.Config{
		Name:        "netstatex",
		DisplayName: "Netstat Exporter",
		Description: "Prometheus netstat exporter",
	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)

	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)

	err = s.Run()

	if err != nil {
		logger.Error(err)
	}
}
