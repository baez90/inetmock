package app

import (
	"reflect"
	"testing"

	"gitlab.com/inetmock/inetmock/internal/endpoint"
)

func Test_config_ReadConfig2(t *testing.T) {
	cfg := CreateConfig()
	_ = cfg.ReadConfigString(`
listeners:
  tcp_80:
    name: ''
    protocol: tcp
    listenAddress: ''
    port: 80
  tcp_443:
    name: ''
    protocol: tcp
    listenAddress: ''
    port: 443
endpoints:
  plainHttp:
    handler: http_mock
    listeners:
      - tcp_80
    options: {}
  https:
    handler: http_mock
    listeners:
      - tcp_443
    options:
      tls: true
`, "yaml")

	t.Log(cfg)
}

func Test_config_ReadConfig(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name          string
		args          args
		want          config
		wantEndpoints map[string]endpoint.MetaSpec
		wantListeners map[endpoint.ListenerReference]endpoint.ListenerSpec
		wantErr       bool
	}{
		{
			name: "Test endpoints config",
			args: args{
				//language=yaml
				config: `
listeners:
  tcp_80:
    name: ''
    protocol: tcp
    listenAddress: ''
    port: 80
  tcp_443:
    name: ''
    protocol: tcp
    listenAddress: ''
    port: 443
endpoints:
  plainHttp:
    handler: http_mock
    listeners:
      - tcp_80
    options: {}
  https:
    handler: http_mock
    listeners:
      - tcp_443
    options:
      tls: true
`,
			},
			wantEndpoints: map[string]endpoint.MetaSpec{
				"plainhttp": {
					Handler:   "http_mock",
					Listeners: []endpoint.ListenerReference{"tcp_80"},
					Options:   nil,
				},
				"https": {
					Handler:   "http_mock",
					Listeners: []endpoint.ListenerReference{"tcp_443"},
					Options: map[string]interface{}{
						"tls": true,
					},
				},
			},
			wantListeners: map[endpoint.ListenerReference]endpoint.ListenerSpec{
				"tcp_80": {
					Name:     "",
					Protocol: "tcp",
					Address:  "",
					Port:     80,
				},
				"tcp_443": {
					Name:     "",
					Protocol: "tcp",
					Address:  "",
					Port:     443,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := CreateConfig()
			if err := cfg.ReadConfigString(tt.args.config, "yaml"); (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.wantEndpoints, cfg.MetaSpecs()) {
				t.Errorf("want = %v, got = %v", tt.wantEndpoints, cfg.MetaSpecs())
			}

			if !reflect.DeepEqual(tt.wantListeners, cfg.ListenerSpecs()) {
				t.Errorf("want = %v, got = %v", tt.wantListeners, cfg.ListenerSpecs())
			}
		})
	}
}
