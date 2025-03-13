package service

import (
	"errors"
	"math"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"backend/internal/repository"
	"backend/internal/transport/rest/req"
	"backend/internal/transport/rest/res"
)

type CalcInsulineService struct {
	insulineRepo *repository.InsulineRepository
}

func NewCalcInsulineService(repo *repository.InsulineRepository) *CalcInsulineService {
	return &CalcInsulineService{insulineRepo: repo}
}

func (s *CalcInsulineService) CalcInsuline(i *req.CalcInsuline, c echo.Context) (res.CalcInsRes, error) {
	// Проверяем наличие обязательных данных
	if i.Days <= 0 {
		return res.CalcInsRes{}, errors.New("Days must be greater than zero")
	}

	// Безопасное извлечение значений (если nil, заменяем на 0)
	breakfast := 0
	if i.Breakfast != nil {
		breakfast = *i.Breakfast + 1
	}

	lunch := 0
	if i.Lunch != nil {
		lunch = *i.Lunch + 1
	}

	dinner := 0
	if i.Dinner != nil {
		dinner = *i.Dinner + 1
	}

	snack := 0
	if i.Snack != nil && i.SnackCount != nil {
		snack = (*i.Snack + 1) * *i.SnackCount
	}

	other := 0
	if i.Other != nil && i.OtherCount != nil {
		other = (*i.Other + 1) * *i.OtherCount
	}

	// Расчёт дозировки на день и на период
	DailyDosage := breakfast + lunch + dinner + snack + other
	PeriodDosage := DailyDosage * i.Days

	// Определение дозировки инсулиновой ручки
	var sizeMl, doseMl int

	if i.ID == nil || *i.ID == "" {
		// Если `ID` не передан, используем `NewInsuline`
		if i.NewInsuline == nil {
			return res.CalcInsRes{}, errors.New("Insuline details are required if ID is not provided")
		}
		if i.NewInsuline.SizeMl == 0 {
			return res.CalcInsRes{}, errors.New("SizeMl is required")
		}
		if i.NewInsuline.UnitsMl == 0 {
			return res.CalcInsRes{}, errors.New("UnitsMl is required")
		}

		sizeMl = i.NewInsuline.SizeMl
		doseMl = i.NewInsuline.UnitsMl
	} else {
		// Если `ID` передан, ищем инсулин в базе
		parsedUUID, err := uuid.Parse(*i.ID)
		if err != nil {
			return res.CalcInsRes{}, errors.New("Invalid UUID format")
		}

		insuline, err := s.insulineRepo.GetInsulineByID(parsedUUID)
		if err != nil {
			return res.CalcInsRes{}, err
		}

		sizeMl = insuline.SizeMl
		doseMl = insuline.UnitsMl
	}

	// Проверка на нулевые значения
	if sizeMl == 0 || doseMl == 0 {
		return res.CalcInsRes{}, errors.New("Invalid insulin data: sizeMl or doseMl cannot be zero")
	}

	// Расчёт дозировки в мл
	DailyDosageMl := float64(DailyDosage) / float64(doseMl)
	PeriodDosageMl := float64(PeriodDosage) / float64(doseMl)

	// Расчёт количества ручек
	DailyPens := DailyDosageMl / float64(sizeMl)
	PeriodPens := PeriodDosageMl / float64(sizeMl)
	PeriodPensRounded := int(math.Ceil(PeriodPens)) // Округление в большую сторону

	// Формируем и возвращаем результат
	result := res.CalcInsRes{
		DailyDosage:    DailyDosage,
		PeriodDosage:   PeriodDosage,
		DailuDosageMl:  DailyDosageMl,
		PeriodDosageMl: PeriodDosageMl,
		PensPeriod:     PeriodPensRounded,
		DailyPens:      DailyPens,
		Name:           i.Name,
		Email:          i.Email,
	}

	return result, nil
}
