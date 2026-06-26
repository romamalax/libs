package metrics

import (
	"sync/atomic"

	"github.com/VictoriaMetrics/metrics"
)

func InitMetricGaugeInt64(name string, valAddr *int64) {
	metrics.GetOrCreateGauge(name, func() float64 {
		return float64(atomic.LoadInt64(valAddr))
	})
}

func GetOrCreateCounter(name string) *metrics.Counter {
	return metrics.GetOrCreateCounter(name)
}

func GetOrCreateDurationHistogram(name string, buckets []float64) *metrics.PrometheusHistogram {
	return metrics.GetOrCreatePrometheusHistogramExt(name, buckets)
}
