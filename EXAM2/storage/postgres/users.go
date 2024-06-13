package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Dilfuza/EXAM2/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) UserGet(id string) (models.Users, error) {
	var user models.Users
	var deletedAt sql.NullString

	err := u.db.QueryRow("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1", id).Scan(
		&user.UserId,
		&user.Name,
		&user.Email,
		&user.Birthday,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&deletedAt,
	)
	if err != nil {
		return user, err
	}

	if deletedAt.Valid {
		user.DeletedAt = &deletedAt.String
	} else {
		user.DeletedAt = nil
	}

	return user, nil
}

func (u *UserRepo) UserCreate(user models.Users) error {
	query := `INSERT INTO users (user_id, name, email, password, birthday, created_at, updated_at, deleted_at)
              VALUES ($1, $2, $3, $4, $5, DEFAULT, DEFAULT, DEFAULT)`
	_, err := u.db.Exec(query, user.UserId, user.Name, user.Email, user.Password, user.Birthday)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func (u *UserRepo) UserUpdate(user models.Users) error {
	query := `UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, updated_at = DEFAULT, deleted_at = $5 WHERE user_id = $6`
	_, err := u.db.Exec(query, user.Name, user.Email, user.Birthday, user.Password, user.DeletedAt, user.UserId)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (u *UserRepo) UserDelete(id string) error {
	_, err := u.db.Exec("DELETE FROM users WHERE user_id = $1", id)
	return err
}
