package handlers

import (
	"backend/internal/transport/service"
	"backend/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GameLinkHandler struct {
	service *service.GameLinkService
}

func NewGameLinkHandler(s *service.GameLinkService) *GameLinkHandler {
	return &GameLinkHandler{service: s}
}

type CreateGameLinkReq struct {
	GameName  string `json:"game_name" validate:"required"`
	SchoolNum string `json:"school_num" validate:"required"`
	Class     string `json:"class" validate:"required"`
	Comment   string `json:"comment"`
}
type CreateGameLinkRes struct {
	Code string `json:"code"`
}

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
