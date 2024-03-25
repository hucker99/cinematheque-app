package repository

import (
	"fmt"
	"github.com/hucker99/cinematheque-app/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash, role) VALUES ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Email, user.Password, "user")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *AuthPostgres) GetUser(email, password string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", userTable)

	err := r.db.Get(&user, query, email, password)

	return user, err
}
