package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

type CommonClientSuite struct {
	CurrentServiceName string
	RegisteryAddr      string
}

func (s CommonClientSuite) Options() []client.Option {

	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithTransportProtocol(transport.GRPC),
		// 使用 OpenTelemetry 的链路追踪
		client.WithSuite(tracing.NewClientSuite()),
	}
	return opts
}
