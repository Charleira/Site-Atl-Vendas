// repositories/user_repositories/user_repository.go
package repositories

import (
	"atlanta-site/config"
	"atlanta-site/models"
	"database/sql"
	"errors"
)

// InsertUser insere um novo usuário no banco
func InsertUser(user models.User) error {
	query := `INSERT INTO users (username, password, role) VALUES (?, ?, ?)`
	_, err := config.DB.Exec(query, user.Username, user.Password, user.Role)
	return err
}

// GetUserByUsername busca um usuário pelo nome de usuário
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, role FROM users WHERE username = ?`
	row := config.DB.QueryRow(query, username)

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID busca um usuário pelo ID
func GetUserByID(id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password, role FROM users WHERE id = ?`
	row := config.DB.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return nil, errors.New("usuário não encontrado")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserRole atualiza o papel do usuário
func UpdateUserRole(user *models.User) error {
	query := `UPDATE users SET role = ? WHERE id = ?`
	_, err := config.DB.Exec(query, user.Role, user.ID)
	return err
}

// UpdateUser atualiza os dados do usuário no banco
func UpdateUser(user *models.User) error {
	query := `UPDATE users SET username = ?, password = ? WHERE id = ?`
	_, err := config.DB.Exec(query, user.Username, user.Password, user.ID)
	return err
}

// DeleteUser remove um usuário do banco
func DeleteUser(userID string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := config.DB.Exec(query, userID)
	return err
}
