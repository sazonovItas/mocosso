package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/iamolegga/enviper"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type (
	Option interface {
		Apply(cfg *enviper.Enviper)
	}

	optionFunc func(cfg *enviper.Enviper)
)

func (of optionFunc) Apply(cfg *enviper.Enviper) {
	of(cfg)
}

func WithConfigPaths(paths ...string) optionFunc {
	return func(cfg *enviper.Enviper) {
		for _, path := range paths {
			cfg.AddConfigPath(path)
		}
	}
}

func WithConfigType(cfgType string) optionFunc {
	return func(cfg *enviper.Enviper) {
		cfg.SetConfigType(cfgType)
	}
}

func WithConfigName(name string) optionFunc {
	return func(cfg *enviper.Enviper) {
		cfg.SetConfigName(name)
	}
}

func WithEnvs(envPrefix string) optionFunc {
	return func(cfg *enviper.Enviper) {
		cfg.SetEnvPrefix(envPrefix)
		cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		cfg.AutomaticEnv()
	}
}

func WithDefaults(defaults map[string]any) optionFunc {
	return func(cfg *enviper.Enviper) {
		if defaults == nil {
			return
		}

		for k, v := range defaults {
			cfg.SetDefault(k, v)
		}
	}
}

// TODO: do something with returning error.
// WithFlagBindings binds flags to viper config from flag set.
func WithFlagBindings(
	flagSet *pflag.FlagSet,
	bindings map[string]func(flagSet *pflag.FlagSet) *pflag.Flag,
) optionFunc {
	return func(cfg *enviper.Enviper) {
		if bindings == nil {
			return
		}

		for key, selector := range bindings {
			if err := cfg.BindPFlag(key, selector(flagSet)); err != nil {
				panic(err)
			}
		}
	}
}

func Load(cfg any, configPath string, ignoreMissingFile bool, options ...Option) error {
	viperCfg := enviper.New(viper.New())
	for _, option := range options {
		option.Apply(viperCfg)
	}

	if configPath != "" {
		_, err := os.Stat(configPath)
		if err == nil {
			viperCfg.SetConfigFile(configPath)
		}
	}

	if err := viperCfg.ReadInConfig(); err != nil {
		viperErr := new(viper.ConfigFileNotFoundError)
		if !errors.As(err, viperErr) || !ignoreMissingFile {
			return err
		}
	}

	if err := viperCfg.Unmarshal(cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
