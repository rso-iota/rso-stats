package db

import (
	"context"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()
var client *redis.Client

func Init(redisURL string) *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.WithError(err).Error("Failed to connect to Redis")
	} else {
		log.Info("Connected to Redis at ", redisURL)
	}

	return client
}

func IncrementFood(playerId string) {
	err := client.Incr(ctx, playerId+":food").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment food")
	}

	err = client.Incr(ctx, "total_food").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment total_food")
	}
}

func IncrementDeaths(playerId string) {
	err := client.Incr(ctx, playerId+":deaths").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment deaths")
	}

	err = client.Incr(ctx, "total_deaths").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment total_deaths")
	}
}

func IncrementKills(playerId string) {
	err := client.Incr(ctx, playerId+":kills").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment kills")
	}

	err = client.Incr(ctx, "total_kills").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment total_kills")
	}
}

func IncrementBotKills() {
	err := client.Incr(ctx, "total_bot_kills").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment total_bot_kills")
	}
}

func IncrementBotDeaths() {
	err := client.Incr(ctx, "total_bot_deaths").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment total_bot_deaths")
	}
}

func IncrementBotFood() {
	err := client.Incr(ctx, "total_bot_food").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment total_bot_food")
	}
}
