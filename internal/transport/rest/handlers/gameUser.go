package handlers

import (
	"backend/internal/transport/service"
	"backend/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GameUserHandler обрабатывает игровые действия участника.
type GameUserHandler struct {
	service *service.GameUserService
}

// NewGameUserHandler создаёт экземпляр GameUserHandler.
func NewGameUserHandler(s *service.GameUserService) *GameUserHandler {
	return &GameUserHandler{service: s}
}

// RegisterPlayerReq содержит данные для регистрации игрока.
//
// @Description Данные для регистрации игрока по коду игры.
// @name RegisterPlayerReq
type RegisterPlayerReq struct {
	Code     string `json:"code" example:"AB12CD"`
	FullName string `json:"full_name" example:"Иван Иванов"`
}

// FinishGameReq содержит данные об окончании игры.
//
// @Description Данные для завершения игры и отправки результатов.
// @name FinishGameReq
type FinishGameReq struct {
	ID    uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Stars int       `json:"stars" example:"3"`
	Score int       `json:"score" example:"1200"`
}

// RegisterPlayer регистрирует игрока по коду.
//
// @Summary      Регистрация игрока
// @Description  Регистрирует участника игры по коду и имени.
// @Tags         GameUser
// @Accept       json
// @Produce      json
// @Param        request body RegisterPlayerReq true "Данные игрока"
// @Success      201 {object} models.GameUser
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /game-user/register [post]
func (h *GameUserHandler) RegisterPlayer(c echo.Context) error {
	var req RegisterPlayerReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.BadRequestError())
	}

	player, err := h.service.RegisterPlayer(req.Code, req.FullName)
	if err != nil {
		return c.JSON(utils.InternalServerError(err.Error()))
	}

	return c.JSON(http.StatusCreated, player)
}

// FinishGame сохраняет результаты игры.
//
// @Summary      Завершить игру
// @Description  Отправляет звезды и очки игрока после окончания игры.
// @Tags         GameUser
// @Accept       json
// @Produce      json
// @Param        id path string true "ID участника"
// @Param        request body FinishGameReq true "Результаты игры"
// @Success      200 {object} map[string]string
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /game-user/{id}/finish [post]
func (h *GameUserHandler) FinishGame(c echo.Context) error {
	idParam := c.Param("id")
	gameUserID, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(utils.BadRequestError())
	}

	var req struct {
		Stars int `json:"stars"`
		Score int `json:"score"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.BadRequestError())
	}

	if err := h.service.FinishGame(gameUserID, req.Stars, req.Score); err != nil {
		return c.JSON(utils.InternalServerError(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Game results updated"})
}
