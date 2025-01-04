package db

import (
	"context"
	"rso-stats/config"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()
var client *redis.Client

func Init(conf config.Config) {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.RedisURL,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to Redis")
	}

	log.Info("Connected to Redis at ", conf.RedisURL)
}

func IncrementFood(playerId string) {
	err := client.Incr(ctx, playerId+":food").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment food")
	}
}

func IncrementDeaths(playerId string) {
	err := client.Incr(ctx, playerId+":deaths").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment deaths")
	}
}

func IncrementKills(playerId string) {
	err := client.Incr(ctx, playerId+":kills").Err()
	if err != nil {
		log.WithError(err).Error("Failed to increment kills")
	}
}
