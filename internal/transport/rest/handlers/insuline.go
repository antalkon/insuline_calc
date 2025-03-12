package handlers

import (
	"backend/internal/models"
	"backend/internal/transport/rest/res"
	"backend/internal/transport/service"
	"backend/internal/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type InsulineHandler struct {
	insuline *service.InsulineService
	validate *validator.Validate
}

func NewInsulineHandler(insuline *service.InsulineService) *InsulineHandler {
	return &InsulineHandler{
		insuline: insuline,
		validate: validator.New(),
	}
}

func (h *InsulineHandler) AddNewInsuline(c echo.Context) error {
	var insuline models.Insuline
	if err := c.Bind(&insuline); err != nil {
		fmt.Println("Bind errir a")

		return c.JSON(utils.BadRequestError())
	}

	if err := h.validate.Struct(insuline); err != nil {
		fmt.Println("Validate errir a")

		return c.JSON(utils.BadRequestError())
	}

	id, err := h.insuline.CreateInsuline(&insuline, c)
	if err != nil {
		fmt.Println("Db errir a")
		return c.JSON(utils.BadRequestError())
	}

	return c.JSON(http.StatusCreated, res.SignUpRes{
		Token:   id,
		Message: "Insuline added successfully"})
}

func (h *InsulineHandler) GetInsuline(c echo.Context) error {
	medication := c.QueryParam("medication")
	ins, err := h.insuline.GetInsulineList(medication, c)
	if err != nil {
		return c.JSON(utils.BadRequestError())
	}

	return c.JSON(http.StatusOK, ins)
}
