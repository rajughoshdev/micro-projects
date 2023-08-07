package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Total number of Http requests",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestTotal)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "INSTRUMENTING A GO APPLICATION FOR PROMETHEUS %s", time.Now())
	httpRequestTotal.WithLabelValues("/").Inc()
}
func main() {
	http.HandleFunc("/", greeting)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
