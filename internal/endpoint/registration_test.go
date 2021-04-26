package endpoint_test

import (
	"testing"
	"testing/fstest"

	"github.com/golang/mock/gomock"
	"github.com/maxatome/go-testdeep/td"

	"gitlab.com/inetmock/inetmock/internal/endpoint"
	dnsmock "gitlab.com/inetmock/inetmock/internal/endpoint/handler/dns/mock"
	httpmock "gitlab.com/inetmock/inetmock/internal/endpoint/handler/http/mock"
	audit_mock "gitlab.com/inetmock/inetmock/internal/mock/audit"
	"gitlab.com/inetmock/inetmock/pkg/logging"
)

func Test_handlerRegistry_AvailableHandlers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                  string
		handlerRegistrySetup  func(tb testing.TB, ctrl *gomock.Controller) endpoint.HandlerRegistry
		wantAvailableHandlers interface{}
	}{
		{
			name: "Empty registry",
			handlerRegistrySetup: func(testing.TB, *gomock.Controller) endpoint.HandlerRegistry {
				return endpoint.NewHandlerRegistry()
			},
			wantAvailableHandlers: td.Nil(),
		},
		{
			name: "Single handler registered",
			handlerRegistrySetup: func(tb testing.TB, ctrl *gomock.Controller) endpoint.HandlerRegistry {
				tb.Helper()
				registry := endpoint.NewHandlerRegistry()
				logger := logging.CreateTestLogger(tb)
				emitter := audit_mock.NewMockEmitter(ctrl)
				if err := httpmock.AddHTTPMock(registry, logger, emitter, new(fstest.MapFS)); err != nil {
					tb.Fatalf("AddHTTPMock() error = %v", err)
				}
				return registry
			},
			wantAvailableHandlers: td.Set(endpoint.HandlerReference("http_mock")),
		},
		{
			name: "Multiple handler registered",
			handlerRegistrySetup: func(tb testing.TB, ctrl *gomock.Controller) endpoint.HandlerRegistry {
				tb.Helper()
				registry := endpoint.NewHandlerRegistry()
				logger := logging.CreateTestLogger(tb)
				emitter := audit_mock.NewMockEmitter(ctrl)
				if err := httpmock.AddHTTPMock(registry, logger, emitter, new(fstest.MapFS)); err != nil {
					tb.Fatalf("AddHTTPMock() error = %v", err)
				}
				if err := dnsmock.AddDNSMock(registry, logger, emitter); err != nil {
					tb.Fatalf("AddHTTPMock() error = %v", err)
				}
				return registry
			},
			wantAvailableHandlers: td.Set(
				endpoint.HandlerReference("dns_mock"),
				endpoint.HandlerReference("http_mock"),
			),
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			gotAvailableHandlers := tt.handlerRegistrySetup(t, ctrl).AvailableHandlers()
			td.Cmp(t, gotAvailableHandlers, tt.wantAvailableHandlers)
		})
	}
}

func Test_handlerRegistry_HandlerForName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                 string
		handlerRegistrySetup func(tb testing.TB, ctrl *gomock.Controller) endpoint.HandlerRegistry
		handlerRef           endpoint.HandlerReference
		wantInstance         interface{}
		wantOk               bool
	}{
		{
			name: "Empty registry",
			handlerRegistrySetup: func(tb testing.TB, _ *gomock.Controller) endpoint.HandlerRegistry {
				tb.Helper()
				return endpoint.NewHandlerRegistry()
			},
			handlerRef:   "http_mock",
			wantInstance: nil,
			wantOk:       false,
		},
		{
			name: "Registry with HTTP mock registered",
			handlerRegistrySetup: func(tb testing.TB, ctrl *gomock.Controller) endpoint.HandlerRegistry {
				tb.Helper()
				registry := endpoint.NewHandlerRegistry()
				logger := logging.CreateTestLogger(tb)
				emitter := audit_mock.NewMockEmitter(ctrl)
				if err := httpmock.AddHTTPMock(registry, logger, emitter, new(fstest.MapFS)); err != nil {
					tb.Fatalf("AddHTTPMock() error = %v", err)
				}
				return registry
			},
			handlerRef:   "http_mock",
			wantInstance: td.NotNil(),
			wantOk:       true,
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			gotInstance, gotOk := tt.handlerRegistrySetup(t, ctrl).HandlerForName(tt.handlerRef)
			td.Cmp(t, gotInstance, tt.wantInstance)
			td.Cmp(t, gotOk, tt.wantOk)
		})
	}
}
