package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// 以下是 Redis Cluster 操作方法

// GetClusterClient 获取 Redis Cluster 客户端
func GetClusterClient() *redis.ClusterClient {
	return ClusterClient
}

// ClusterSet 设置键值对，带过期时间
func ClusterSet(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.Set(ctx, key, value, expiration).Err()
}

// ClusterGet 获取键值
func ClusterGet(ctx context.Context, key string) (string, error) {
	if ClusterClient == nil {
		return "", redis.Nil
	}
	return ClusterClient.Get(ctx, key).Result()
}

// ClusterDel 删除键
func ClusterDel(ctx context.Context, keys ...string) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.Del(ctx, keys...).Err()
}

// ClusterExists 检查键是否存在
func ClusterExists(ctx context.Context, keys ...string) (bool, error) {
	if ClusterClient == nil {
		return false, redis.Nil
	}
	result, err := ClusterClient.Exists(ctx, keys...).Result()
	return result > 0, err
}

// ClusterExpire 设置键的过期时间
func ClusterExpire(ctx context.Context, key string, expiration time.Duration) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.Expire(ctx, key, expiration).Err()
}

// ClusterHGet 获取哈希表中的字段值
func ClusterHGet(ctx context.Context, key, field string) (string, error) {
	if ClusterClient == nil {
		return "", redis.Nil
	}
	return ClusterClient.HGet(ctx, key, field).Result()
}

// ClusterHSet 设置哈希表字段值
func ClusterHSet(ctx context.Context, key string, values ...interface{}) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.HSet(ctx, key, values...).Err()
}

// ClusterHGetAll 获取哈希表中的所有字段和值
func ClusterHGetAll(ctx context.Context, key string) (map[string]string, error) {
	if ClusterClient == nil {
		return nil, redis.Nil
	}
	return ClusterClient.HGetAll(ctx, key).Result()
}

// ClusterHDel 删除哈希表中的字段
func ClusterHDel(ctx context.Context, key string, fields ...string) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.HDel(ctx, key, fields...).Err()
}

// ClusterLPush 向列表头部添加元素
func ClusterLPush(ctx context.Context, key string, values ...interface{}) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.LPush(ctx, key, values...).Err()
}

// ClusterRPush 向列表尾部添加元素
func ClusterRPush(ctx context.Context, key string, values ...interface{}) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.RPush(ctx, key, values...).Err()
}

// ClusterLRange 获取列表指定范围内的元素
func ClusterLRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	if ClusterClient == nil {
		return nil, redis.Nil
	}
	return ClusterClient.LRange(ctx, key, start, stop).Result()
}

// ClusterLPop 从列表中弹出头部元素
func ClusterLPop(ctx context.Context, key string) (string, error) {
	if ClusterClient == nil {
		return "", redis.Nil
	}
	return ClusterClient.LPop(ctx, key).Result()
}

// ClusterRPop 从列表中弹出尾部元素
func ClusterRPop(ctx context.Context, key string) (string, error) {
	if ClusterClient == nil {
		return "", redis.Nil
	}
	return ClusterClient.RPop(ctx, key).Result()
}

// ClusterSAdd 向集合添加元素
func ClusterSAdd(ctx context.Context, key string, members ...interface{}) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.SAdd(ctx, key, members...).Err()
}

// ClusterSMembers 获取集合中的所有元素
func ClusterSMembers(ctx context.Context, key string) ([]string, error) {
	if ClusterClient == nil {
		return nil, redis.Nil
	}
	return ClusterClient.SMembers(ctx, key).Result()
}

// ClusterSIsMember 判断元素是否是集合的成员
func ClusterSIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	if ClusterClient == nil {
		return false, redis.Nil
	}
	return ClusterClient.SIsMember(ctx, key, member).Result()
}

// ClusterSRem 从集合中删除元素
func ClusterSRem(ctx context.Context, key string, members ...interface{}) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.SRem(ctx, key, members...).Err()
}

// ClusterZAdd 向有序集合添加元素
func ClusterZAdd(ctx context.Context, key string, members ...redis.Z) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.ZAdd(ctx, key, members...).Err()
}

// ClusterZRange 获取有序集合指定范围内的元素
func ClusterZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	if ClusterClient == nil {
		return nil, redis.Nil
	}
	return ClusterClient.ZRange(ctx, key, start, stop).Result()
}

// ClusterZRem 从有序集合中删除元素
func ClusterZRem(ctx context.Context, key string, members ...interface{}) error {
	if ClusterClient == nil {
		return redis.Nil
	}
	return ClusterClient.ZRem(ctx, key, members...).Err()
}

// ClusterZScore 获取有序集合元素的分数
func ClusterZScore(ctx context.Context, key string, member string) (float64, error) {
	if ClusterClient == nil {
		return 0, redis.Nil
	}
	return ClusterClient.ZScore(ctx, key, member).Result()
}

// ClusterEval 执行Lua脚本
func ClusterEval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	if ClusterClient == nil {
		return nil, redis.Nil
	}
	return ClusterClient.Eval(ctx, script, keys, args...).Result()
}
