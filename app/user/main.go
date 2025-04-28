package main

import (
	"context"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"zqzqsb.com/gomall/app/user/biz/dal"
	"zqzqsb.com/gomall/app/user/conf"
	"zqzqsb.com/gomall/app/user/kitex_gen/user/userservice"
	"zqzqsb.com/gomall/common/mtl"
	"zqzqsb.com/gomall/common/serversuite"
)

var (
	// rigister server to consul
	serviceID      = "user-rpc-001"
	serviceName    = "user-service-rpc"
	serviceAddress = "192.168.110.112"
	servicePort    = 8888
	consulAddr     = "127.0.0.1:8500"
)

func main() {
	err := godotenv.Load()
	mtl.InitMetrics(serviceName, conf.GetConf().Kitex.MetricsPort, consulAddr)
	p := mtl.InitTracing(serviceName)
	defer p.Shutdown(context.Background())

	if err != nil {
		klog.Error(err.Error())
	}

	dal.Init()
	opts := kitexInit()
	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	opts = append(opts, server.
		WithTransHandlerFactory(&mixTransHandlerFactory{nil}))

	opts = append(opts, server.
		WithTransHandlerFactory(&mixTransHandlerFactory{nil}))

	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))
	opts = append(opts, server.WithSuite(&serversuite.CommonServerSuite{
		CurrentServiceName: serviceName,
		RegisteryAddr:      consulAddr,
	}))
	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	registerServiceWithConsul(serviceID, serviceName, serviceAddress, servicePort)

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
