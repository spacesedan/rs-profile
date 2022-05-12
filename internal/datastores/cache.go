package datastores

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

type Cache interface {
	Get(key string) *string
	Set(key string, value interface{}) error
	Delete(key string) error
}

type cache struct {
}

var Rdb *redis.Client

func NewCache(rdb *redis.Client) Cache {
	Rdb = rdb
	return &cache{}
}

func NewRedis() (*redis.Client, error) {
	ctx := context.Background()
	fmt.Println("REDIS HOST:", os.Getenv("REDIS_HOST"))
	opt, err := redis.ParseURL(os.Getenv("REDIS_HOST"))
	if err != nil {
		log.Printf("Could not parse redis host url: %v\n", err)
		return nil, err
	}

	rdb := redis.NewClient(opt)

	status := rdb.Ping(ctx)
	if status.Err() != nil {
		log.Printf("Could not ping redis: %v\n", status.Err())
		return nil, status.Err()
	}

	return rdb, nil
}

func (c *cache) Get(key string) *string {
	ctx := context.Background()
	val, err := Rdb.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
		return nil
	case err != nil:
		fmt.Println("Get failed", err)
		return nil
	case val == "":
		fmt.Println("value is empty")
		return nil
	default:
		return &val
	}
}

func (c *cache) Set(key string, value interface{}) error {
	ctx := context.Background()
	status := Rdb.Set(ctx, key, value, 0)
	if status.Err() != nil {
		log.Printf("Could not set value for KEY: %v\nERR: %v\n", key, status.Err())
		return status.Err()
	}

	return nil
}

func (c *cache) Delete(key string) error {
	ctx := context.Background()
	err := Rdb.Del(ctx, key).Err()
	if err != redis.Nil {
		log.Printf("Something went wrong when deleting value for: KEY: %v\nERR: %v\n", key, err)
		return err
	}

	return nil
}
