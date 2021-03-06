package mock

import (
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/inetmock/inetmock/internal/endpoint"
	"gitlab.com/inetmock/inetmock/pkg/metrics"
)

var (
	totalRequestCounter      *prometheus.CounterVec
	requestDurationHistogram *prometheus.HistogramVec
)

func AddHTTPMock(registry endpoint.HandlerRegistry) (err error) {
	if totalRequestCounter == nil {
		if totalRequestCounter, err = metrics.Counter(
			name,
			"total_requests",
			"",
			handlerNameLblName,
			ruleMatchedLblName,
		); err != nil {
			return
		}
	}

	if requestDurationHistogram == nil {
		if requestDurationHistogram, err = metrics.Histogram(
			name,
			"request_duration",
			"",
			nil,
			handlerNameLblName,
		); err != nil {
			return
		}
	}

	registry.RegisterHandler(name, func() endpoint.ProtocolHandler {
		return &httpHandler{}
	})

	return
}
