package main

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/Kartochnik010/test-effectivemobile/db/postgres"
	"github.com/Kartochnik010/test-effectivemobile/internal/config"
	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	repo "github.com/Kartochnik010/test-effectivemobile/internal/repository"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// ctx, _ := context.WithDeadline(context.Background(), time.Date(0, 0, 0, 0, 0, 5, 0, time.UTC))
	db, err := postgres.New(context.Background(), cfg.DSN)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			slog.Error(err.Error())
		}
	}()

	repos := repo.New(db)

	p := models.Person{
		Name:    "Dmitriy",
		Surname: "Ushakov",
	}

	if err := repos.PersonRepo.InsertPerson(p); err != nil {
		slog.Error(err.Error())
	}
	fmt.Println("as")
}
