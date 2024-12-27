package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/arianvp/cgroup-exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	addr := flag.String("listen-address", ":13232", "address to listen on")
	cgroup := flag.String("cgroup", "", "what cgroup to monitor. Can be a blob. If empty all cgroups are monitored.")
	flag.Parse()
	cgroupfs := os.DirFS("/sys/fs/cgroup")
	registry := prometheus.NewRegistry()
	registry.MustRegister(collector.New(cgroupfs, *cgroup))
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	// TODO: cancellation
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
