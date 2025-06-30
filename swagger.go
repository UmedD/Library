package main

import (
	_ "Library/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// <- сюда swag поместит сгенерированный docs/swagger.json
)

// setupSwagger настраивает endpoint /swagger/*any для просмотра UI
func setupSwagger(r *gin.Engine) {
	// URL, по которому будет доступна спецификация в JSON
	url := ginSwagger.URL("/swagger/doc.json")
	// регистрируем handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
