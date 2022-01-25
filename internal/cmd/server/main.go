package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehditeymorian/hermes/internal/config"
	"github.com/mehditeymorian/hermes/internal/db/mongo"
	"github.com/mehditeymorian/hermes/internal/db/store"
	"github.com/mehditeymorian/hermes/internal/emq"
	"github.com/mehditeymorian/hermes/internal/http/handler"
	"github.com/mehditeymorian/hermes/internal/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func Command(cfgFile string) *cobra.Command {
	serverCommand := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "server",
		Short: "signaling server",
		Run:   run,
	}

	return serverCommand
}

func run(cmd *cobra.Command, _ []string) {
	cfgFile := cmd.Flag("config").Value.String()

	cfg := config.Load(cfgFile)

	emqClient := emq.Connect(cfg.Emq)

	emqClient.Publish("test", 0, false, "Hello World")

	dbClient, _ := mongo.Connect(cfg.DB)

	dbStore := store.New(dbClient)

	logger := log.New(cfg.Logger)

	app := fiber.New()

	handler.Room{
		Logger: logger,
		Store:  dbStore,
	}.Register(app)

	zap.L().Fatal("failed to run app", zap.Error(app.Listen(":3000")))
}
