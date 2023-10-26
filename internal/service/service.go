package service

import (
	"github.com/Kartochnik010/test-effectivemobile/internal/models"
	"github.com/Kartochnik010/test-effectivemobile/internal/repository"
)

type Service struct {
	repository.Person
	requests
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Person:   NewPersonService(repo),
		requests: NewRequestsService(),
	}
}

type requests interface {
	AddData(person *models.Person)
}
