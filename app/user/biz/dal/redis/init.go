package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"zqzqsb.com/gomall/app/user/conf"
)

var (
	RedisClient   *redis.Client
	ClusterClient *redis.ClusterClient
)

// Init 初始化 Redis 客户端
func Init() {
	// 初始化单节点 Redis 客户端
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	// 如果启用了 Redis Cluster，则初始化 Cluster 客户端
	if conf.GetConf().RedisCluster.Enabled {
		if err := initRedisCluster(); err != nil {
			log.Printf("Failed to initialize Redis Cluster: %v", err)
			panic(err)
		}
		log.Println("Redis Cluster initialized successfully")
	}
}

// Close 关闭 Redis 连接
func Close() {
	if RedisClient != nil {
		RedisClient.Close()
	}
	
	// 如果启用了 Redis Cluster，则关闭 Cluster 连接
	if conf.GetConf().RedisCluster.Enabled && ClusterClient != nil {
		closeRedisCluster()
	}
}

// initRedisCluster 初始化 Redis Cluster 客户端
func initRedisCluster() error {
	// 创建 Redis Cluster 客户端
	ClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:           conf.GetConf().RedisCluster.Addrs,
		Username:        conf.GetConf().RedisCluster.Username,
		Password:        conf.GetConf().RedisCluster.Password,
		MaxRetries:      conf.GetConf().RedisCluster.MaxRetries,
		MinRetryBackoff: conf.GetConf().RedisCluster.MinRetryBackoff,
		MaxRetryBackoff: conf.GetConf().RedisCluster.MaxRetryBackoff,
		RouteByLatency:  conf.GetConf().RedisCluster.RouteByLatency,
		RouteRandomly:   conf.GetConf().RedisCluster.RouteRandomly,
	})

	// 测试连接
	if err := ClusterClient.Ping(context.Background()).Err(); err != nil {
		return err
	}

	return nil
}

// closeRedisCluster 关闭 Redis Cluster 客户端连接
func closeRedisCluster() error {
	if ClusterClient != nil {
		return ClusterClient.Close()
	}
	return nil
}
