package config

import (
	"time"
)

const (
	ConfigHomeDir = "$HOME/.mocosso"
	ConfigEtcDir  = "/etc/mocosso"

	ConfigType     = "yaml"
	ConfigFileName = "mocosso.yaml"

	ConfigEnvPrefix = "mocosso"
)

// Config is config structure.
type Config struct {
	Core     SectionCore     `yaml:"core"     mapstructure:"core"`
	Auth     SectionAuth     `yaml:"auth"     mapstructure:"auth"`
	Log      SectionLog      `yaml:"log"      mapstructure:"log"`
	Postgres SectionPostgres `yaml:"postgres" mapstructure:"postgres"`
	Redis    SectionRedis    `yaml:"redis"    mapstructure:"redis"`
	HTTP     SectionHTTP     `yaml:"grpc"     mapstructure:"http"`
	GRPC     SectionGRPC     `yaml:"http"     mapstructure:"grpc"`
}

// SectionCore is subsection of config with core options.
type SectionCore struct {
	Env             string        `yaml:"env"              mapstructure:"env"`
	Host            string        `yaml:"host"             mapstructure:"host"`
	ServiceName     string        `yaml:"service_name"     mapstructure:"service_name"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" mapstructure:"shutdown_timeout"`
}

// SectionAuth is subsection of config with auth options.
type SectionAuth struct {
	Token SectionToken `yaml:"token" mapstructure:"token"`
}

type SectionToken struct {
	JWTPublicKey    string        `yaml:"jwt_public_key"   mapstructure:"jwt_public_key"`
	JWTPrivateKey   string        `yaml:"jwt_private_key"  mapstructure:"jwt_private_key"`
	AccessLiveness  time.Duration `yaml:"access_liveness"  mapstructure:"access_liveness"`
	RefreshLiveness time.Duration `yaml:"refresh_liveness" mapstructure:"refresh_liveness"`
}

type SectionVerification struct {
	EmailVerifictionLiveness time.Duration `yaml:"email_verifiction_liveness" mapstructure:"email_verifiction_liveness"`
	PasswordResetLiveness    time.Duration `yaml:"password_reset_liveness"    mapstructure:"password_reset_liveness"`
}

// SectionLog is subsection of config with log options.
type SectionLog struct {
	Encoding      string `yaml:"encoding"   mapstructure:"encoding"`
	Level         string `yaml:"level"      mapstructure:"level"`
	AccessLogPath string `yaml:"access_log" mapstructure:"access_log"`
	ErrorLogPath  string `yaml:"error_log"  mapstructure:"error_log"`
}

// SectionPostgres is subsection of config with storage options.
type SectionPostgres struct {
	PostrgresURI string `yaml:"postgres_uri" mapstructure:"postgres_uri"`
}

// SectionRedis is subsection of config with cache options.
type SectionRedis struct {
	RedisURI string `yaml:"redis_uri" mapstructure:"redis_uri"`
}

// SectionHTTP is sub section of config with http server options.
type SectionHTTP struct {
	Enabled bool   `yaml:"enabled" mapstructure:"enabled"`
	Port    string `yaml:"port"    mapstructure:"port"`

	SSL      bool   `yaml:"ssl"       mapstructure:"ssl"`
	CertPath string `yaml:"cert_path" mapstructure:"cert_path"`
	KeyPath  string `yaml:"key_path"  mapstructure:"key_path"`
}

// SectionGRPC is subsection of config with grpc server options.
type SectionGRPC struct {
	Enabled bool   `yaml:"enabled" mapstructure:"enabled"`
	Port    string `yaml:"port"    mapstructure:"port"`

	SSL      bool   `yaml:"ssl"       mapstructure:"ssl"`
	CertPath string `yaml:"cert_path" mapstructure:"cert_path"`
	KeyPath  string `yaml:"key_path"  mapstructure:"key_path"`
}
