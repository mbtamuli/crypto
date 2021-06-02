package sqlite

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mbtamuli/telegram-bot/model"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) User(id uuid.UUID) (model.User, error) {
	var u model.User
	if err := s.Get(&u, `SELECT * FROM users WHERE id=$1`, id); err != nil {
		return model.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return u, nil
}

func (s *UserStore) UserByUsername(username string) (model.User, error) {
	var u model.User
	err := s.Get(&u, `SELECT * FROM users WHERE username = ?`, username)
	if err != nil {
		return model.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return u, nil
}

func (s *UserStore) Users() ([]model.User, error) {
	var us []model.User
	if err := s.Select(&us, `SELECT * FROM users`); err != nil {
		return []model.User{}, fmt.Errorf("error getting users: %w", err)
	}
	return us, nil
}

func (s *UserStore) CreateUser(u *model.User) error {
	if _, err := s.NamedExec(`INSERT INTO users(id, username) VALUES (:id, :username)`, u); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (s *UserStore) UpdateUser(u *model.User) error {
	if _, err := s.NamedExec(`UPDATE users SET chat_id = :chat_id WHERE id = :id`, u); err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (s *UserStore) DeleteUser(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM users WHERE id = ?`, id); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (s *UserStore) DeleteUsers() error {
	if _, err := s.Exec(`Delete FROM users`); err != nil {
		return fmt.Errorf("error deleting all users: %w", err)
	}
	return nil
}
