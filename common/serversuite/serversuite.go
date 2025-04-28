package serversuite

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"zqzqsb.com/gomall/common/mtl"
)

type CommonServerSuite struct {
	CurrentServiceName string
}

func (s *CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithTracer(prometheus.NewServerTracer("", "",
			prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Register))),
		// 使用 OpenTelemetry 的链路追踪
		server.WithSuite(tracing.NewServerSuite()),
	}
	return opts
}
