// controllers/user_controllers.go
package user_controllers

import (
	"atlanta-site/models"
	services "atlanta-site/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserDetails obtém detalhes do usuário autenticado
// @Summary Obtém os detalhes de um usuário
// @Description Retorna informações do usuário pelo ID
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string "Usuário não encontrado"
// @Router /users/{id} [get]
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
// @Summary Atualiza os detalhes de um usuário
// @Description Permite que o usuário atualize seu perfil com novas informações
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Param user body models.User true "Dados do usuário para atualização"
// @Success 200 {object} map[string]string "Perfil atualizado com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 500 {object} map[string]string "Erro ao atualizar perfil"
// @Router /users/{id} [put]
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
// @Summary Altera a senha de um usuário
// @Description Permite que um usuário altere sua senha fornecendo a senha antiga e a nova senha
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Param passwordData body object true "Dados para alteração de senha"
// @Success 200 {object} map[string]string "Senha alterada com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 500 {object} map[string]string "Erro ao alterar senha"
// @Router /users/{id}/password [put]
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
// @Summary Deleta um usuário
// @Description Permite que um usuário exclua sua conta
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário"
// @Success 200 {object} map[string]string "Usuário deletado com sucesso"
// @Failure 500 {object} map[string]string "Erro ao deletar usuário"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := services.DeleteUserService(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}
