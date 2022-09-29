package repository

import (
	"database/sql"
	"errors"
	"github.com/labstack/gommon/log"
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
			log.Error("Error while scanning book " + err.Error())
			return nil, errors.New("Unexpected database error")
		}
	}
	return &user, nil
}

func (u *userRepositoryDB) InsertSecretById(capsule domain.SecretCapsule) error {
	_, err := u.db.Exec("UPDATE users SET secret = $1 WHERE id = $2;", capsule.Secret, capsule.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryDB) SelectSecretAndAddCountById(id string) (string, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	var secret string
	query := `SELECT secret FROM users WHERE id=$1`
	err = u.db.QueryRow(query, id).Scan(&secret)
	if err != nil {
		return "", err
	}

	_, err = u.db.Exec("UPDATE users SET show_count = show_count + 1 WHERE id = $1;", id)
	if err != nil {
		return "", err
	}

	if err = tx.Commit(); err != nil {
		return "", err
	}

	return secret, nil
}

func (u *userRepositoryDB) GetShowCountById(id string) (int, error) {
	var showCount int
	query := `SELECT show_count FROM users WHERE id=$1`
	err := u.db.QueryRow(query, id).Scan(&showCount)
	if err != nil {
		return -1, err
	}
	return showCount, nil
}
