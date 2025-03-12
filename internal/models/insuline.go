package models

import "github.com/google/uuid"

type Insuline struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey"`
	InsulineType string    `json:"insuline_type" validate:"required"` // Ультракороткий, Короткий, Средний, Длительный, Комбинированный
	Medication   string    `json:"medication" validate:"required"`    // Aspart, Glatgin, Levemir, Novorapid, Tresiba
	TradeName    string    `json:"trade_name" validate:"required"`    // Ringlar, Rinfast, Rinlev, Rinliz, Rinliz
	UnitsMl      int       `json:"units_ml" validate:"required"`      // 100 (100ед/мл), 200 (200ед/мл)
	SizeMl       int       `json:"size_ml" validate:"required"`       // 3, 5, 10, 15, 20, 30 (мл)
	DisplayName  string    `json:"display_name" validate:"required"`  // РинГлар (Aspart), (100ед/мл), 3мл
	PhotoUrl     string    `json:"photo_url"`                         // https://example.com/photo.jpg
}
