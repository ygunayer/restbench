package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type HttpConfig struct {
	Hostname string `mapstructure:"HTTP__HOST"`
	Port     int16  `mapstructure:"HTTP__PORT"`
}

func (cfg *HttpConfig) GetListenAddr() string {
	return fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port)
}

type DatabaseConfig struct {
	Url string `mapstructure:"DATABASE__URL"`
}

type CassandraConfig struct {
	Hosts    []string `mapstructure:"CASSANDRA__HOSTS"`
	Keyspace string   `mapstructure:"CASSANDRA__KEYSPACE"`
}

type Config struct {
	Http      HttpConfig      `mapstructure:",squash"`
	Database  DatabaseConfig  `mapstructure:",squash"`
	Cassandra CassandraConfig `mapstructure:",squash"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	env, err := godotenv.Read()
	if err != nil {
		return nil, err
	}

	var cfg Config

	decoderConfig := mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &cfg,
	}

	decoder, err := mapstructure.NewDecoder(&decoderConfig)
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(env); err != nil {
		return nil, err
	}

	return &cfg, nil
}
