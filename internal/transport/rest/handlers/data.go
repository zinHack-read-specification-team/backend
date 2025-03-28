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

// GetUser возвращает информацию о текущем пользователе.
//
// @Summary      Получить текущего пользователя
// @Description  Возвращает профиль пользователя по JWT-токену (userID из контекста).
// @Tags         User
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} models.User
// @Failure      401 {object} utils.ErrorResponse
// @Failure      404 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /user [get]
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

// GetGameStats возвращает список пользователей и статистику по коду игры.
//
// @Summary      Получить статистику по игре
// @Description  Возвращает список участников и их результаты по коду игры.
// @Tags         Game
// @Produce      json
// @Param        code path string true "Код игры"
// @Success      200 {array} models.GameUser
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /game/{code}/stats [get]
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

// GenerateCertificate генерирует PDF-сертификат для пользователя.
//
// @Summary      Генерация сертификата
// @Description  Генерирует и возвращает PDF-сертификат участника по его ID.
// @Tags         Certificate
// @Produce      application/pdf
// @Param        id path string true "ID пользователя"
// @Success      200 {file} file
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /certificate/{id} [get]
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

	return c.Blob(http.StatusOK, "application/pdf", cert)
}
