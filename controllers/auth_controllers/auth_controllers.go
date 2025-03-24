package auth_controller

import (
	"atlanta-site/repositories"
	"atlanta-site/services/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController gerencia autenticação de usuários
type AuthController struct {
	// O serviço agora será um campo simples
}

// NewAuthController cria uma nova instância de AuthController
func NewAuthController() *AuthController {
	return &AuthController{}
}

// Structs para definir os dados de entrada
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	Token string `json:"token"`
}

// Login autentica um usuário e retorna um token JWT
func Login(c *gin.Context) {
	var loginData LoginRequest

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Criar uma instância do repositório AuthRepository
	repo := repositories.NewAuthRepository()

	// Passa o repositório para a função de login
	token, err := auth_service.Login(loginData.Email, loginData.Password, repo)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Logout encerra a sessão do usuário removendo o token JWT
func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token não informado"})
		return
	}

	err := auth_service.Logout(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deslogar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout realizado com sucesso"})
}

// RedirectToDiscord redireciona o usuário para o OAuth2 do Discord
func RedirectToDiscord(c *gin.Context) {
	clientID := os.Getenv("DISCORD_CLIENT_ID")
	redirectURI := os.Getenv("DISCORD_REDIRECT_URI") // Exemplo: "https://seusite.com/auth/callback"

	if clientID == "" || redirectURI == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Configuração do Discord não encontrada"})
		return
	}

	discordAuthURL := fmt.Sprintf(
		"https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=identify email",
		clientID,
		redirectURI,
	)

	c.Redirect(http.StatusFound, discordAuthURL)
}

// RefreshToken gera um novo token JWT baseado em um token antigo
func RefreshToken(c *gin.Context) {
	var requestData RefreshTokenRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token inválido"})
		return
	}

	newToken, err := auth_service.RefreshToken(requestData.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
