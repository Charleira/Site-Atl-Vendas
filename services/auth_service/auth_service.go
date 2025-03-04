package auth_service

import (
	"atlanta-site/repositories"
	"atlanta-site/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	Repo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{Repo: repo}
}

var jwtSecret = []byte("SUA_CHAVE_SECRETA")

// Login do usuário
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmail(email)
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
func (s *AuthService) Logout(token string) error {
	// Aqui poderia ser implementado um sistema de blacklist para invalidar tokens
	return nil
}

// Refresh Token (gera um novo token JWT)
func (s *AuthService) RefreshToken(oldToken string) (string, error) {
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
