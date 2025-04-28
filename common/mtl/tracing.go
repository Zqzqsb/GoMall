package mtl

import (
	// ...
	"github.com/hertz-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
		provider.WithExportEndpoint("localhost:4317"), // 添加 OpenTelemetry Collector 导出端点
	)
	return p
}
