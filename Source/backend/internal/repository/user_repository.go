package repository

import (
	"database/sql"
	"sports-backend/Source/backend/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(username, hashedPassword string) (int, error) {
	var id int
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, username, hashedPassword).Scan(&id)
	return id, err
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, password_hash FROM users WHERE username = $1`

	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUsername(userID int, username string) error {
	query := `UPDATE users SET username = $1 WHERE id = $2`
	_, err := r.db.Exec(query, username, userID)
	return err
}

func (r *UserRepository) DeleteUser(userID int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, userID)
	return err
}
