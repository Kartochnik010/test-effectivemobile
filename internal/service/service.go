package service

import (
	"github.com/Kartochnik010/test-effectivemobile/internal/repository"
)

type Service struct {
	repository.Person
	Requests
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Person:   NewPersonService(repo),
		Requests: NewRequestsService(),
	}
}

type Requests interface {
	GetURL(url string) string
}
