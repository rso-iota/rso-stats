package db

import (
	"context"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()
var client *redis.Client
var url string

func Init(redisURL string) *redis.Client {
	url = redisURL
	client = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.WithError(err).Error("Failed to connect to Redis")
	} else {
		log.Info("Connected to Redis at ", url)
	}

	return client
}

func DbError(msg string, err error) {
	log.WithError(err).Error(msg)
	log.Debug("Trying to reconnect to Redis")
	if client.Ping(ctx).Err() != nil {
		Init(url)
	}
}

func IncrementFood(playerId string) {
	err := client.Incr(ctx, playerId+":food").Err()
	if err != nil {
		DbError("Failed to increment food", err)
	}

	err = client.Incr(ctx, "total_food").Err()
	if err != nil {
		DbError("Failed to increment total_food", err)
	}
}

func IncrementDeaths(playerId string) {
	err := client.Incr(ctx, playerId+":deaths").Err()
	if err != nil {
		DbError("Failed to increment deaths", err)
	}
}

func IncrementKills(playerId string) {
	err := client.Incr(ctx, playerId+":kills").Err()
	if err != nil {
		DbError("Failed to increment kills", err)
	}

	err = client.Incr(ctx, "total_kills").Err()
	if err != nil {
		DbError("Failed to increment total_kills", err)
	}
}
