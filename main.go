package main

import (
	"rso-stats/config"
	"rso-stats/db"
	"rso-stats/stats"
)

func main() {
	config := config.Init()

	db.Init(config)
	stats.Start(config)
}
