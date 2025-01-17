package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"rso-stats/graph/model"
)

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
	totalPlayerFood := getOrZero(r.RedisClient, ctx, "total_food")
	totalPlayerKills := getOrZero(r.RedisClient, ctx, "total_kills")
	totalPlayerDeaths := getOrZero(r.RedisClient, ctx, "total_deaths")
	totalBotFood := getOrZero(r.RedisClient, ctx, "total_bot_food")
	totalBotKills := getOrZero(r.RedisClient, ctx, "total_bot_kills")
	totalBotDeaths := getOrZero(r.RedisClient, ctx, "total_bot_deaths")

	stats := model.GlobalStats{
		PlayerFoodEaten: int32(totalPlayerFood),
		PlayerKills:     int32(totalPlayerKills),
		PlayerDeaths:    int32(totalPlayerDeaths),
		BotFoodEaten:    int32(totalBotFood),
		BotKills:        int32(totalBotKills),
		BotDeaths:       int32(totalBotDeaths),
	}

	return &stats, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
