package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/TensoRaws/FinalRip/module/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClients(clients map[RDB]*Client) {
	for k := range clients {
		r := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", config.RedisConfig.Host, config.RedisConfig.Port),
			Password: fmt.Sprintf("%v", config.RedisConfig.Password),
			DB:       int(k) + 1,
			PoolSize: config.RedisConfig.PoolSize,
		})
		ctx := context.Background()
		clients[k] = &Client{C: r, Ctx: ctx}
	}
}

func NewRedisClient(n int) *Client {
	ctx := context.Background()
	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.RedisConfig.Host, config.RedisConfig.Port),
		Password: fmt.Sprintf("%v", config.RedisConfig.Password),
		DB:       n,
	})
	return &Client{C: r, Ctx: ctx}
}

// ClientGetName returns the name of the connection.
func (c Client) ClientGetName() *redis.StringCmd {
	return c.C.ClientGetName(c.Ctx)
}

func (c Client) Echo(message interface{}) *redis.StringCmd {
	return c.C.Echo(c.Ctx, message)
}

func (c Client) Ping() *redis.StatusCmd {
	return c.C.Ping(c.Ctx)
}

func (c Client) Del(keys ...string) *redis.IntCmd {
	return c.C.Del(c.Ctx, keys...)
}

func (c Client) Unlink(keys ...string) *redis.IntCmd {
	return c.C.Unlink(c.Ctx, keys...)
}

func (c Client) Dump(key string) *redis.StringCmd {
	return c.C.Dump(c.Ctx, key)
}

func (c Client) Exists(keys ...string) *redis.IntCmd {
	return c.C.Exists(c.Ctx, keys...)
}

func (c Client) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.Expire(c.Ctx, key, expiration)
}

func (c Client) ExpireNX(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.ExpireNX(c.Ctx, key, expiration)
}

func (c Client) ExpireXX(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.ExpireXX(c.Ctx, key, expiration)
}

func (c Client) ExpireGT(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.ExpireGT(c.Ctx, key, expiration)
}

func (c Client) ExpireLT(key string, expiration time.Duration) *redis.BoolCmd {
	return c.C.ExpireLT(c.Ctx, key, expiration)
}

func (c Client) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return c.C.ExpireAt(c.Ctx, key, tm)
}

func (c Client) ExpireTime(key string) *redis.DurationCmd {
	return c.C.ExpireTime(c.Ctx, key)
}

func (c Client) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	return c.C.ZRemRangeByScore(c.Ctx, key, min, max)
}

func (c Client) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return c.C.ZRange(c.Ctx, key, start, stop)
}

func (c Client) ZAddNX(key string, members ...redis.Z) *redis.IntCmd {
	return c.C.ZAddNX(c.Ctx, key, members...)
}

func (c Client) SMembers(key string) *redis.StringSliceCmd {
	return c.C.SMembers(c.Ctx, key)
}

func (c Client) SIsMember(key string, member interface{}) *redis.BoolCmd {
	return c.C.SIsMember(c.Ctx, key, member)
}

func (c Client) SAdd(key string, members ...interface{}) *redis.IntCmd {
	return c.C.SAdd(c.Ctx, key, members...)
}

func (c Client) SRem(key string, members ...interface{}) *redis.IntCmd {
	return c.C.SRem(c.Ctx, key, members...)
}

func (c Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.C.Set(c.Ctx, key, value, expiration)
}

func (c Client) Get(key string) *redis.StringCmd {
	return c.C.Get(c.Ctx, key)
}

func (c Client) HMSet(key string, fields map[string]interface{}) *redis.BoolCmd {
	return c.C.HMSet(c.Ctx, key, fields)
}

func (c Client) HMGet(key string, fields ...string) *redis.SliceCmd {
	return c.C.HMGet(c.Ctx, key, fields...)
}

func (c Client) HSet(key, field string, value interface{}) *redis.IntCmd {
	return c.C.HSet(c.Ctx, key, field, value)
}

func (c Client) HGet(key, field string) *redis.StringCmd {
	return c.C.HGet(c.Ctx, key, field)
}

func (c Client) HGetAll(key string) *redis.MapStringStringCmd {
	return c.C.HGetAll(c.Ctx, key)
}
