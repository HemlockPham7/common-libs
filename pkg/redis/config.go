package redis

import "github.com/kelseyhightower/envconfig"

type config struct {
	Address  string `default:"localhost:6379" envconfig:"REDIS_ADDR"`
	Password string `default:"" envconfig:"REDIS_PWD"`
	DB       int    `default:"0" envconfig:"REDIS_DB"`
}

// newConfig creates a new config
func newConfig(envPrefix string) (*config, error) {
	cfg := &config{}
	err := envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
