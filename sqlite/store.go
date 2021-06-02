package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		UserStore: &UserStore{DB: db},
	}, nil
}

type Store struct {
	*UserStore
}
