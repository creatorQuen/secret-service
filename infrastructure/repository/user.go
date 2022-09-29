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

func (u *userRepositoryDB) CreateUser(user domain.User) (string, error) {
	var lastInsertID string
	err := u.db.QueryRow("INSERT INTO users(id, created_at, email, password, full_name, show_count, secret) VALUES($1, $2, $3, $4, $5, $6, $7) returning id;",
		user.Id, user.CreatedAt, user.Email, user.Password, user.FullName, user.ShowCount, user.Secret).Scan(&lastInsertID)
	if err != nil {
		return "", err
	}

	return lastInsertID, nil
}

func (u *userRepositoryDB) GetUserByEmail(email string) (*domain.User, error) {
	query := `SELECT id, created_at, email, password, full_name, show_count, secret FROM users WHERE email=$1`

	var user domain.User
	err := u.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.ShowCount,
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
