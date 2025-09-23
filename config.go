package main

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Web   WebConfig   `yaml:"API"`
	Debug DebugConfig `yaml:"DEBUG"`
}

type WebConfig struct {
	Host            string        `validate:"required" yaml:"HOST"`
	ReadTimeout     time.Duration `validate:"required" yaml:"READ_TIMEOUT"`
	IdleTimeout     time.Duration `validate:"required" yaml:"IDLE_TIMEOUT"`
	WriteTimeout    time.Duration `validate:"required" yaml:"WRITE_TIMEOUT"`
	ShutdownTimeout time.Duration `validate:"required" yaml:"SHUTDOWN_TIMEOUT"`
}
type DebugConfig struct {
	Host            string        `validate:"required" yaml:"HOST"`
	ReadTimeout     time.Duration `validate:"required" yaml:"READ_TIMEOUT"`
	IdleTimeout     time.Duration `validate:"required" yaml:"IDLE_TIMEOUT"`
	WriteTimeout    time.Duration `validate:"required" yaml:"WRITE_TIMEOUT"`
	ShutdownTimeout time.Duration `validate:"required" yaml:"SHUTDOWN_TIMEOUT"`
}

func loadConfig(ctx context.Context) (Config, error) {
	var config Config
	var configData []byte
	var err error

	env := os.Getenv("env")
	if env == "local" {
		// Cargar configuración desde archivo local
		configData, err = os.ReadFile("config.yaml")

		if err != nil {
			return Config{}, err
		}
	}

	// Parsear YAML (tanto para local como para secret manager)
	if err := yaml.Unmarshal(configData, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

// LoadConfig reads configuration from file or environment variables.
func readConfig() (Config, error) {
	viper.AddConfigPath("./app/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	// if err := modelvalidator.Check(&cfg, false); err != nil {
	// 	return Config{}, err
	// }

	return cfg, nil
}
