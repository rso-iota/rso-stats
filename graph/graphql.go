package graph

import (
	"net/http"
	"rso-stats/config"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func Init(conf config.Config, redisClient *redis.Client) {
	srv := handler.New(NewExecutableSchema(Config{Resolvers: &Resolver{RedisClient: redisClient}}))

	srv.AddTransport(&transport.POST{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Info("GraphQL server running on port ", conf.GrapQLPort)
	log.Fatal(http.ListenAndServe(":"+conf.GrapQLPort, nil))
}
