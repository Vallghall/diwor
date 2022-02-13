package storage

import (
	"database/sql"
	"fmt"
	"gitlab.com/Valghall/diwor/internal/users"

	"github.com/jmoiron/sqlx"
)

type AuthPostgress struct {
	db *sqlx.DB
}

func NewAuthPostgress(db *sqlx.DB) *AuthPostgress {
	return &AuthPostgress{db: db}
}
func (ap *AuthPostgress) CreateUser(user users.User) (int, error) {
	var id int

	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash)"+
			" VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := ap.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ap *AuthPostgress) GetUser(username, password string) (users.User, error) {
	var user users.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := ap.db.Get(&user, query, username, password)

	return user, err
}

func (ap *AuthPostgress) GetUserById(id int) (users.User, error) {
	user := users.User{Id: id}

	query := fmt.Sprintf("SELECT name, username FROM %s WHERE id=$1", usersTable)
	err := ap.db.Get(&user, query, id)

	return user, err
}

func (ap *AuthPostgress) LookUpUser(username string) bool {
	var user users.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", usersTable)
	err := ap.db.Get(&user, query, username)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}
