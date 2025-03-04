package services

import (
	"atlanta-site/models"
	"atlanta-site/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// GetUserDetailsService busca os detalhes do usuário
func GetUserDetailsService(userID string) (*models.User, error) {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUserDetailsService permite ao usuário atualizar suas informações
func UpdateUserDetailsService(userID string, updatedUser models.User) error {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Atualiza apenas os campos permitidos
	user.Username = updatedUser.Username

	return repositories.UpdateUser(user)
}

// ChangePasswordService permite ao usuário alterar sua senha
func ChangePasswordService(userID, oldPassword, newPassword string) error {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Verifica se a senha antiga está correta
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("senha antiga incorreta")
	}

	// Gera nova senha criptografada
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao criptografar senha")
	}

	user.Password = string(hashedPassword)

	return repositories.UpdateUser(user)
}

// DeleteUserService remove um usuário do sistema
func DeleteUserService(userID string) error {
	return repositories.DeleteUser(userID)
}
