package spanattributeprocessor

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
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

func createTracesProcessor(ctx context.Context, set processor.CreateSettings, cfg component.Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	return processorhelper.NewTracesProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		newSpanAttributesProcessor(set.Logger).processTraces,
		processorhelper.WithCapabilities(processorCapabilities))
}

func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		processor.WithTraces(createTracesProcessor, stability))
}
