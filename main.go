package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Use JSON Formatter when deployed on Kubernetes
	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.Info("starting server")

	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/{rest:.*}", handlerOK).Methods("GET")

	// Listen and Serve
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handlerOK(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{"path": r.URL.Path}).Info("Received request")
	w.Write([]byte(r.URL.Path))
}
