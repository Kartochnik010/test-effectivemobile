package transport

import (
	"github.com/Kartochnik010/test-effectivemobile/internal/service"
	"github.com/Kartochnik010/test-effectivemobile/pkg/logger"
)

type Handler struct {
	service *service.Service
	logger  *logger.Logger
}

func NewHandler(service *service.Service, logger *logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
