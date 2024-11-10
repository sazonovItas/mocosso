//go:build unit
// +build unit

package app

import (
	"reflect"
	"testing"
	"time"

	"github.com/sazonovItas/mocosso/pkg/config"
)

func TestConfig(t *testing.T) {
	type args struct {
		configPath        string
		ignoreMissingFile bool
		options           []config.Option
	}

	tests := []struct {
		name    string
		envs    map[string]string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Test setup default values",
			args: args{
				ignoreMissingFile: true,
				options: []config.Option{
					config.WithDefaults(map[string]any{
						"core.env":              "local",
						"core.host":             "localhost",
						"core.service_name":     "mocosso",
						"core.shutdown_timeout": 15 * time.Second,
					}),
				},
			},
			want: Config{
				Core: SectionCore{
					Env:             "local",
					Host:            "localhost",
					ServiceName:     "mocosso",
					ShutdownTimeout: 15 * time.Second,
				},
			},
		},
		{
			name: "Test load from environment variables",
			envs: map[string]string{
				"MOCOSSO_AUTH_TOKEN_JWT_PUBLIC_KEY":   "public_key",
				"MOCOSSO_AUTH_TOKEN_JWT_PRIVATE_KEY":  "private_key",
				"MOCOSSO_AUTH_TOKEN_ACCESS_LIVENESS":  "15m",
				"MOCOSSO_AUTH_TOKEN_REFRESH_LIVENESS": "720h",
			},
			args: args{
				ignoreMissingFile: true,
				options: []config.Option{
					config.WithEnvs(ConfigEnvPrefix),
				},
			},
			want: Config{
				Auth: SectionAuth{
					Token: SectionToken{
						JWTPublicKey:    "public_key",
						JWTPrivateKey:   "private_key",
						AccessLiveness:  15 * time.Minute,
						RefreshLiveness: 30 * 24 * time.Hour,
					},
				},
			},
		},
		{
			name: "Test environment variables covers defaults",
			envs: map[string]string{
				"MOCOSSO_AUTH_TOKEN_JWT_PUBLIC_KEY":  "public_key",
				"MOCOSSO_AUTH_TOKEN_JWT_PRIVATE_KEY": "private_key",
				"MOCOSSO_AUTH_TOKEN_ACCESS_LIVENESS": "15m",
			},
			args: args{
				ignoreMissingFile: true,
				options: []config.Option{
					config.WithDefaults(map[string]any{
						"core.env":                    "local",
						"core.host":                   "localhost",
						"auth.token.access_liveness":  30 * time.Minute,
						"auth.token.refresh_liveness": 720 * time.Hour,
					}),
					config.WithEnvs(ConfigEnvPrefix),
				},
			},
			want: Config{
				Core: SectionCore{
					Env:  "local",
					Host: "localhost",
				},
				Auth: SectionAuth{
					Token: SectionToken{
						JWTPublicKey:    "public_key",
						JWTPrivateKey:   "private_key",
						AccessLiveness:  15 * time.Minute,
						RefreshLiveness: 30 * 24 * time.Hour,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envs {
				t.Setenv(key, value)
			}

			var cfg Config
			if err := config.Load(&cfg, tt.args.configPath, tt.args.ignoreMissingFile, tt.args.options...); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.want, cfg) {
				t.Errorf("Load() config = %v, want = %v", cfg, tt.want)
			}
		})
	}
}
