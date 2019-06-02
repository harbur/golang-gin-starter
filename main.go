package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": "v0.1.0",
		},
	})
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})
)

func main() {
	// Use JSON Formatter when deployed on Kubernetes
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.Info("starting server")
	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(version)

	m := mux.NewRouter()
	m.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	m.Handle("/{rest:.*}", promhttp.InstrumentHandlerCounter(httpRequestsTotal, http.HandlerFunc(handlerOK))).Methods("GET")

	// Listen and Serve
	log.Fatal(http.ListenAndServe(":8080", m))
}

func handlerOK(w http.ResponseWriter, r *http.Request) {
	log.WithField("path", r.URL.Path).Info("Received request")
	w.Write([]byte(r.URL.Path))
}
