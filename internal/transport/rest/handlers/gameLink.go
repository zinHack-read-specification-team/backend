package handlers

import (
	"backend/internal/transport/service"
	"backend/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GameLinkHandler обрабатывает маршруты, связанные с игровыми ссылками.
type GameLinkHandler struct {
	service *service.GameLinkService
}

// NewGameLinkHandler создает новый экземпляр GameLinkHandler.
func NewGameLinkHandler(s *service.GameLinkService) *GameLinkHandler {
	return &GameLinkHandler{service: s}
}

// CreateGameLinkReq содержит параметры для создания ссылки на игру.
//
// @Description Запрос на создание ссылки на игру.
// @name CreateGameLinkReq
type CreateGameLinkReq struct {
	GameName  string `json:"game_name" validate:"required" example:"Историческая викторина"`
	SchoolNum string `json:"school_num" validate:"required" example:"42"`
	Class     string `json:"class" validate:"required" example:"7А"`
	Comment   string `json:"comment" example:"Для параллели 7-х классов"`
}

// CreateGameLinkRes возвращает код созданной ссылки.
//
// @Description Ответ после успешного создания ссылки.
// @name CreateGameLinkRes
type CreateGameLinkRes struct {
	Code string `json:"code" example:"AB12CD"`
}

// CreateGameLink создает новую ссылку на игру.
//
// @Summary      Создание ссылки на игру
// @Description  Генерирует уникальный код и сохраняет ссылку с параметрами
// @Tags         GameLink
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request body CreateGameLinkReq true "Параметры для создания ссылки"
// @Success      201 {object} CreateGameLinkRes
// @Failure      400 {object} utils.ErrorResponse
// @Failure      401 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /game-link [post]
func (h *GameLinkHandler) CreateGameLink(c echo.Context) error {
	var req CreateGameLinkReq
	if err := c.Bind(&req); err != nil || req.GameName == "" {
		return c.JSON(utils.BadRequestError())
	}

	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return c.JSON(utils.UnauthorizedError())
	}

	link, err := h.service.CreateLink(userID, req.GameName, req.SchoolNum, req.Class, req.Comment)
	if err != nil {
		return c.JSON(utils.InternalServerError(err.Error()))
	}

	return c.JSON(http.StatusCreated, CreateGameLinkRes{
		Code: link.Code,
	})
}

// GetUserLinks возвращает все ссылки, созданные пользователем.
//
// @Summary      Получить свои ссылки
// @Description  Возвращает список всех игровых ссылок, созданных текущим пользователем.
// @Tags         GameLink
// @Security     BearerAuth
// @Produce      json
// @Success      200 {array} models.GameLink
// @Failure      401 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /game-link [get]
func (h *GameLinkHandler) GetUserLinks(c echo.Context) error {
	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return c.JSON(utils.UnauthorizedError())
	}

	links, err := h.service.GetLinksByUserID(userID)
	if err != nil {
		return c.JSON(utils.InternalServerError(err.Error()))
	}

	return c.JSON(http.StatusOK, links)
}

// CheckLinkReq — структура запроса для проверки ссылки.
//
// @Description Проверка ссылки по коду
// @name CheckLinkReq
type CheckLinkReq struct {
	Code string `json:"code" validate:"required,len=6" example:"AB12CD"`
}

// CheckLink проверяет существование ссылки по её коду.
//
// @Summary      Проверить ссылку по коду
// @Description  Проверяет, существует ли ссылка с данным кодом.
// @Tags         GameLink
// @Produce      json
// @Param        code path string true "Код ссылки (6 символов)"
// @Success      200 {object} models.GameLink
// @Failure      400 {object} utils.ErrorResponse
// @Failure      404 {object} utils.ErrorResponse
// @Router       /game-link/{code} [get]
func (h *GameLinkHandler) CheckLink(c echo.Context) error {
	code := c.Param("code")
	if len(code) != 6 {
		return c.JSON(utils.BadRequestError())
	}

	link, err := h.service.CheckLink(code)
	if err != nil {
		return c.JSON(utils.NotFoundError())
	}

	return c.JSON(http.StatusOK, link)
}
