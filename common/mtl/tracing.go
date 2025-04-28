package mtl

import (
	"github.com/hertz-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithEnableTracing(true), // 确保启用链路追踪
	)
	return p
}
