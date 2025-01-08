package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Follower struct {
	FollowUserID int64 `json:"follow_user_id"`
	UserID       int64 `json:"user_id"`
}

type PostgresFollowerStore struct {
	db *sql.DB
}

func (s *PostgresFollowerStore) Follow(ctx context.Context, followUserID, userID int64) error {
	query := `
		INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followUserID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}
	return err
}

func (s *PostgresFollowerStore) Unfollow(ctx context.Context, unfollowUserID, userID int64) error {
	query := `
		DELETE FROM followers WHERE user_id = $1 AND follower_id = $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, unfollowUserID)
	return err
}
