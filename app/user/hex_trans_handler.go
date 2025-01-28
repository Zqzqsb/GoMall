package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"regexp"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	hertzServer "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/detection"
	"github.com/cloudwego/kitex/pkg/remote/trans/netpoll"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/cors"
	"zqzqsb.com/gomall/app/user/biz/router"
	mw "zqzqsb.com/gomall/app/user/biz/router/middleware"
)

type mixTransHandlerFactory struct {
	originFactory remote.ServerTransHandlerFactory
}

type transHandler struct {
	remote.ServerTransHandler
}

// SetInvokeHandleFunc is used to set invoke handle func.
func (t *transHandler) SetInvokeHandleFunc(inkHdlFunc endpoint.Endpoint) {
	t.ServerTransHandler.(remote.InvokeHandleFuncSetter).SetInvokeHandleFunc(inkHdlFunc)
}

func (m mixTransHandlerFactory) NewTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	var kitexOrigin remote.ServerTransHandler
	var err error

	if m.originFactory != nil {
		kitexOrigin, err = m.originFactory.NewTransHandler(opt)
	} else {
		// if no customized factory just use the default factory under detection pkg.
		kitexOrigin, err = detection.NewSvrTransHandlerFactory(netpoll.NewSvrTransHandlerFactory(), nphttp2.NewSvrTransHandlerFactory()).NewTransHandler(opt)
	}
	if err != nil {
		return nil, err
	}
	return &transHandler{ServerTransHandler: kitexOrigin}, nil
}

var httpReg = regexp.MustCompile(`^(?:GET |POST|PUT|DELE|HEAD|OPTI|CONN|TRAC|PATC)$`)

func (t *transHandler) OnRead(ctx context.Context, conn net.Conn) error {
	c, ok := conn.(network.Conn)
	if ok {
		pre, _ := c.Peek(4)
		if httpReg.Match(pre) {
			klog.Info("using Hertz to process request")
			err := hertzEngine.Serve(ctx, c)
			if err != nil {
				err = fmt.Errorf("HERTZ: %w", err)
			}
			return err
		}
	}
	return t.ServerTransHandler.OnRead(ctx, conn)
}

func initHertz() *route.Engine {
	h := hertzServer.New(hertzServer.WithIdleTimeout(0))
	log.Println("init hertz")
	
	// 配置 CORS
	h.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET"},
		AllowHeaders:    []string{"*"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))


	// 注册 session 和 csrf
	mw.InitJwt()
	mw.InitSession(h)
	mw.InitCSRF(h)

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	router.GeneratedRegister(h)
	
	if err := h.Engine.Init(); err != nil {
		panic(err)
	}
	if err := h.Engine.MarkAsRunning(); err != nil {
		panic(err)
	}
	
	return h.Engine
}

func registerServiceWithConsul(serviceID, serviceName, serviceAddress string, servicePort int) {
	// 配置 Consul 客户端
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)

	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// 定义服务注册信息
	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: serviceAddress,
		Port:    servicePort,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/ping", serviceAddress, servicePort),
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	// 注册服务
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Failed to register service with Consul: %v", err)
	}
	log.Printf("Service %s registered with Consul", serviceName)
}

var hertzEngine *route.Engine

func init() {

	hertzEngine = initHertz()
	serviceID := "user-http-001"
	serviceName := "user-service-http"
	serviceAddress := "192.168.110.112"
	servicePort := 8888

	registerServiceWithConsul(serviceID, serviceName, serviceAddress, servicePort)
}
