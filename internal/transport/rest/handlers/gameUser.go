// handlers/game_user.go
package handlers

import (
	"backend/internal/transport/service"
	"backend/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GameUserHandler struct {
	service *service.GameUserService
}

func NewGameUserHandler(s *service.GameUserService) *GameUserHandler {
	return &GameUserHandler{service: s}
}

type RegisterPlayerReq struct {
	Code     string `json:"code"`
	FullName string `json:"full_name"`
}

type FinishGameReq struct {
	ID    uuid.UUID `json:"id"`
	Stars int       `json:"stars"`
	Score int       `json:"score"`
}

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
