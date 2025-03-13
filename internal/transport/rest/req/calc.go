package req

type CalcInsuline struct {
	ID          *string       `json:"id,omitempty"`              // Опционально
	Name        string        `json:"name" validate:"required"`  // Обязательное поле
	Email       string        `json:"email" validate:"required"` // Обязательное поле
	NewInsuline *noneInsuline `json:"insuline_type,omitempty"`   // Опционально, но если передано — все поля внутри обязательны
	Breakfast   *int          `json:"breakfast,omitempty"`       // Опционально
	Lunch       *int          `json:"lunch,omitempty"`           // Опционально
	Dinner      *int          `json:"dinner,omitempty"`          // Опционально
	Snack       *int          `json:"snack,omitempty"`           // Опционально
	SnackCount  *int          `json:"snack_count,omitempty"`     // Опционально
	Other       *int          `json:"other,omitempty"`           // Опционально
	OtherCount  *int          `json:"other_count,omitempty"`     // Опционально
	Days        int           `json:"days" validate:"required"`  // Обязательное поле
}

type noneInsuline struct {
	InsulineType string `json:"insuline_type" validate:"required"` // Ультракороткий, Короткий, Средний, Длительный, Комбинированный
	Medication   string `json:"medication" validate:"required"`    // Aspart, Glatgin, Levemir, Novorapid, Tresiba
	TradeName    string `json:"trade_name" validate:"required"`    // Ringlar, Rinfast, Rinlev, Rinliz, Rinliz
	UnitsMl      int    `json:"units_ml" validate:"required"`      // 100 (100ед/мл), 200 (200ед/мл)
	SizeMl       int    `json:"size_ml" validate:"required"`       // 3, 5, 10, 15, 20, 30 (мл)
}
