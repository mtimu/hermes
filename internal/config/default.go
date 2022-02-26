package config

import (
	"github.com/mehditeymorian/hermes/internal/db"
	"github.com/mehditeymorian/hermes/internal/emq"
	"github.com/mehditeymorian/hermes/internal/log"
)

func Default() Config {
	return Config{
		Emq: emq.Config{
			ClientID: "test-client",
			URL:      "tcp://emq:1883",
		},
		DB: db.Config{
			Name: "Hermes",
			URI:  "mongodb://mongo:27017",
		},
		Logger: log.Config{
			Production: false,
			Encoding:   "console",
			Level:      "info",
		},
	}
}
