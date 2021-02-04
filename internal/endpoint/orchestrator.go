package endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/soheilhy/cmux"
	"gitlab.com/inetmock/inetmock/pkg/audit"
	"gitlab.com/inetmock/inetmock/pkg/cert"
	"gitlab.com/inetmock/inetmock/pkg/health"
	"gitlab.com/inetmock/inetmock/pkg/logging"
	"go.uber.org/zap"
)

const (
	startupTimeoutDuration = 100 * time.Millisecond
)

var (
	ErrStartupTimeout = errors.New("endpoint did not start in time")
)

type Endpoint struct {
	lifecycle Lifecycle
	cancel    context.CancelFunc
	Name      string
}

type Orchestrator interface {
	RegisterListener(spec ListenerSpec) error
	StartEndpoints() (errChan chan error)
	ShutdownEndpoints()
}

func NewOrchestrator(appCtx context.Context, certStore cert.Store, registry HandlerRegistry, emitter audit.Emitter, logging logging.Logger, checker health.Checker) Orchestrator {
	return &orchestrator{
		appCtx:                   appCtx,
		registry:                 registry,
		logger:                   logging,
		certStore:                certStore,
		emitter:                  emitter,
		checker:                  checker,
		properlyStartedEndpoints: make(map[string]Endpoint),
	}
}

type orchestrator struct {
	appCtx    context.Context
	registry  HandlerRegistry
	logger    logging.Logger
	certStore cert.Store
	emitter   audit.Emitter
	checker   health.Checker

	endpointListeners        []endpointListener
	properlyStartedEndpoints map[string]Endpoint
	muxes                    []cmux.CMux
}

type endpointListener struct {
	name   string
	uplink Uplink
	spec   Spec
}

func (e *orchestrator) RegisterListener(spec ListenerSpec) (err error) {
	for name, s := range spec.Endpoints {
		if handler, registered := e.registry.HandlerForName(s.HandlerRef); registered {
			s.Handler = handler
			spec.Endpoints[name] = s
		}
	}

	var epListeners []endpointListener
	var muxes []cmux.CMux
	if epListeners, muxes, err = endpointListenersFromSpec(spec, e.certStore.TLSConfig()); err != nil {
		return
	}

	e.endpointListeners = append(e.endpointListeners, epListeners...)
	e.muxes = append(e.muxes, muxes...)

	return
}

func (e *orchestrator) StartEndpoints() (errChan chan error) {
	errChan = make(chan error)
	for _, epListener := range e.endpointListeners {
		endpointLogger := e.logger.With(
			zap.String("epListener", epListener.name),
		)
		endpointLogger.Info("Starting epListener")
		epCtx, cancel := context.WithCancel(e.appCtx)
		lifecycle := NewEndpointLifecycleFromContext(
			epListener.name,
			epCtx,
			e.logger.With(zap.String("epListener", epListener.name)),
			e.certStore,
			e.emitter,
			epListener.uplink,
			epListener.spec.Options,
		)

		ep := Endpoint{
			lifecycle: lifecycle,
			cancel:    cancel,
			Name:      epListener.name,
		}

		if err := startEndpoint(epListener.spec, lifecycle, endpointLogger); err == nil {
			_ = e.checker.RegisterCheck(
				endpointComponentName(ep),
				health.StaticResultCheckWithMessage(health.HEALTHY, "Successfully started"),
			)
			e.properlyStartedEndpoints[epListener.name] = ep
			endpointLogger.Info("successfully started epListener")
		} else {
			_ = e.checker.RegisterCheck(
				endpointComponentName(ep),
				health.StaticResultCheckWithMessage(health.UNHEALTHY, "failed to start"),
			)
			endpointLogger.Error("error occurred during epListener startup - will be skipped for now")
		}
	}
	e.logger.Info("Startup of all endpoints completed")

	for _, mux := range e.muxes {
		go func(mux cmux.CMux) {
			mux.HandleError(func(err error) bool {
				errChan <- err
				return true
			})
			if err := mux.Serve(); err != nil && !errors.Is(err, cmux.ErrListenerClosed) {
				errChan <- err
			}
		}(mux)
	}

	return
}

func (e *orchestrator) ShutdownEndpoints() {
	for name, endpoint := range e.properlyStartedEndpoints {
		e.logger.Info("Triggering shutdown of endpoint", zap.String("endpoint", name))
		endpoint.cancel()
		delete(e.properlyStartedEndpoints, name)
	}
}

func startEndpoint(ep Spec, lifecycle Lifecycle, logger logging.Logger) (err error) {
	startupResult := make(chan error)
	ctx, cancel := context.WithTimeout(lifecycle.Context(), startupTimeoutDuration)
	defer cancel()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Fatal(
					"recovered panic during startup of endpoint",
					zap.Any("recovered", r),
				)
				startupResult <- fmt.Errorf("recovered: %v", r)
			}
		}()

		startupResult <- ep.Handler.Start(lifecycle)
	}()

	select {
	case err = <-startupResult:
	case <-ctx.Done():
		err = ErrStartupTimeout
	}

	return
}

func endpointComponentName(ep Endpoint) string {
	return fmt.Sprintf("endpoint_%s", ep.Name)
}
