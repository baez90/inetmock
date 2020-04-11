package main

import (
	"github.com/baez90/inetmock/internal/plugins"
	"github.com/baez90/inetmock/pkg/api"
	"github.com/baez90/inetmock/pkg/logging"
	"go.uber.org/zap"
	"sync"
)

func init() {
	logger, _ := logging.CreateLogger()
	logger = logger.With(
		zap.String("ProtocolHandler", name),
	)

	plugins.Registry().RegisterHandler(name, func() api.ProtocolHandler {
		return &tlsInterceptor{
			logger:                  logger,
			currentConnectionsCount: &sync.WaitGroup{},
		}
	}, generateCACmd(logger))
}