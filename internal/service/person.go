package service

import (
	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	repository "github.com/Kartochnik010/test-effectivemobile/internal/repository"
)

type PersonService struct {
	repo repository.Person // пока что эти методы никак не отличаются
}

func NewPersonService(repo repository.Person) *PersonService {
	return &PersonService{repo: repo}
}

func (p *PersonService) InsertPerson(person models.Person) (models.Person, error) {
	return models.Person{}, nil
}
func (p *PersonService) FindPersonById(id int) (models.Person, error) {
	return models.Person{}, nil
}
func (p *PersonService) DeletePersonById(id int) error {
	return nil
}
func (p *PersonService) UpdatePersonById(id int, person models.Person) (models.Person, error) {
	return models.Person{}, nil
}
func (p *PersonService) SearchPerson(filters models.Filters) ([]models.Person, models.Metadata, error) {
	return nil, models.Metadata{}, nil
}
