package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type Driver struct {
}

func NewDriver() *Driver {
	return &Driver{}
}

// Single
// @Description: 获取单节点redis连接
// @receiver d
// @param c 单节点配置
// @return *redis.Client
func (d *Driver) Single(c *redis.Options) *redis.Client {
	if len(c.Addr) == 0 {
		panic("config addr lost in single redis")
	}

	rdb := redis.NewClient(c)

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	log.Println("redis服务连接成功，当前模式【单机】")

	return rdb
}

// Cluster
// @Description: redis集群连接
// @receiver d
// @param c
// @return *redis.ClusterClient
func (d *Driver) Cluster(c *redis.ClusterOptions) *redis.ClusterClient {
	rdb := redis.NewClusterClient(c)

	ctx := context.Background()
	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}

	log.Println("redis服务连接成功，当前模式【集群】")
	return rdb
}
