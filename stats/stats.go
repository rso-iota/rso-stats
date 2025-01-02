package stats

import (
	"rso-stats/config"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func FoodEaten(msg *nats.Msg) {
	log.WithField("food", string(msg.Data)).Info("Food eaten")
}

func Start(config config.Config) {
	conn, err := nats.Connect(config.NatsURL)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to nats")
	}
	defer conn.Close()

	log.Info("Connected to nats at ", config.NatsURL)

	conn.QueueSubscribe("food", "stats", FoodEaten)

	select {}
}
