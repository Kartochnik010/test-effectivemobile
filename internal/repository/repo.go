package repos

import (
	"github.com/Kartochnik010/test-effectivemobile/db/postgres"
	personRepository "github.com/Kartochnik010/test-effectivemobile/internal/repository/person"
)

type Repos struct {
	PersonRepo personRepository.Repository
}

func New(postgresRepository postgres.Repository) *Repos {
	return &Repos{
		PersonRepo: personRepository.New(postgresRepository),
	}
}
