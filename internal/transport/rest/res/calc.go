package res

type CalcInsRes struct {
	DailyDosage  int `json:"daily_dosage"`
	PeriodDosage int `json:"period_dosage"`

	DailuDosageMl  float64 `json:"daily_dosage_ml"`
	PeriodDosageMl float64 `json:"period_dosage_ml"`

	DailyPens  float64 `json:"daily_pens"`
	PensPeriod int     `json:"pens_period"`

	Name  string `json:"name"`
	Email string `json:"email"`

	Period int `json:"period"`
}
