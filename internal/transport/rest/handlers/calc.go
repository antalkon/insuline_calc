package handlers

import (
	"backend/internal/transport/rest/req"
	"backend/internal/transport/service"
	"backend/internal/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type InsulineCalcHandler struct {
	calc     *service.CalcInsulineService
	validate *validator.Validate
}

func NewCalcInsulineHandler(calc *service.CalcInsulineService) *InsulineCalcHandler {
	return &InsulineCalcHandler{
		calc:     calc,
		validate: validator.New(),
	}
}

func (h *InsulineCalcHandler) CalcInsuline(c echo.Context) error {
	var insuline req.CalcInsuline
	if err := c.Bind(&insuline); err != nil {
		fmt.Println("Bind errir a")
		return c.JSON(utils.BadRequestError())
	}

	if err := h.validate.Struct(insuline); err != nil {
		fmt.Println("Validate errir a")

		return c.JSON(utils.BadRequestError())
	}

	id, err := h.calc.CalcInsuline(&insuline, c)
	if err != nil {
		fmt.Println("Db errir a")
		return c.JSON(utils.BadRequestError())
	}

	return c.JSON(http.StatusCreated, id)
}
