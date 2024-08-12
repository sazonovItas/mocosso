package config

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMissingFile(t *testing.T) {
	_, err := LoadConfig("test")
	assert.NotNil(t, err, "should be error file missing: %v", err)
}

func TestLoadDefaultConfig(t *testing.T) {
	_, err := LoadConfig()
	assert.Nil(t, err, "default config should be right: %s", err)
}

func TestLoadConfigFromEnv(t *testing.T) {
	envPrefix := strings.ToUpper(ConfigEnvPrefix)

	// TODO: add cache section
	// TODO: add error test
	testCases := []struct {
		name    string
		envs    map[string]string
		want    interface{}
		got     Loader
		wantErr bool
	}{
		{
			name: "core section",
			envs: map[string]string{
				envPrefix + "_CORE_MODE":             "development",
				envPrefix + "_CORE_ADDRESS":          "localhost",
				envPrefix + "_CORE_SSL":              "true",
				envPrefix + "_CORE_CERT_PATH":        "localhost.cert",
				envPrefix + "_CORE_KEY_PATH":         "localhost.key",
				envPrefix + "_CORE_SHUTDOWN_TIMEOUT": "10s",
			},
			want: &SectionCore{
				Mode:            "development",
				Address:         "localhost",
				SSL:             true,
				CertPath:        "localhost.cert",
				KeyPath:         "localhost.key",
				ShutdownTimeout: 10 * time.Second,
			},
			got:     &SectionCore{},
			wantErr: false,
		},
		{
			name: "log section",
			envs: map[string]string{
				envPrefix + "_LOG_FORMAT": "console",
				envPrefix + "_LOG_LEVEL":  "error",
				envPrefix + "_LOG_PATH":   "error.log",
			},
			want: &SectionLog{
				Format:  "console",
				Level:   "error",
				LogPath: "error.log",
			},
			got:     &SectionLog{},
			wantErr: false,
		},
		{
			name: "storage section",
			envs: map[string]string{
				envPrefix + "_STORAGE_URI": "postgres://test_user:test_password@localhost:3030/mixery",
			},
			want: &SectionStorage{
				URI: "postgres://test_user:test_password@localhost:3030/mixery",
			},
			got:     &SectionStorage{},
			wantErr: false,
		},
		{
			name: "http section",
			envs: map[string]string{
				envPrefix + "_HTTP_ENABLED": "false",
				envPrefix + "_HTTP_PORT":    "3030",
			},
			want: &SectionHTTP{
				Enabled: false,
				Port:    "3030",
			},
			got:     &SectionHTTP{},
			wantErr: false,
		},
		{
			name: "grpc section",
			envs: map[string]string{
				envPrefix + "_GRPC_ENABLED": "true",
				envPrefix + "_GRPC_PORT":    "4040",
			},
			want: &SectionGRPC{
				Enabled: true,
				Port:    "4040",
			},
			got:     &SectionGRPC{},
			wantErr: false,
		},
		{
			name: "log section",
			envs: map[string]string{
				envPrefix + "_LOG_FORMAT": "console",
				envPrefix + "_LOG_LEVEL":  "error",
				envPrefix + "_LOG_PATH":   "error.log",
			},
			want: &SectionLog{
				Format:  "text",
				Level:   "error",
				LogPath: "error.log",
			},
			got:     &SectionLog{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for env, value := range tc.envs {
				os.Setenv(env, value)
			}
			defer os.Clearenv()

			if err := LoadViper(); err != nil {
				panic(err)
			}

			tc.got.Load()
			if tc.wantErr {
				assert.NotEqual(
					t,
					tc.want,
					tc.got,
					"should not be equal want %v, got %v, want error %v",
					tc.want,
					tc.got,
					tc.wantErr,
				)
			} else {
				assert.Equal(
					t,
					tc.want,
					tc.got,
					"should be equal want %v, got %v, want error %v",
					tc.want,
					tc.got,
					tc.wantErr)
			}
		})
	}
}
