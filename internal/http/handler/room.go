package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Room struct {
	Logger zap.Logger
}

func (r Room) Register(app *fiber.App) {
	app.Get("/room/:id", r.get)
	app.Post("/room", r.create)
	app.Delete("/room/:id", r.del)
}

// get room's info.
func (r Room) get(ctx *fiber.Ctx) error {
	panic(context.TODO())
}

// create a room.
func (r Room) create(ctx *fiber.Ctx) error {
	panic(context.TODO())
}

// del deletes a room.
func (r Room) del(ctx *fiber.Ctx) error {
	panic(context.TODO())
}
