package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitCache() *redis.Client {

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal("failed to parse redis url: ", err)
	}

	rdb := redis.NewClient(opt)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("failed to connect redis: ", err)
	}

	log.Println("Connected to Redis successfully")
	return rdb

}
