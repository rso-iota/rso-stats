package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"rso-stats/graph/model"

	"github.com/redis/go-redis/v9"
)

func playerData(ctx context.Context, redis *redis.Client, id string) *model.Player {
	foodEaten, err := redis.Get(ctx, id+":food").Int()
	if err != nil {
		foodEaten = 0
	}

	deaths, err := redis.Get(ctx, id+":deaths").Int()
	if err != nil {
		deaths = 0
	}

	kills, err := redis.Get(ctx, id+":kills").Int()
	if err != nil {
		kills = 0
	}

	player := model.Player{
		FoodEaten: int32(foodEaten),
		Deaths:    int32(deaths),
		Kills:     int32(kills),
	}

	return &player
}

// Player is the resolver for the player field.
func (r *queryResolver) Player(ctx context.Context, id string) (*model.Player, error) {
	return playerData(ctx, r.RedisClient, id), nil
}

// Players is the resolver for the players field.
func (r *queryResolver) Players(ctx context.Context, ids []*string) ([]*model.Player, error) {
	players := make([]*model.Player, len(ids))

	for i, id := range ids {
		players[i] = playerData(ctx, r.RedisClient, *id)
	}

	return players, nil
}

// Stats is the resolver for the stats field.
func (r *queryResolver) Stats(ctx context.Context) (*model.GlobalStats, error) {
	totalFood, err := r.RedisClient.Get(ctx, "total_food").Int()
	if err != nil {
		return nil, err
	}

	totalKills, err := r.RedisClient.Get(ctx, "total_kills").Int()
	if err != nil {
		return nil, err
	}

	stats := model.GlobalStats{
		FoodEaten: int32(totalFood),
		Kills:     int32(totalKills),
	}

	return &stats, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
