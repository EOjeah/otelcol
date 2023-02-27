package spanattributeprocessor // import github.com/eojeah/otelcol/collector/processor/spanattributeprocessor

import (
	"errors"

	"go.opentelemetry.io/collector/component"
)

type Config struct {
	OsType string `mapstructure:"os_type"`
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error {
	osType := cfg.OsType
	if osType != "linux" {
		return errors.New("os_type must be equal to linux")
	}
	return nil
}
