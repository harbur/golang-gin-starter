package prometheus

import "github.com/prometheus/client_golang/prometheus"

var appActions = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "app_actions",
	Help: "Counter of successful actions",
}, []string{"action"})

// IncrementAction counts actions
func IncrementAction(action string) {
	appActions.WithLabelValues(action).Inc()
}

// Register prometheus counters
func Register() {
	prometheus.Register(appActions)
}
