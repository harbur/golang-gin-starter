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
	appHTTPRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "app_http_requests_total",
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

	m := mux.NewRouter()
	m.Handle("/metrics", promhttp.Handler())
	m.Handle("/healthz", http.HandlerFunc(healthCheck)).Methods("GET")
	m.Handle("/app/{rest:.*}", promhttp.InstrumentHandlerCounter(appHTTPRequestsTotal, http.HandlerFunc(handlerOK))).Methods("GET")
	prometheus.MustRegister(appHTTPRequestsTotal)

	// Listen and Serve
	log.Fatal(http.ListenAndServe(":8080", m))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func handlerOK(w http.ResponseWriter, r *http.Request) {
	log.WithField("path", r.URL.Path).Info("Received request")
	w.Write([]byte(r.URL.Path))
}
