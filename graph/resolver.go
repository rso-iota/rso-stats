package graph

import (
	"context"
	"rso-stats/graph/model"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RedisClient *redis.Client
}

func getOrZero(client *redis.Client, ctx context.Context, key string) int {
	val, err := client.Get(ctx, key).Int()
	if err != nil {
		log.WithError(err).Errorf("Failed to get %s", key)
		return 0
	}

	return val
}

func playerData(ctx context.Context, client *redis.Client, id string) *model.Player {
	foodEaten := getOrZero(client, ctx, id+":food")
	deaths := getOrZero(client, ctx, id+":deaths")
	kills := getOrZero(client, ctx, id+":kills")

	player := model.Player{
		FoodEaten: int32(foodEaten),
		Deaths:    int32(deaths),
		Kills:     int32(kills),
	}

	return &player
}
