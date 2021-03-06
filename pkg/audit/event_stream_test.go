package audit_test

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"gitlab.com/inetmock/inetmock/pkg/audit"
	"gitlab.com/inetmock/inetmock/pkg/audit/details"
	"gitlab.com/inetmock/inetmock/pkg/audit/sink"
	"gitlab.com/inetmock/inetmock/pkg/logging"
	"gitlab.com/inetmock/inetmock/pkg/wait"
)

var (
	noOpSink   = sink.NewNoOpSink("test defaultSink")
	testEvents = []*audit.Event{
		{
			Transport:       audit.TransportProtocol_TCP,
			Application:     audit.AppProtocol_HTTP,
			SourceIP:        net.ParseIP("127.0.0.1").To4(),
			DestinationIP:   net.ParseIP("127.0.0.1").To4(),
			SourcePort:      32344,
			DestinationPort: 80,
			TLS: &audit.TLSDetails{
				Version:     audit.TLSVersionToEntity(tls.VersionTLS13).String(),
				CipherSuite: tls.CipherSuiteName(tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA),
				ServerName:  "localhost",
			},
			ProtocolDetails: details.HTTP{
				Method: "GET",
				Host:   "localhost",
				URI:    "http://localhost/asdf",
				Proto:  "HTTP 1.1",
				Headers: http.Header{
					"Accept": []string{"application/json"},
				},
			},
		},
		{
			Transport:       audit.TransportProtocol_TCP,
			Application:     audit.AppProtocol_DNS,
			SourceIP:        net.ParseIP("::1").To16(),
			DestinationIP:   net.ParseIP("::1").To16(),
			SourcePort:      32344,
			DestinationPort: 80,
		},
	}
)

func wgMockSink(t testing.TB, wg *sync.WaitGroup) audit.Sink {
	return sink.NewGenericSink(
		"WG mock sink",
		func(event audit.Event) {
			t.Logf("Got event = %v", event)
			wg.Done()
		},
	)
}

func Test_eventStream_RegisterSink(t *testing.T) {
	type args struct {
		s audit.Sink
	}
	type testCase struct {
		name    string
		args    args
		setup   func(e audit.EventStream)
		wantErr bool
	}
	tests := []testCase{
		{
			name: "Register test defaultSink",
			args: args{
				s: noOpSink,
			},
			wantErr: false,
		},
		{
			name: "Fail due to already registered defaultSink",
			args: args{
				s: noOpSink,
			},
			setup: func(e audit.EventStream) {
				_ = e.RegisterSink(context.Background(), noOpSink)
			},
			wantErr: true,
		},
	}
	scenario := func(tt testCase) func(t *testing.T) {
		return func(t *testing.T) {
			var err error
			var e audit.EventStream
			if e, err = audit.NewEventStream(logging.CreateTestLogger(t)); err != nil {
				t.Errorf("NewEventStream() error = %v", err)
			}

			t.Cleanup(func() {
				_ = e.Close()
			})

			if tt.setup != nil {
				tt.setup(e)
			}

			if err := e.RegisterSink(context.Background(), tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("RegisterSink() error = %v, wantErr %v", err, tt.wantErr)
			}

			found := false
			for _, s := range e.Sinks() {
				if found = s == tt.args.s.Name(); found {
					break
				}
			}
			if !found {
				t.Errorf("expected defaultSink name %s not found in registered sinks %v", tt.args.s.Name(), e.Sinks())
			}
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, scenario(tt))
	}
}

func Test_eventStream_Emit(t *testing.T) {
	type args struct {
		evs  []*audit.Event
		opts []audit.EventStreamOption
	}
	type testCase struct {
		name      string
		args      args
		subscribe bool
	}
	tests := []testCase{
		{
			name:      "Expect to get a single event",
			subscribe: true,
			args: args{
				opts: []audit.EventStreamOption{},
				evs:  testEvents[:1],
			},
		},
		{
			name:      "Expect to get multiple events",
			subscribe: true,
			args: args{
				opts: []audit.EventStreamOption{},
				evs:  testEvents,
			},
		},
		{
			name: "Emit without subscribe sink",
			args: args{
				opts: []audit.EventStreamOption{audit.WithBufferSize(0)},
				evs:  testEvents[:1],
			},
			subscribe: false,
		},
	}

	scenario := func(tt testCase) func(t *testing.T) {
		return func(t *testing.T) {
			var err error
			var e audit.EventStream
			if e, err = audit.NewEventStream(logging.CreateTestLogger(t), tt.args.opts...); err != nil {
				t.Errorf("NewEventStream() error = %v", err)
			}

			t.Cleanup(func() {
				_ = e.Close()
			})

			emittedWaitGroup := new(sync.WaitGroup)
			receivedWaitGroup := new(sync.WaitGroup)

			emittedWaitGroup.Add(len(tt.args.evs))

			if tt.subscribe {
				receivedWaitGroup.Add(len(tt.args.evs))
				if err := e.RegisterSink(context.Background(), wgMockSink(t, receivedWaitGroup)); err != nil {
					t.Errorf("RegisterSink() error = %v", err)
				}
			}

			go func(evs []*audit.Event, wg *sync.WaitGroup) {
				for _, ev := range evs {
					e.Emit(*ev)
					wg.Done()
				}
			}(tt.args.evs, emittedWaitGroup)

			select {
			case <-wait.ForWaitGroupDone(emittedWaitGroup):
			case <-time.After(100 * time.Millisecond):
				t.Errorf("not all events emitted in time")
			}

			select {
			case <-wait.ForWaitGroupDone(receivedWaitGroup):
			case <-time.After(5 * time.Second):
				t.Errorf("did not get all expected events in time")
			}
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, scenario(tt))
	}
}
