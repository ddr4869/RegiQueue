package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1234",
	})
}

func IncrementQueue(ctx context.Context, userId string) (int64, error) {
	queuePosition, err := rdb.LPush(ctx, "registration_queue", userId).Result()
	if err != nil {
		log.Printf("Failed to add user %s to queue: %v", userId, err)
		return 0, err
	}
	return queuePosition, nil
}

func DecrementQueue(ctx context.Context, userId string) {
	_, err := rdb.LRem(ctx, "registration_queue", 0, userId).Result()
	if err != nil {
		log.Printf("Failed to remove user %s from queue: %v", userId, err)
	}
}

func GetQueuePosition(ctx context.Context, userId string) (int64, error) {
	queueLength, err := rdb.LLen(ctx, "registration_queue").Result()
	if err != nil {
		log.Printf("Failed to get queue length: %v", err)
		return 0, err
	}

	position := queueLength - int64(rdb.LPos(ctx, "registration_queue", userId, redis.LPosArgs{}).Val()) - 1
	return position, nil
}
