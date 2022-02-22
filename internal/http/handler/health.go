package handler

import (
	"context"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Health struct {
	DB  *mongo.Database
	Emq mqtt.Client
}

func (h Health) Register(app *fiber.App) {
	app.Get("/healthz", h.health)
}

func (h Health) health(ctx *fiber.Ctx) error {

	timeoutCtx, done := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer done()

	err := h.DB.Client().Ping(timeoutCtx, readpref.Primary())
	if err != nil {

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":       err.Error(),
			"description": "failed to ping database",
		})
	}

	if !h.Emq.IsConnectionOpen() {

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":       "there is no active connection to mqtt broker",
			"description": "the client is disconnected or reconnecting",
		})
	}

	return ctx.SendStatus(http.StatusNoContent)
}
