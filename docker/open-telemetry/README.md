# GoMall OpenTelemetry 链路追踪

> 本文档介绍了 GoMall 项目中 OpenTelemetry 链路追踪的架构、配置和使用方法。

## 1. 架构概述

> OpenTelemetry 提供了统一的可观测性框架，用于收集、处理和导出分布式系统的链路追踪数据。

GoMall 项目使用 OpenTelemetry 实现了分布式链路追踪，主要组件包括：

- **OpenTelemetry Collector**：收集、处理和导出链路追踪数据
- **Jaeger**：主要的链路追踪可视化工具
- **Zipkin**：备选的链路追踪可视化工具
- **Prometheus**：用于收集和存储指标数据（复用现有的 Prometheus 实例）
- **Grafana**：用于可视化指标和链路追踪数据（复用现有的 Grafana 实例）

## 2. 部署与配置

> 通过 Docker Compose 可以快速部署 OpenTelemetry 链路追踪后端服务。

### 2.1 启动服务

```bash
cd /path/to/GomallBackend/docker/open-telemetry
docker-compose up -d
```

### 2.2 访问链路追踪界面

- Jaeger UI: http://localhost:16686
- Zipkin UI: http://localhost:9411

### 2.3 配置说明

- **otel-collector-config.yaml**：OpenTelemetry Collector 的配置文件
  - 配置了接收器、处理器和导出器
  - 支持导出到 Jaeger、Zipkin 和 Prometheus

## 3. 应用集成

> 在 GoMall 项目中，链路追踪已集成到服务套件中，无需额外配置。

### 3.1 服务端集成

服务端通过 `serversuite` 包集成了 OpenTelemetry 链路追踪：

```go
// 在 main 函数中初始化 OpenTelemetry
mtl.InitTracing(serviceName)

// 使用服务套件
opts := kitexInit()
svr := userservice.NewServer(new(UserServiceImpl), opts...)
```

### 3.2 客户端集成

客户端通过 `clientsuite` 包集成了 OpenTelemetry 链路追踪：

```go
// 在 main 函数中初始化 OpenTelemetry
mtl.InitTracing(serviceName)

// 使用客户端套件
client, err := userservice.NewClient(
    "user-service",
    client.WithHostPorts("127.0.0.1:8888"),
    client.WithSuite(clientsuite.NewClientSuite("user-client")),
)
```

## 4. 最佳实践

> 以下是在 GoMall 项目中使用 OpenTelemetry 链路追踪的最佳实践。

### 4.1 服务命名规范

- 服务名应该使用小写字母和连字符，例如 `user-service`
- 客户端名应该使用小写字母和连字符，并添加 `-client` 后缀，例如 `user-service-client`

### 4.2 自定义 Span

在需要追踪的代码块中添加自定义 Span：

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
)

func SomeFunction() {
    ctx, span := otel.Tracer("my-service").Start(ctx, "SomeFunction")
    defer span.End()
    
    // 添加自定义属性
    span.SetAttributes(attribute.String("key", "value"))
    
    // 执行业务逻辑
    // ...
}
```

### 4.3 错误处理

在 Span 中记录错误：

```go
import (
    "go.opentelemetry.io/otel/codes"
)

func SomeFunction() error {
    ctx, span := otel.Tracer("my-service").Start(ctx, "SomeFunction")
    defer span.End()
    
    // 执行业务逻辑
    err := doSomething()
    if err != nil {
        span.SetStatus(codes.Error, err.Error())
        span.RecordError(err)
        return err
    }
    
    return nil
}
```

## 5. 故障排查

> 当链路追踪出现问题时，可以通过以下步骤进行排查。

### 5.1 检查服务状态

```bash
docker-compose ps
```

### 5.2 查看日志

```bash
docker-compose logs otel-collector
docker-compose logs jaeger
docker-compose logs zipkin
```

### 5.3 常见问题

1. **链路追踪数据不显示**
   - 检查 OpenTelemetry Collector 是否正常运行
   - 检查应用是否正确初始化了 OpenTelemetry
   - 检查网络连接是否正常

2. **链路追踪数据不完整**
   - 检查所有服务是否都集成了 OpenTelemetry
   - 检查上下文传播是否正确
   - 检查 Span 是否正确结束
