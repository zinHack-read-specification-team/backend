package router

import (
	"backend/internal/repository"
	rmiddleware "backend/internal/transport/rest/Rmiddleware"
	"backend/internal/transport/rest/handlers"
	"backend/internal/transport/service"
	"backend/pkg/cache"
	"backend/pkg/config"
	"backend/pkg/db"
	"backend/pkg/logger"
	"backend/pkg/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(e *echo.Echo, cfg *config.Config, log *logger.Logger, db *db.Database, cache *cache.RedisClient, s3 *storage.MinIOClient) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // разрешаем все источники
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	ddbb := db.DB

	authRepo := repository.NewAuthRepository(ddbb)
	dataRepo := repository.NewDataRepository(ddbb)
	GameLinkRepo := repository.NewGameLinkRepository(ddbb)

	authService := service.NewAuthService(authRepo)
	dataService := service.NewDataService(dataRepo)
	GgameLinkService := service.NewGameLinkService(GameLinkRepo)

	authHandler := handlers.NewAuthHandler(authService)
	dataHandler := handlers.NewDataHandler(dataService)
	gameLinkHandler := handlers.NewGameLinkHandler(GgameLinkService)

	api := e.Group("/api/v1")
	api.GET("/ping", handlers.Ping)

	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", authHandler.SignUpUser)
		auth.POST("/sign-in", authHandler.SignInUser)
	}

	data := api.Group("/data", rmiddleware.JWTMiddleware)
	{
		data.GET("/get-user", dataHandler.GetUser)
		data.POST("/create-link", gameLinkHandler.CreateGameLink)
		data.GET("/links", gameLinkHandler.GetUserLinks)
		data.GET("/stats/:code", dataHandler.GetGameStats)

	}
	user := api.Group("/user")
	{
		user.GET("/check-link/:code", gameLinkHandler.CheckLink)

	}

	gameUserRepo := repository.NewGameUserRepository(ddbb)
	gameUserService := service.NewGameUserService(gameUserRepo, GameLinkRepo)
	gameUserHandler := handlers.NewGameUserHandler(gameUserService)

	game := api.Group("/game")
	{
		game.POST("/register", gameUserHandler.RegisterPlayer)
		game.PUT("/finish/:id", gameUserHandler.FinishGame) // <-- исправлено
		game.GET("/certificate/:id", dataHandler.GenerateCertificate)

	}
}
