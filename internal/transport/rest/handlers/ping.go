package handlers

import "github.com/labstack/echo/v4"

// Ping godoc
//
// @Summary      Ping endpoint
// @Description  Проверка, что сервер работает (health-check)
// @Tags         health
// @Produce      json
// @Success      200 {object} PingResponse
// @Router       /ping [get]
func Ping(c echo.Context) error {
	return c.JSON(200, PingResponse{Message: "pong"})
}

// PingResponse структура ответа для ping.
//
// @Description Ответ на пинг-запрос.
// @name PingResponse
type PingResponse struct {
	Message string `json:"message" example:"pong"`
}
