package repositories

import (
	"atlanta-site/models"
	"database/sql"
	"errors"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

// Buscar usuário por email
func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, email, password, role FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	return &user, nil
}
