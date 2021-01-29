package mock

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"

	"gitlab.com/inetmock/inetmock/internal/endpoint"
	imHttp "gitlab.com/inetmock/inetmock/internal/endpoint/handler/http"
	"gitlab.com/inetmock/inetmock/pkg/logging"
	"go.uber.org/zap"
)

const (
	name               = "http_mock"
	handlerNameLblName = "handler_name"
	ruleMatchedLblName = "rule_matched"
)

type httpHandler struct {
	logger logging.Logger
	server *http.Server
}

func (p *httpHandler) Start(lifecycle endpoint.Lifecycle) (err error) {
	p.logger = lifecycle.Logger().With(
		zap.String("protocol_handler", name),
	)

	var options httpOptions
	if options, err = loadFromConfig(lifecycle); err != nil {
		return
	}

	p.logger = p.logger.With(
		zap.String("address", lifecycle.Uplink().Addr().String()),
	)

	router := &RegexpHandler{
		logger:      p.logger,
		emitter:     lifecycle.Audit(),
		handlerName: lifecycle.Name(),
	}
	p.server = &http.Server{
		Handler:     router,
		ConnContext: imHttp.StoreConnPropertiesInContext,
	}

	if options.TLS {
		p.server.TLSConfig = lifecycle.CertStore().TLSConfig()
		p.server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler))
	}

	for _, rule := range options.Rules {
		router.setupRoute(rule)
	}

	go p.startServer(options.TLS, lifecycle.Uplink().Listener)
	go p.shutdownOnCancel(lifecycle.Context())
	return
}

func (p *httpHandler) shutdownOnCancel(ctx context.Context) {
	<-ctx.Done()
	p.logger.Info("Shutting down HTTP mock")
	if err := p.server.Close(); err != nil {
		p.logger.Error(
			"failed to shutdown HTTP server",
			zap.Error(err),
		)
	}
	return
}

func (p *httpHandler) startServer(tls bool, listener net.Listener) {
	var serve func(listener net.Listener) error
	if tls {
		serve = func(listener net.Listener) error {
			return p.server.ServeTLS(listener, "", "")
		}
	} else {
		serve = p.server.Serve
	}

	if err := serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		p.logger.Error(
			"failed to start http listener",
			zap.Error(err),
		)
	}
}
