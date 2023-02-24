package spanattributeprocessor

import (
	"fmt"
	"go.opentelemetry.io/collector/config"
	"strings"
)

type Config struct {
	config.ProcessorSettings `mapstructure:",squash"`
	OsType                   string `mapstructure:"os_type"`
}

func (cfg *Config) Validate() error {
	osType := cfg.OsType
	if strings.ToLower(osType) != "linux" {
		return fmt.Errorf("os_type must be equal to linux")
	}
	return nil
}
