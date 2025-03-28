package handlers

import (
	"backend/internal/transport/service"
	"backend/internal/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DataHandler struct {
	data     *service.DataService
	validate *validator.Validate
}

func NewDataHandler(data *service.DataService) *DataHandler {
	return &DataHandler{
		data:     data,
		validate: validator.New(),
	}
}

func (h *DataHandler) GetUser(c echo.Context) error {
	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		code, resp := utils.UnauthorizedError()
		return c.JSON(code, resp)
	}

	user, err := h.data.GetUserByID(userID)
	if err != nil {
		if err.Error() == "record not found" {
			code, resp := utils.NotFoundError()
			return c.JSON(code, resp)
		}
		code, resp := utils.InternalServerError("failed to fetch user from DB")
		return c.JSON(code, resp)
	}

	return c.JSON(http.StatusOK, user)
}

// internal/transport/rest/handlers/data.go

func (h *DataHandler) GetGameStats(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return c.JSON(utils.BadRequestError())
	}

	users, err := h.data.GetGameStatsByCode(code)
	if err != nil {
		return c.JSON(utils.InternalServerError(err.Error()))
	}

	return c.JSON(http.StatusOK, users)
}

func (h *DataHandler) GenerateCertificate(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(utils.BadRequestError())
	}

	cert, err := h.data.GenerateCertificate(id)
	if err != nil {
		return c.JSON(utils.InternalServerError("Ошибка генерации PDF: " + err.Error()))
	}

	// Отдаем PDF файл
	return c.Blob(http.StatusOK, "application/pdf", cert)
}
