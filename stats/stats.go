package stats

import (
	"rso-stats/config"
	"rso-stats/db"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func FoodEaten(msg *nats.Msg) {
	db.IncrementFood(string(msg.Data))
}

func PlayerDied(msg *nats.Msg) {
	db.IncrementDeaths(string(msg.Data))
}

func PlayerKilled(msg *nats.Msg) {
	db.IncrementKills(string(msg.Data))
}

func Start(config config.Config) {
	conn, err := nats.Connect(config.NatsURL)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to nats")
	}
	defer conn.Close()

	log.Info("Connected to nats at ", config.NatsURL)

	conn.QueueSubscribe("food", "stats", FoodEaten)
	conn.QueueSubscribe("died", "stats", PlayerDied)
	conn.QueueSubscribe("kill", "stats", PlayerKilled)

	select {}
}
