package endpoint

import (
	"crypto/tls"
	"errors"
	"fmt"
	"sort"

	"github.com/soheilhy/cmux"
)

var (
	ErrUDPMultiplexer           = errors.New("UDP listeners don't support multiplexing")
	ErrMultiplexingNotSupported = errors.New("not all handlers do support multiplexing")
)

func endpointListenersFromSpec(spec ListenerSpec, tlsConfig *tls.Config) (listeners []endpointListener, muxes []cmux.CMux, err error) {
	var uplink Uplink
	if uplink, err = spec.Uplink(); err != nil {
		return
	}

	if len(spec.Endpoints) <= 1 {
		for name, s := range spec.Endpoints {
			listeners = append(listeners, endpointListener{
				name:   fmt.Sprintf("%s:%s", spec.Name, name),
				uplink: uplink,
				spec:   s,
			})
			return
		}
	}

	if uplink.Proto == NetProtoUDP {
		err = ErrUDPMultiplexer
		return
	}

	var epNames []string
	var multiplexEndpoints = make(map[string]MultiplexHandler)
	for name, spec := range spec.Endpoints {
		epNames = append(epNames, name)
		if ep, ok := spec.Handler.(MultiplexHandler); !ok {
			err = fmt.Errorf("handler %s %w", spec.HandlerRef, ErrMultiplexingNotSupported)
			return
		} else {
			multiplexEndpoints[name] = ep
		}
	}

	sort.Strings(epNames)

	plainMux := cmux.New(uplink.Listener)
	tlsListener := plainMux.Match(cmux.TLS())
	tlsListener = tls.NewListener(tlsListener, tlsConfig)
	tlsMux := cmux.New(tlsListener)

	var tlsRequired = false

	for _, epName := range epNames {
		epSpec := spec.Endpoints[epName]
		var epMux = plainMux
		if epSpec.TLS {
			epMux = tlsMux
			tlsRequired = true
		}
		epListener := endpointListener{
			name: fmt.Sprintf("%s:%s", spec.Name, epName),
			uplink: Uplink{
				Proto:    NetProtoTCP,
				Listener: epMux.Match(multiplexEndpoints[epName].Matchers()...),
			},
			spec: epSpec,
		}

		listeners = append(listeners, epListener)
	}

	muxes = append(muxes, plainMux)

	if tlsRequired {
		muxes = append(muxes, tlsMux)
	} else {
		_ = tlsListener.Close()
	}

	return
}
