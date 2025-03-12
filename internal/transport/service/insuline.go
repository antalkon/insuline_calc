package service

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"backend/internal/models"
	"backend/internal/repository"
)

type InsulineService struct {
	insulineRepo *repository.InsulineRepository
}

func NewInsulineService(repo *repository.InsulineRepository) *InsulineService {
	return &InsulineService{insulineRepo: repo}
}

func (s *InsulineService) CreateInsuline(i *models.Insuline, c echo.Context) (string, error) {
	i.ID = uuid.New()
	if err := s.insulineRepo.CreateInsuline(i); err != nil {
		return "", err
	}

	return i.ID.String(), nil
}

func (s *InsulineService) GetInsulineList(md string, c echo.Context) ([]models.Insuline, error) {
	ins, err := s.insulineRepo.GetInsulinesByMedication(&md)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
