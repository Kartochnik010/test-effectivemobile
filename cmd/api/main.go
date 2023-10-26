package main

import (
	"context"

	"github.com/Kartochnik010/test-effectivemobile/internal/config"
	repo "github.com/Kartochnik010/test-effectivemobile/internal/repository"
	"github.com/Kartochnik010/test-effectivemobile/internal/repository/postgres"
	"github.com/Kartochnik010/test-effectivemobile/internal/service"
	transport "github.com/Kartochnik010/test-effectivemobile/internal/transport/http"
	"github.com/Kartochnik010/test-effectivemobile/pkg/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := postgres.InitDB(context.Background(), cfg.DSN)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			log.Err(err).Msg("failed to close db connection")
		}
	}()

	repos := repo.NewPostgresRepo(db.GetConnection())
	service := service.NewService(repos)
	t := transport.NewHandler(service, logger.NewLogger(false, ""))
	var srv transport.HttpServer
	if err := srv.Run(cfg.Port, t.Routes()); err != nil {
		log.Err(err).Msg("Failed to start server")
	}
}
