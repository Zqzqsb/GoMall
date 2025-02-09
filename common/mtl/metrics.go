package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"

	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Register *prometheus.Registry

// InitMetrics 初始化并配置 Prometheus 监控指标系统
// 参数:
//   - serviceName: 服务名称，用于在 Consul 中标识该服务
//   - metricsPort: metrics 服务监听的端口，例如 ":8080"
//   - registeryAddr: Consul 注册中心的地址
//
// 功能:
//   1. 创建并配置 Prometheus Registry
//   2. 注册基础的 Go 运行时和进程收集器
//   3. 将 metrics 端点注册到 Consul 以便服务发现
//   4. 暴露 /metrics 端点用于 Prometheus 抓取
//   5. 优雅关闭时自动从 Consul 注销服务

func InitMetrics(serviceName, metricsPort, registeryAddr string) {
	Register = prometheus.NewRegistry()
	Register.MustRegister(collectors.NewGoCollector())
	Register.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	r, _ := consul.NewConsulRegister(registeryAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	registerInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	r.Register(registerInfo)
	server.RegisterShutdownHook(func() {
		r.Deregister(registerInfo)
	})
	http.Handle("/metrics", promhttp.HandlerFor(Register, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil)
}
