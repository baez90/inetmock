package main

import (
	"github.com/spf13/cobra"
	"gitlab.com/inetmock/inetmock/internal/rpc"
	"go.uber.org/zap"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the INetMock server",
		Long:  ``,
		RunE:  startINetMock,
	}
)

func startINetMock(_ *cobra.Command, _ []string) (err error) {
	rpcAPI := rpc.NewINetMockAPI(serverApp)
	logger := serverApp.Logger()

	cfg := serverApp.Config()
	endpointOrchestrator := serverApp.EndpointManager()

	for ref, spec := range cfg.ListenerSpecs() {
		if err = endpointOrchestrator.RegisterListener(ref, spec); err != nil {
			logger.Error("Failed to register listener", zap.Error(err))
			return
		}
	}

	for endpointName, endpointHandler := range cfg.MetaSpecs() {
		if err := endpointOrchestrator.RegisterEndpoint(endpointName, endpointHandler); err != nil {
			logger.Warn(
				"error occurred while creating endpoint",
				zap.String("endpointName", endpointName),
				zap.String("handlerName", string(endpointHandler.Handler)),
				zap.Error(err),
			)
		}
	}

	serverApp.EndpointManager().StartEndpoints()
	if err = rpcAPI.StartServer(); err != nil {
		serverApp.Shutdown()
		logger.Error(
			"failed to start gRPC API",
			zap.Error(err),
		)
	}

	<-serverApp.Context().Done()

	logger.Info("App context canceled - shutting down")

	rpcAPI.StopServer()
	serverApp.EndpointManager().ShutdownEndpoints()
	return
}
