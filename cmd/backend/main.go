package main

import (
	"backend/internal/app"
	"log"

	_ "backend/docs" // подключение сгенерированных Swagger-документов
)

// @title           Безопасная школа API
// @version         1.0
// @description     Backend-сервис для работы с безопасной школьной платформой. Используйте спецификацию ниже для интеграции.

// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html

// @host           zin-hack-25.antalkon.ru
// @BasePath       /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI спецификация
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Application initialization failed: %v", err)
	}
	application.Run()

	if err := application.RunServer(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
