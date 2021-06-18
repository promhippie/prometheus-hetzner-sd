package action

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/promhippie/prometheus-hetzner-sd/pkg/version"
)

var (
	registry  = prometheus.NewRegistry()
	namespace = "prometheus_hetzner_sd"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "request_duration_seconds",
			Help:      "Histogram of latencies for requests to the Hetzner API.",
			Buckets:   []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.0, 5.0, 10.0},
		},
		[]string{"project"},
	)

	requestFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "request_failures_total",
			Help:      "Total number of failed requests to the Hetzner API.",
		},
		[]string{"project"},
	)
)

func init() {
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{
		Namespace: namespace,
	}))

	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(version.Collector(namespace))

	registry.MustRegister(requestDuration)
	registry.MustRegister(requestFailures)
}

type promLogger struct {
	logger log.Logger
}

func (pl promLogger) Println(v ...interface{}) {
	level.Error(pl.logger).Log(
		"msg", fmt.Sprintln(v...),
	)
}
