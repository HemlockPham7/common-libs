package nrtrace

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type config struct {
	AppName           string `envconfig:"NR_APP_NAME" default:"bookmark-service"`
	License           string `envconfig:"NR_LICENSE" default:""`
	LogForwardEnabled bool   `envconfig:"NR_LOG_FORWARD_ENABLED" default:"true"`
}

func newConfig(envPrefix string) (*config, error) {
	cfg := &config{}
	err := envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}

func NewClient(envPrefix string) (*newrelic.Application, error) {
	conf, err := newConfig(envPrefix)
	if err != nil {
		return nil, err
	}

	if conf.License == "" {
		return nil, errors.New("NR_LICENSE is required")
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(conf.AppName),
		newrelic.ConfigLicense(conf.License),
		newrelic.ConfigAppLogForwardingEnabled(conf.LogForwardEnabled),
	)
	if err != nil {
		return nil, err
	}

	return app, nil
}
