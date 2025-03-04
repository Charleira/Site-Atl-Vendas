package auth_service

import (
	"atlanta-site/repositories"
	"atlanta-site/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("SUA_CHAVE_SECRETA")

// Login do usuário
// Verifica as credenciais do usuário e gera um token JWT
func Login(email, password string, repo *repositories.AuthRepository) (string, error) {
	user, err := repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Verifica a senha
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("credenciais inválidas")
	}

	// Gera o token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Logout (no caso de JWT, o logout é gerenciado no frontend, mas pode ser implementado via blacklist)
// Esse método pode ser implementado futuramente, caso se deseje invalidar tokens.
func Logout(token string) error {
	// Aqui poderia ser implementado um sistema de blacklist para invalidar tokens
	return nil
}

// Refresh Token (gera um novo token JWT)
// Verifica se o token está válido e gera um novo token
func RefreshToken(oldToken string) (string, error) {
	// Valida o token existente
	token, err := jwt.Parse(oldToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("token inválido ou expirado")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("erro ao processar token")
	}

	// Gera um novo token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": claims["userID"],
		"role":   claims["role"],
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := newToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
