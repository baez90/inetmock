package main

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"gitlab.com/inetmock/inetmock/internal/app"
	"gitlab.com/inetmock/inetmock/internal/endpoint"
	dns "gitlab.com/inetmock/inetmock/internal/endpoint/handler/dns/mock"
	http "gitlab.com/inetmock/inetmock/internal/endpoint/handler/http/mock"
	"gitlab.com/inetmock/inetmock/internal/endpoint/handler/http/proxy"
	"gitlab.com/inetmock/inetmock/internal/endpoint/handler/metrics"
	"gitlab.com/inetmock/inetmock/internal/endpoint/handler/tls/interceptor"
	"gitlab.com/inetmock/inetmock/pkg/cert"
)

var (
	serverApp     app.App
	cfg           appConfig
	registrations = []endpoint.Registration{
		http.AddHTTPMock,
		dns.AddDNSMock,
		interceptor.AddTLSInterceptor,
		proxy.AddHTTPProxy,
		metrics.AddMetricsExporter,
	}
)

type Data struct {
	PCAP  string
	Audit string
}

func (d *Data) setup() (err error) {
	if d.PCAP, err = ensureDataDir(d.PCAP); err != nil {
		return
	}
	if d.Audit, err = ensureDataDir(d.Audit); err != nil {
		return
	}

	return
}

func ensureDataDir(dataDirPath string) (cleanedPath string, err error) {
	cleanedPath = dataDirPath
	if !filepath.IsAbs(cleanedPath) {
		if cleanedPath, err = filepath.Abs(cleanedPath); err != nil {
			return
		}
	}

	err = os.MkdirAll(cleanedPath, 0640)
	return
}

type appConfig struct {
	TLS       cert.Options
	Listeners map[string]endpoint.ListenerSpec
	API       struct {
		Listen string
	}
	Data Data
}

func (c *appConfig) APIURL() *url.URL {
	if u, err := url.Parse(c.API.Listen); err != nil {
		u, _ = url.Parse("tcp://:0")
		return u
	} else {
		return u
	}
}

func main() {
	serverApp = app.NewApp(
		app.Spec{
			Name:        "inetmock",
			Short:       "INetMock is lightweight internet mock",
			Config:      &cfg,
			SubCommands: []*cobra.Command{serveCmd, generateCaCmd},
			Defaults: map[string]interface{}{
				"api.listen": "tcp://:0",
				"data.pcap":  "/var/lib/inetmock/data/pcap",
				"data.audit": "/var/lib/inetmock/data/audit",
			},
		},
	)

	serverApp.MustRun()
}
