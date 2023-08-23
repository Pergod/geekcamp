package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// 创建一个Histogram指标
	HttpRequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_latency_seconds",
			Help:    "Latency of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets, // 使用默认的分桶范围
		},
		[]string{"handler"},
	)
)

func Register() {
	// 创建一个Histogram指标
	prometheus.MustRegister(HttpRequestLatency)
}
