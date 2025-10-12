package metrics

import "github.com/prometheus/client_golang/prometheus"



var (
    HttpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests processed",
        },
        []string{"group", "method", "path", "status"},
    )

    HttpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request latency distributions",
            Buckets: prometheus.DefBuckets,
        },
        []string{"group", "path"},
    )
)