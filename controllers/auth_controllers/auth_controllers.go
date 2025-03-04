package controllers

import (
	services "atlanta-site/services/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{Service: service}
}

// Login handler
func (ctrl *AuthController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	token, err := ctrl.Service.Login(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Logout handler (caso use blacklist de tokens)
func (ctrl *AuthController) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token não informado"})
		return
	}

	err := ctrl.Service.Logout(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deslogar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout realizado com sucesso"})
}

// Refresh Token handler
func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	var requestData struct {
		Token string `json:"token"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token inválido"})
		return
	}

	newToken, err := ctrl.Service.RefreshToken(requestData.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
