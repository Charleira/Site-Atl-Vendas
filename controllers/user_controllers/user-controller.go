// user_controllers/user_controller.go
package user_controllers

import (
	"atlanta-site/models"
	services "atlanta-site/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser lida com a criação de um novo usuário
func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Chama o service para fazer as validações e inserir o usuário
	if err := services.RegisterUserService(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário registrado com sucesso!"})
}

// PromoteUserToAdmin lida com a promoção de um usuário a administrador
func PromoteUserToAdmin(c *gin.Context) {
	id := c.Param("id")

	// Chama o service para promover o usuário
	if err := services.PromoteUserToAdminService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário promovido a administrador!"})
}
