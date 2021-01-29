package endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	RegisteredEndpoints() []string
	StartedEndpoints() []Endpoint
	RegisterEndpoint(name string, multiHandlerConfig MetaSpec) error
	RegisterListener(ref ListenerReference, spec ListenerSpec) error
	StartEndpoints()
	ShutdownEndpoints()
}

func NewEndpointManager(appCtx context.Context, certStore cert.Store, registry HandlerRegistry, emitter audit.Emitter, logging logging.Logger, checker health.Checker) Orchestrator {
	return &endpointManager{
		appCtx:                   appCtx,
		registry:                 registry,
		logger:                   logging,
		certStore:                certStore,
		emitter:                  emitter,
		checker:                  checker,
		listeners:                make(map[ListenerReference]Uplink),
		registeredSpecs:          make(map[string]Spec),
		properlyStartedEndpoints: make(map[string]Endpoint),
	}
}

type endpointManager struct {
	appCtx                   context.Context
	registry                 HandlerRegistry
	logger                   logging.Logger
	certStore                cert.Store
	emitter                  audit.Emitter
	checker                  health.Checker
	listeners                map[ListenerReference]Uplink
	registeredSpecs          map[string]Spec
	properlyStartedEndpoints map[string]Endpoint
}

func (e endpointManager) RegisteredEndpoints() (eps []string) {
	for n := range e.registeredSpecs {
		eps = append(eps, n)
	}
	return
}

func (e endpointManager) StartedEndpoints() (eps []Endpoint) {
	for _, ep := range e.properlyStartedEndpoints {
		eps = append(eps, ep)
	}
	return
}

func (e *endpointManager) RegisterListener(ref ListenerReference, spec ListenerSpec) (err error) {
	var uplink Uplink
	if uplink, err = spec.Uplink(); err != nil {
		return
	}
	e.listeners[ref] = uplink
	return
}

func (e *endpointManager) RegisterEndpoint(name string, endpointConfig MetaSpec) error {
	for _, spec := range endpointConfig.EndpointSpecs() {
		if handler, ok := e.registry.HandlerForName(endpointConfig.Handler); ok {
			spec.Handler = handler
		} else {
			return fmt.Errorf("no matching handler registered for names %s", endpointConfig.Handler)
		}
		if _, ok := e.listeners[spec.ListenerRef]; !ok {
			return fmt.Errorf("no matching uplink registered for reference %s", spec.ListenerRef)
		}
		e.registeredSpecs[fmt.Sprintf("%s_%s", name, spec.ListenerRef)] = spec
	}

	return nil
}

func (e *endpointManager) StartEndpoints() {
	startTime := time.Now()
	for name, spec := range e.registeredSpecs {
		endpointLogger := e.logger.With(
			zap.String("spec", name),
		)
		endpointLogger.Info("Starting spec")
		epCtx, cancel := context.WithCancel(e.appCtx)
		lifecycle := NewEndpointLifecycleFromContext(
			name,
			epCtx,
			e.logger.With(zap.String("spec", name)),
			e.certStore,
			e.emitter,
			e.listeners[spec.ListenerRef],
			spec.Options,
		)

		ep := Endpoint{
			lifecycle: lifecycle,
			cancel:    cancel,
			Name:      name,
		}

		if err := startEndpoint(spec, lifecycle, endpointLogger); err == nil {
			_ = e.checker.RegisterCheck(
				endpointComponentName(ep),
				health.StaticResultCheckWithMessage(health.HEALTHY, "Successfully started"),
			)
			e.properlyStartedEndpoints[name] = ep
			endpointLogger.Info("successfully started spec")
		} else {
			_ = e.checker.RegisterCheck(
				endpointComponentName(ep),
				health.StaticResultCheckWithMessage(health.UNHEALTHY, "failed to start"),
			)
			endpointLogger.Error("error occurred during spec startup - will be skipped for now")
		}
	}
	endpointStartupDuration := time.Since(startTime)
	e.logger.Info(
		"Startup of all endpoints completed",
		zap.Duration("startupTime", endpointStartupDuration),
	)
}

func (e *endpointManager) ShutdownEndpoints() {
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
