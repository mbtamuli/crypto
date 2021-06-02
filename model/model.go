package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID      `db:"id"`
	Username string         `db:"username"`
	ChatID   sql.NullString `db:"chat_id"`
}

type UserStore interface {
	User(username uuid.UUID) (User, error)
	UserByUsername(username string) (User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(username uuid.UUID) error
}

type Store interface {
	UserStore
}
