package user_controllers

import (
	"atlanta-site/models"
	services "atlanta-site/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserDetails obtém detalhes do usuário autenticado
func GetUserDetails(c *gin.Context) {
	userID := c.Param("id")

	user, err := services.GetUserDetailsService(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserDetails permite ao usuário atualizar seu perfil
func UpdateUserDetails(c *gin.Context) {
	userID := c.Param("id")
	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err := services.UpdateUserDetailsService(userID, updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Perfil atualizado com sucesso"})
}

// ChangePassword permite ao usuário alterar sua senha
func ChangePassword(c *gin.Context) {
	userID := c.Param("id")
	var passwordData struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&passwordData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err := services.ChangePasswordService(userID, passwordData.OldPassword, passwordData.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Senha alterada com sucesso"})
}

// DeleteUser permite ao usuário deletar sua conta
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := services.DeleteUserService(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}
