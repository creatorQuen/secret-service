package repository

import (
	"database/sql"
	"errors"
	"log"
	"secret-service/domain"
)

type userRepositoryDB struct {
	db *sql.DB
}

func NewUserRepositoryDb(dbClinet *sql.DB) *userRepositoryDB {
	return &userRepositoryDB{dbClinet}
}

func (u *userRepositoryDB) CreateUser(user domain.User) (int, error) {
	var lastInsertID int
	err := u.db.QueryRow("INSERT INTO users(created_at, email, password, full_name, show_count, secret) VALUES($1, $2, $3, $4, $5, $6) returning id;",
		user.CreatedAt, user.Email, user.Password, user.FullName, user.ShowCount, user.Secret).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (u *userRepositoryDB) GetUserByEmail(email string) (*domain.User, error) {
	query := `SELECT id, created_at, email, password, full_name, show_count, unique_id, secret FROM users WHERE email=$1`

	var user domain.User
	err := u.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.ShowCount,
		&user.UniqueId,
		&user.Secret,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			log.Println("Error while scanning book " + err.Error())
			return nil, errors.New("Unexpected database error")
		}
	}
	return &user, nil
}
