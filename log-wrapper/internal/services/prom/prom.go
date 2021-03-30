package prom

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestResponseTime    = InitializeHttpResponseTime("http", "response_time", "Http Request response time for all endpoints")
	DependancyResponseTime = InitializeDependancyResponseTime("dependancy", "response_time", "Response time for all dependancies")
)

func InitializeHttpResponseTime(namespace, name, help string) *prometheus.SummaryVec {
	summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, []string{"status_class", "request", "method"})

	prometheus.MustRegister(summary)
	return summary
}

func RecordHttpResponseTime(status int, req, method string, st time.Time) {
	switch {
	case status >= 500:
		RequestResponseTime.WithLabelValues("5xx", req, method).Observe(float64(time.Since(st).Seconds()))
	case status >= 400:
		RequestResponseTime.WithLabelValues("4xx", req, method).Observe(float64(time.Since(st).Seconds()))
	case status >= 300:
		RequestResponseTime.WithLabelValues("3xx", req, method).Observe(float64(time.Since(st).Seconds()))
	case status >= 200:
		RequestResponseTime.WithLabelValues("2xx", req, method).Observe(float64(time.Since(st).Seconds()))
	default:
		RequestResponseTime.WithLabelValues("2xx", req, method).Observe(float64(time.Since(st).Seconds()))
	}
}

func InitializeDependancyResponseTime(namespace, name, help string) *prometheus.SummaryVec {
	summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, []string{"type", "request", "status_class"})

	prometheus.MustRegister(summary)
	return summary
}

func RecordDependancyResponseTime(depType, req, status string, v float64) {
	DependancyResponseTime.WithLabelValues(depType, req, status).Observe(v)
}
