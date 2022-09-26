package repository

import (
	"database/sql"
	"secret-service/domain"
)

type userRepositoryDB struct {
	db *sql.DB
}

func NewUserRepositoryDb(dbClinet *sql.DB) *userRepositoryDB {
	return &userRepositoryDB{dbClinet}
}

func (b *userRepositoryDB) CreateUser(user domain.User) (int, error) {
	var lastInsertID int
	err := b.db.QueryRow("INSERT INTO users(created_at, email, password, full_name, show_count, secret) VALUES($1, $2, $3, $4, $5, $6) returning id;",
		user.CreatedAt, user.Email, user.Password, user.FullName, user.ShowCount, user.Secret).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
