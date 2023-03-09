package spanattributeprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

const (
	typeStr       = "spanattributeprocessor"
	defaultOsType = "linux"
	stability     = component.StabilityLevelDevelopment
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

func createDefaultConfig() component.Config {
	return &Config{
		OsType: defaultOsType,
	}
}

func createTracesProcessor(_ context.Context, set processor.CreateSettings, cfg component.Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	return nil, nil
}

func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, stability))
}
