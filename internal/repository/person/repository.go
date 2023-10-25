package personrepository

import (
	"context"

	"github.com/Kartochnik010/test-effectivemobile/db/postgres"
	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetPerson(id string) (models.Person, error)
	InsertPerson(user models.Person) error
	DeletePersonById(id string) error
	FindPersonById(id string) (models.Person, error)

	GetPersonWithContext(ctx context.Context, id string) (models.Person, error)
	InsertPersonWithContext(ctx context.Context, user models.Person) error
	DeletePersonByIdWithContext(ctx context.Context, id string) error
	FindPersonByIdWithContext(ctx context.Context, id string) (models.Person, error)
}

const PersonTable = "person"

type repositoryDB struct {
	dbClient *pgx.Conn
}

func New(postgresRepository postgres.Repository) Repository {
	return &repositoryDB{
		dbClient: postgresRepository.GetConnection(),
	}
}

func (r *repositoryDB) GetPerson(id string) (models.Person, error) {
	return models.Person{}, nil
}
func (r *repositoryDB) InsertPerson(user models.Person) error {

	return nil
}
func (r *repositoryDB) DeletePersonById(id string) error {
	return nil
}
func (r *repositoryDB) FindPersonById(id string) (models.Person, error) {
	return models.Person{}, nil
}

func (r *repositoryDB) GetPersonWithContext(ctx context.Context, id string) (models.Person, error) {
	return models.Person{}, nil
}
func (r *repositoryDB) InsertPersonWithContext(ctx context.Context, user models.Person) error {
	return nil
}
func (r *repositoryDB) DeletePersonByIdWithContext(ctx context.Context, id string) error { return nil }
func (r *repositoryDB) FindPersonByIdWithContext(ctx context.Context, id string) (models.Person, error) {
	return models.Person{}, nil
}
