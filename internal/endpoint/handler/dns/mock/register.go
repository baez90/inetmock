package mock

import (
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/inetmock/inetmock/internal/endpoint"
	"gitlab.com/inetmock/inetmock/pkg/metrics"
)

const (
	name = "dns_mock"
)

var (
	handlerNameLblName          = "handler_name"
	totalHandledRequestsCounter *prometheus.CounterVec
	unhandledRequestsCounter    *prometheus.CounterVec
	requestDurationHistogram    *prometheus.HistogramVec
)

func AddDNSMock(registry endpoint.HandlerRegistry) (err error) {
	if totalHandledRequestsCounter, err = metrics.Counter(
		name,
		"handled_requests_total",
		"",
		handlerNameLblName,
	); err != nil {
		return
	}

	if unhandledRequestsCounter, err = metrics.Counter(
		name,
		"unhandled_requests_total",
		"",
		handlerNameLblName,
	); err != nil {
		return
	}

	if requestDurationHistogram, err = metrics.Histogram(
		name,
		"request_duration",
		"",
		nil,
		handlerNameLblName,
	); err != nil {
		return
	}

	registry.RegisterHandler(name, func() endpoint.ProtocolHandler {
		return &dnsHandler{}
	})

	return
}
