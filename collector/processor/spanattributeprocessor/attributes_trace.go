package spanattributeprocessor

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type spanAttributesProcessor struct {
	logger *zap.Logger
	//attrProc *attraction.AttrProc
	//skipExpr expr.BoolExpr[ottlspan.TransformContext]
}

func newSpanAttributesProcessor(logger *zap.Logger) *spanAttributesProcessor {
	return &spanAttributesProcessor{
		logger: logger,
	}
}

func (sap *spanAttributesProcessor) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (sap *spanAttributesProcessor) Shutdown(_ context.Context) error {
	return nil
}

func (a *spanAttributesProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	rss := td.ResourceSpans()
	for i := 0; i < rss.Len(); i++ {
		rs := rss.At(i)
		ilss := rs.ScopeSpans()
		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)
			spans := ils.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				//span.Attributes().PutStr(linuxcpu.cpuPercent, "testValue")
				span.Attributes().PutStr("cpu.percent", getCpuPercent())
			}
		}
	}
	return td, nil
}

func getCpuPercent() string {
	cpuPercent, err := cpu.Percent(0, false)
	//cpuTime, err = cpu.Times(false)
	if err != nil {
		fmt.Println(err)
		return "unknown"
	}
	cpuString := fmt.Sprintf("%.2f", cpuPercent[0])
	//fmt.Println("CPU utilization: ", cpuString+"%")
	//fmt.Println("CPU utilization: ", cpuPercent)
	return cpuString
}
