package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	ConfigHomeDir = "$HOME/.gosso"
	ConfigEtcDir  = "/etc/gosso"

	ConfigType     = "yaml"
	ConfigFileName = "gosso.conf"

	ConfigEnvPrefix = "gosso"
)

type Loader interface {
	Load()
}

// TODO: add connection timeouts and limitations to storage
// TODO: add memory and redis cache with options
var defaultConfig = []byte(`
core: 
  mode: "release" # default is "release", "local", "development"
  address: # ip address to bind, default is "any"

  ssl: false # default is false (disabled)
  cert_path: "cert.pem"
  key_path: "key.pem"

  shutdown_timeout: 30s # default is 30s

log:
  format: "json" # text or json
  level: "debug"  # default is debug
  path: "stdout" # log path default is stdout console, but could file

storage:
  database_uri: "postgres://user:password@localhost:5432/mixert?sslmode=disable"

cache:
  enabled: false
	type: "memory"

  redis:
    
  memory:

http:
  enabled: true
  port: "8080"

grpc:
  enabled: false
  port: "9090"
`)

// ConfigYaml is config structure.
type ConfigYaml struct {
	Core    SectionCore    `yaml:"core"`
	Log     SectionLog     `yaml:"log"`
	Storage SectionStorage `yaml:"storage"`
	Cache   SectionCache   `yaml:"cache"`
	HTTP    SectionHTTP    `yaml:"grpc"`
	GRPC    SectionGRPC    `yaml:"http"`
}

// Load method is implementation of loader interface.
func (cfg *ConfigYaml) Load() {
	(&cfg.Core).Load()
	(&cfg.Log).Load()
	(&cfg.Storage).Load()
	(&cfg.Cache).Load()
	(&cfg.HTTP).Load()
	(&cfg.GRPC).Load()
}

// SectionCore is subsection of config with core options.
type SectionCore struct {
	Mode    string `yaml:"mode"`
	Address string `yaml:"address"`

	SSL      bool   `yaml:"ssl"`
	CertPath string `yaml:"cert_path"`
	KeyPath  string `yaml:"key_path"`

	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

// Load method is implementation of loader interface.
func (sc *SectionCore) Load() {
	sc.Mode = viper.GetString("core.mode")
	sc.Address = viper.GetString("core.address")

	sc.SSL = viper.GetBool("core.ssl")
	sc.CertPath = viper.GetString("core.cert_path")
	sc.KeyPath = viper.GetString("core.key_path")

	sc.ShutdownTimeout = viper.GetDuration("core.shutdown_timeout")
}

// SectionLog is subsection of config with log options.
type SectionLog struct {
	Format  string `yaml:"format"`
	Level   string `yaml:"level"`
	LogPath string `yaml:"path"`
}

// Load method is implementation of loader interface.
func (sc *SectionLog) Load() {
	sc.Format = viper.GetString("log.format")
	sc.Level = viper.GetString("log.level")
	sc.LogPath = viper.GetString("log.path")
}

// SectionLog is subsection of config with storage options.
type SectionStorage struct {
	URI string `yaml:"uri"`
}

// Load method is implementation of loader interface.
func (sc *SectionStorage) Load() {
	sc.URI = viper.GetString("storage.uri")
}

// SectionLog is subsection of config with cache options.
type SectionCache struct{}

// Load method is implementation of loader interface.
func (sc *SectionCache) Load() {}

// SectionHTTP is sub section of config with http server options.
type SectionHTTP struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
}

// Load method is implementation of loader interface.
func (sc *SectionHTTP) Load() {
	sc.Enabled = viper.GetBool("http.enabled")
	sc.Port = viper.GetString("http.port")
}

// SectionGRPC is subsection of config with grpc server options.
type SectionGRPC struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
}

// Load method is implementation of loader interface.
func (sc *SectionGRPC) Load() {
	sc.Enabled = viper.GetBool("grpc.enabled")
	sc.Port = viper.GetString("grpc.port")
}

// LoadConfig function load config from viper.
func LoadConfig(confPath ...string) (conf *ConfigYaml, err error) {
	const op = "config.LoadConfig"

	conf = &ConfigYaml{}
	if err := LoadViper(confPath...); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	conf.Load()
	return conf, nil
}

// MustLoadConfig function load config from viper.
// If any of errors occurred It panics.
func MustLoadConfig(confPath ...string) (conf *ConfigYaml) {
	const op = "config.MustLoadConfig"

	conf = &ConfigYaml{}
	if err := LoadViper(confPath...); err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

	conf.Load()
	return conf
}

// LoadViper function load config in viper from file and read env vars that match.
func LoadViper(confPath ...string) (err error) {
	viper.SetConfigType(ConfigType)

	// setup environment for viper
	viper.SetEnvPrefix(ConfigEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if len(confPath) > 0 && confPath[0] != "" {
		content, err := os.ReadFile(confPath[0])
		if err != nil {
			return fmt.Errorf("failed to read config file: %w", err)
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return fmt.Errorf("failed to load config in viper from confPath: %w", err)
		}
	} else {
		viper.AddConfigPath(ConfigEtcDir)
		viper.AddConfigPath(ConfigHomeDir)
		viper.AddConfigPath("./configs")

		viper.SetConfigName(ConfigFileName)

		if err := viper.ReadInConfig(); err == nil {
			fmt.Printf("[%s - config file used]\n", viper.ConfigFileUsed())
		} else if err := viper.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
			return fmt.Errorf("failed to load default config in viper: %w", err)
		}
	}

	return nil
}
