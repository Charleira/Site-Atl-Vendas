package main

import (
	"atlanta-site/config"
	"atlanta-site/controllers"
	_ "atlanta-site/docs"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Vendas
// @version 1.0
// @description API para gerenciamento de produtos, pedidos e usuários.
// @contact.name Suporte da API
// @contact.url http://www.seusite.com/support
// @contact.email suporte@seusite.com
// @host localhost:8080
// @BasePath /
func main() {
	config.ConnectToDatabase()

	router := gin.Default()
	// Configuração do CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // URL do seu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	controllers.SetupRoutes(router)

	// Rota para a documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
