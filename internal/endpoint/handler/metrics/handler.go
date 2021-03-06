package metrics

import (
	"errors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/inetmock/inetmock/internal/endpoint"
	"gitlab.com/inetmock/inetmock/pkg/logging"
	"go.uber.org/zap"
)

const (
	name = "metrics_exporter"
)

type metricsExporter struct {
	logger logging.Logger
	server *http.Server
}

func (m *metricsExporter) Start(lifecycle endpoint.Lifecycle) (err error) {
	var exporterOptions metricsExporterOptions
	if err = lifecycle.UnmarshalOptions(&exporterOptions); err != nil {
		return
	}

	m.logger = m.logger.With(
		zap.String("handler_name", lifecycle.Name()),
		zap.String("address", lifecycle.Uplink().Addr().String()),
	)

	mux := http.NewServeMux()
	mux.Handle(exporterOptions.Route, promhttp.Handler())
	m.server = &http.Server{
		Handler: mux,
	}

	go func() {
		if err := m.server.Serve(lifecycle.Uplink().Listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			m.logger.Error(
				"Error occurred while serving metrics",
				zap.Error(err),
			)
		}
	}()

	go func() {
		<-lifecycle.Context().Done()
		if err := m.server.Close(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			m.logger.Error("failed to stop metrics server", zap.Error(err))
		}
	}()
	return
}
