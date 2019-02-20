package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("starting server")

	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())

	// Listen and Serve
	log.Fatal(http.ListenAndServe(":8080", r))
}
