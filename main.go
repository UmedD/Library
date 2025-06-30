// @title         OnlineLibrary API
// @version       1.0
// @description   RESTful API сервиса OnlineLibrary
// @host          localhost:9091
// @BasePath      /

package main

import (
	"Library/internal/config"
	"Library/internal/controller"
	"Library/internal/db"
	"Library/logger"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 1) Читаем .env и JSON-конфиг (логгер, JWT, app-параметры)
	if err := config.ReadSettings(); err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	// 2) Инициализируем ротационные логгеры (lumberjack + gin.DefaultWriter)
	if err := logger.Init(); err != nil {
		log.Fatalf("cannot init logger: %v", err)
	}

	// 3) Подключаемся к базе (PostgresParams берутся из config.AppSettings)
	if err := db.ConnectDB(config.AppSettings.PostgresParams); err != nil {
		logger.Error.Fatalf("DB connection failed: %v", err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			logger.Error.Printf("DB close error: %v", err)
		}
	}()

	// 4) Выбираем режим Gin (release/debug)
	gin.SetMode(config.AppSettings.AppParams.GinMode)

	// 5) Инициализируем роутер
	r := gin.Default()

	setupSwagger(r)
	// 6) Регистрируем публичные и защищённые маршруты
	controller.RegisterAuthRoutes(r)   // /auth/sign-up, /auth/sign-in
	controller.RegisterUserRoutes(r)   // /users (GET открытые, POST/PUT/DELETE через JWT+AdminOnly)
	controller.RegisterAuthorRoutes(r) // /authors
	controller.RegisterBookRoutes(r)   // /books

	// 7) Старт сервера на порту из конфига
	addr := config.AppSettings.AppParams.PortRun
	logger.Info.Printf("Starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		logger.Error.Fatalf("Server stopped with error: %v", err)
	}
}
