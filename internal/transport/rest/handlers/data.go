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
