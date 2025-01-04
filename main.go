package main

import (
	"rso-stats/config"
	"rso-stats/db"
	"rso-stats/graph"
	"rso-stats/stats"
)

func main() {
	config := config.Init()

	redis := db.Init(config)
	go graph.Init(config, redis)

	stats.Start(config)
}
