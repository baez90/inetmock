package mock_test

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"gitlab.com/inetmock/inetmock/internal/endpoint"
	"gitlab.com/inetmock/inetmock/internal/endpoint/handler/http/mock"
	"gitlab.com/inetmock/inetmock/pkg/audit"
	"gitlab.com/inetmock/inetmock/pkg/logging"
)

var (
	availableExtensions = []string{"gif", "html", "ico", "jpg", "png", "txt"}
	charSet             = "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Benchmark_httpHandler(b *testing.B) {
	var listener net.Listener
	var err error

	var uplink endpoint.Uplink

	if uplink.Listener, err = net.Listen("tcp", ":0"); err != nil {
		b.Errorf("Uplink() error = %v", err)
	}

	b.Cleanup(func() {
		_ = uplink.Close()
	})

	var port int

	if tcpListener, ok := listener.(*net.TCPListener); ok {
		if tcpAddr, ok := tcpListener.Addr().(*net.TCPAddr); ok {
			port = tcpAddr.Port
		} else {
			b.Errorf("not an TCP addr")
		}
	} else {
		b.Errorf("not a TCP listener")
	}

	var cancel context.CancelFunc
	if cancel, err = setupHandler(b, uplink); err != nil {
		b.Errorf("setupHandler() error = %v", err)
	}
	b.Cleanup(cancel)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		extension := availableExtensions[rand.Intn(len(availableExtensions))]
		if resp, err := http.Get(fmt.Sprintf("http://localhost:%d/%s.%s", port, randomString(15), extension)); err != nil {
			b.Error(err)
		} else if resp.StatusCode != 200 {
			b.Errorf("Got status code %d", resp.StatusCode)
		}
	}
}

func randomString(length int) (result string) {
	buffer := strings.Builder{}
	for i := 0; i < length; i++ {
		buffer.WriteByte(charSet[rand.Intn(len(charSet))])
	}
	return buffer.String()
}

func setupHandler(b *testing.B, uplink endpoint.Uplink) (cancel context.CancelFunc, err error) {
	b.Helper()

	registry := endpoint.NewHandlerRegistry()
	if err := mock.AddHTTPMock(registry); err != nil {
		b.Errorf("AddHTTPMock() error = %v", err)
	}
	handler, ok := registry.HandlerForName("http_mock")
	if !ok {
		b.Error("handler not registered")
	}

	opts := map[string]interface{}{
		"rules": []struct {
			Pattern  string
			Response string
		}{
			{
				Pattern:  ".*\\.(?i)gif",
				Response: "./../../../../../assets/fakeFiles/default.gif",
			},
			{
				Pattern:  ".*\\.(?i)html",
				Response: "./../../../../../assets/fakeFiles/default.html",
			},
			{
				Pattern:  ".*\\.(?i)ico",
				Response: "./../../../../../assets/fakeFiles/default.ico",
			},
			{
				Pattern:  ".*\\.(?i)jpg",
				Response: "./../../../../../assets/fakeFiles/default.jpg",
			},
			{
				Pattern:  ".*\\.(?i)png",
				Response: "./../../../../../assets/fakeFiles/default.png",
			},
			{
				Pattern:  ".*\\.(?i)txt",
				Response: "./../../../../../assets/fakeFiles/default.txt",
			},
		},
	}

	var ctx context.Context
	ctx, cancel = context.WithCancel(context.Background())

	logger := logging.CreateTestLogger(b)

	var eventStream audit.Emitter
	eventStream, err = audit.NewEventStream(logger)

	lifecycle := endpoint.NewEndpointLifecycleFromContext(
		"test",
		ctx,
		logger,
		nil,
		eventStream,
		uplink,
		opts,
	)

	err = handler.Start(lifecycle)

	return
}
