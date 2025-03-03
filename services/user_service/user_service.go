// services/user_services/user_service.go
package services

import (
	"atlanta-site/models"
	"atlanta-site/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUserService valida os dados do usuário e chama o repositório para inserção
func RegisterUserService(user models.User) error {
	// Validar se o usuário já existe
	existingUser, _ := repositories.GetUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("usuário já existe")
	}

	// Criptografar senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao criptografar senha")
	}
	user.Password = string(hashedPassword)

	// Inserir no repositório
	if err := repositories.InsertUser(user); err != nil {
		return errors.New("erro ao criar usuário")
	}

	return nil
}

// PromoteUserToAdminService promove o usuário para administrador
func PromoteUserToAdminService(id string) error {
	// Chama o repositório para buscar o usuário
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return err
	}

	// Alterar o papel para "admin"
	user.Role = "admin"

	// Atualizar no repositório
	if err := repositories.UpdateUserRole(user); err != nil {
		return err
	}

	return nil
}
