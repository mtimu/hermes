package config

import (
	"github.com/mehditeymorian/hermes/internal/db"
	"github.com/mehditeymorian/hermes/internal/emq"
)

func Default() Config {
	return Config{
		Emq: emq.Config{
			ClientID: "test-client",
			URL:      "tcp://localhost:18083",
		},
		DB: db.Config{
			Name: "Hermes",
			URI:  "http://localhost:27017",
		},
	}
}
