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

// newConfig loads New Relic configuration from environment variables
// using the specified environment variable prefix.
//
// Parameters:
//   - envPrefix: the prefix used to load New Relic configuration values.
//
// Returns:
//   - The loaded configuration, or an error if the configuration cannot be processed.
func newConfig(envPrefix string) (*config, error) {
	cfg := &config{}
	err := envconfig.Process(envPrefix, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}

// NewClient creates a New Relic application using configuration loaded from environment variables.
//
// The New Relic license is required. An error is returned if the license is missing
// or the application cannot be initialized.
//
// Parameters:
//   - envPrefix: the prefix used to load New Relic configuration values.
//
// Returns:
//   - A configured New Relic application, or an error if the configuration is invalid
//     or the application cannot be initialized.
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
