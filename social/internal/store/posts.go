package store

import (
	"context"
	"database/sql"
)

type PostgresPostsStore struct {
	db *sql.DB
}

func (s *PostgresPostsStore) Create(ctx context.Context) error {
	return nil
}
