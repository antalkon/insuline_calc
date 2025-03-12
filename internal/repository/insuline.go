package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/internal/models"
)

type InsulineRepository struct {
	db *gorm.DB
}

func NewInsulineRepository(db *gorm.DB) *InsulineRepository {
	if db == nil {
		panic("Database connection is nil in repository")
	}
	return &InsulineRepository{db: db}
}

func (r *InsulineRepository) CreateInsuline(insuline *models.Insuline) error {
	return r.db.Table("insulines").Create(insuline).Error
}
func (r *InsulineRepository) GetInsulinesByMedication(medication *string) ([]models.Insuline, error) {
	var insulines []models.Insuline
	query := r.db.Table("insulines")

	if medication != nil && *medication != "" {
		query = query.Where("medication = ?", *medication)
	}

	err := query.Find(&insulines).Error
	if err != nil {
		return nil, err
	}

	return insulines, nil
}
func (r *InsulineRepository) GetInsulineByID(id uuid.UUID) (*models.Insuline, error) {
	var insuline models.Insuline
	err := r.db.Table("insulines").Where("id = ?", id).First(&insuline).Error
	if err != nil {
		return nil, err
	}
	return &insuline, nil
}
