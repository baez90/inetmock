package proxy

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
	imHttp "gitlab.com/inetmock/inetmock/internal/endpoint/handler/http"
	"gitlab.com/inetmock/inetmock/pkg/audit"
	"gitlab.com/inetmock/inetmock/pkg/logging"
	"go.uber.org/zap"
	"gopkg.in/elazarl/goproxy.v1"
)

type proxyHttpsHandler struct {
	options   httpProxyOptions
	tlsConfig *tls.Config
	emitter   audit.Emitter
}

func (p *proxyHttpsHandler) HandleConnect(_ string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
	p.emitter.Emit(imHttp.EventFromRequest(ctx.Req, audit.AppProtocol_HTTP_PROXY))

	return &goproxy.ConnectAction{
		Action: goproxy.ConnectAccept,
		TLSConfig: func(host string, ctx *goproxy.ProxyCtx) (*tls.Config, error) {
			return p.tlsConfig, nil
		},
	}, p.options.Target.host()
}

type proxyHttpHandler struct {
	handlerName string
	options     httpProxyOptions
	logger      logging.Logger
	emitter     audit.Emitter
}

func (p *proxyHttpHandler) Handle(req *http.Request, ctx *goproxy.ProxyCtx) (retReq *http.Request, resp *http.Response) {
	timer := prometheus.NewTimer(requestDurationHistogram.WithLabelValues(p.handlerName))
	defer timer.ObserveDuration()

	retReq = req
	p.emitter.Emit(imHttp.EventFromRequest(req, audit.AppProtocol_HTTP_PROXY))

	var err error
	if resp, err = ctx.RoundTrip(redirectHTTPRequest(p.options.Target.host(), req)); err != nil {
		p.logger.Error(
			"error while doing roundtrip",
			zap.Error(err),
		)
		return req, nil
	}

	return
}

func redirectHTTPRequest(targetHost string, originalRequest *http.Request) (redirectReq *http.Request) {
	redirectReq = &http.Request{
		Method: originalRequest.Method,
		URL: &url.URL{
			Host:       targetHost,
			Path:       originalRequest.URL.Path,
			ForceQuery: originalRequest.URL.ForceQuery,
			Fragment:   originalRequest.URL.Fragment,
			Opaque:     originalRequest.URL.Opaque,
			RawPath:    originalRequest.URL.RawPath,
			RawQuery:   originalRequest.URL.RawQuery,
			User:       originalRequest.URL.User,
		},
		Proto:            originalRequest.Proto,
		ProtoMajor:       originalRequest.ProtoMajor,
		ProtoMinor:       originalRequest.ProtoMinor,
		Header:           originalRequest.Header,
		Body:             originalRequest.Body,
		GetBody:          originalRequest.GetBody,
		ContentLength:    originalRequest.ContentLength,
		TransferEncoding: originalRequest.TransferEncoding,
		Close:            false,
		Host:             originalRequest.Host,
		Form:             originalRequest.Form,
		PostForm:         originalRequest.PostForm,
		MultipartForm:    originalRequest.MultipartForm,
		Trailer:          originalRequest.Trailer,
	}
	redirectReq = redirectReq.WithContext(originalRequest.Context())

	return
}
