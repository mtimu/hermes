package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mehditeymorian/hermes/internal/db/store"
	"github.com/mehditeymorian/hermes/internal/http/request"
	"github.com/mehditeymorian/hermes/internal/model"
	"go.uber.org/zap"
)

type Room struct {
	Logger *zap.Logger
	Store  store.Store
}

func (r Room) Register(app *fiber.App) {
	app.Get("/room/:id", r.get)
	app.Post("/room", r.create)
	app.Delete("/room/:id", r.del)
}

// get room's info.
func (r Room) get(ctx *fiber.Ctx) error {
	roomID := ctx.Params("id")

	r.Logger.Info("http.room.get", zap.String("roomID", roomID))

	room, err := r.Store.RoomCollection.Get(ctx.Context(), roomID)
	if err != nil {
		r.Logger.Error("failed to get room", zap.Error(err))

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"message": "failed to retrieve room",
			"error":   err.Error(),
		})
	}

	r.Logger.Info("http.room.get", zap.String("status", "ok"))

	return ctx.Status(http.StatusOK).JSON(room) //nolint:wrapcheck
}

// create a room.
func (r Room) create(ctx *fiber.Ctx) error {
	r.Logger.Info("http.room.create")

	var req request.NewRoom
	if err := ctx.BodyParser(&req); err != nil {
		r.Logger.Error("failed to parse request body", zap.Error(err))

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{ //nolint:wrapcheck
			"message": "failed to parse request body",
			"error":   err.Error(),
		})
	}

	room := model.NewRoom(req.HostID)

	err := r.Store.RoomCollection.Create(ctx.Context(), *room)
	if err != nil {
		r.Logger.Error("failed to create room", zap.Error(err))

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create room",
			"error":   err.Error(),
		})
	}

	r.Logger.Info("http.room.create", zap.String("status", "ok"))

	return ctx.Status(http.StatusOK).JSON(room) //nolint:wrapcheck
}

// del deletes a room.
func (r Room) del(ctx *fiber.Ctx) error {
	r.Logger.Info("http.room.delete")

	roomID := ctx.Params("id")

	err := r.Store.RoomCollection.Del(ctx.Context(), roomID)
	if err != nil {
		r.Logger.Error("failed to delete room", zap.Error(err))

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"message": "failed to delete room",
			"error":   err.Error(),
		})
	}

	r.Logger.Info("http.room.delete", zap.String("status", "ok"))

	return ctx.SendStatus(http.StatusNoContent) //nolint:wrapcheck
}
