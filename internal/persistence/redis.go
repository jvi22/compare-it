package persistence

import (
    "context"
    "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(addr string) error {
    RedisClient = redis.NewClient(&redis.Options{Addr: addr})
    return RedisClient.Ping(context.Background()).Err()
}

// Session is marshalled to JSON before storing.
func SetSession(ctx context.Context, key string, value []byte) error {
    return RedisClient.Set(ctx, key, value, 0).Err()
}

func GetSession(ctx context.Context, key string) ([]byte, error) {
    return RedisClient.Get(ctx, key).Bytes()
}
