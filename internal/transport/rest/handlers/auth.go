package handlers

import (
	"backend/internal/models"
	"backend/internal/transport/rest/req"
	"backend/internal/transport/rest/res"
	"backend/internal/transport/service"
	"backend/internal/utils"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	auth     *service.AuthService
	validate *validator.Validate
}

func NewAuthHandler(auth *service.AuthService) *AuthHandler {
	return &AuthHandler{
		auth:     auth,
		validate: validator.New(),
	}
}

// SignUpUser регистрирует нового пользователя.
//
// @Summary      Регистрация пользователя
// @Description  Создаёт нового пользователя и выдаёт access-токен
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body models.User true "Данные пользователя"
// @Success      201 {object} res.SignUpRes
// @Failure      400 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /auth/sign-up [post]
func (h *AuthHandler) SignUpUser(c echo.Context) error {
	var client models.User
	if err := c.Bind(&client); err != nil {
		return c.JSON(utils.BadRequestError())
	}

	if err := h.validate.Struct(client); err != nil {
		return c.JSON(utils.BadRequestError())
	}

	token, err := h.auth.SignUpClient(&client, c)
	if err != nil {
		return c.JSON(utils.BadRequestError())
	}

	return c.JSON(http.StatusCreated, res.SignUpRes{
		Token:   token,
		Message: "User created successfully"})
}

// SignInUser аутентифицирует пользователя и возвращает access-токен.
//
// @Summary      Авторизация пользователя
// @Description  Проверяет телефон и пароль, затем выдаёт access-токен
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials body req.SignInReq true "Данные для входа"
// @Success      200 {object} res.SignInRes
// @Failure      400 {object} utils.ErrorResponse
// @Failure      401 {object} utils.ErrorResponse
// @Failure      500 {object} utils.ErrorResponse
// @Router       /auth/sign-in [post]
func (h *AuthHandler) SignInUser(c echo.Context) error {
	var req req.SignInReq
	if err := c.Bind(&req); err != nil {
		log.Println("Bind error:", err)
		return c.JSON(utils.BadRequestError())
	}

	if err := h.validate.Struct(req); err != nil {
		log.Println("Validation error:", err)
		return c.JSON(utils.BadRequestError())
	}
	client := models.User{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	token, err := h.auth.SignInClient(&client, c)
	if err != nil {
		log.Println("SignInClient error:", err)
		return c.JSON(utils.BadRequestError())
	}

	log.Println("User signed in successfully")

	return c.JSON(http.StatusOK, res.SignInRes{
		Token:   token,
		Message: "User signed in successfully",
	})
}
