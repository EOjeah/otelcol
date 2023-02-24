package spanattributeprocessor

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

const (
	typeStr       = "spanattributeprocessor"
	defaultOsType = "linux"
)

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(typeStr)),
		OsType:            defaultOsType,
	}
}

func createTracesProcessor(_ context.Context, params component.ProcessorCreateSettings, baseCfg config.Processor, consumer consumer.Traces) (component.TracesProcessor, error) {
	return nil, nil
}

func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
		component.WithTracesProcessor(createTracesProcessor))
}
