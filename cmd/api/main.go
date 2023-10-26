package main

import (
	"context"

	"github.com/Kartochnik010/test-effectivemobile/internal/config"
	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	repo "github.com/Kartochnik010/test-effectivemobile/internal/repository"
	"github.com/Kartochnik010/test-effectivemobile/internal/repository/postgres"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

var ()

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
			log.Error().Msg(err.Error())
		}
	}()

	repos := repo.NewPostgresRepo(db.GetConnection())

	p := models.Person{
		Name:    "D",
		Surname: "U",
	}

	newP, err := repos.InsertPerson(p)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	if _, err := repos.UpdatePersonById(newP.ID, models.Person{Surname: "B", Patronymic: "A"}); err != nil {
		log.Err(err)
	}
}
