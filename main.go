package main

import (
	"rso-stats/config"
	"rso-stats/stats"
)

func main() {
	config := config.Init()
	stats.Start(config)
}
