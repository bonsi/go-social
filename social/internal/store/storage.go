package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("Resource not found.")
	ErrConflict          = errors.New("Resource already exists.")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followUserID, userID int64) error
		Unfollow(ctx context.Context, unfollowUserID, userID int64) error
	}
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetadata, error)
		DeleteByID(context.Context, int64) error
		Update(context.Context, *Post) error
	}
	Users interface {
		Activate(context.Context, string) error
		Create(context.Context, *sql.Tx, *User) error
		GetByID(context.Context, int64) (*User, error)
		CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error
		Delete(context.Context, int64) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{

		Comments:  &PostgresCommentStore{db},
		Followers: &PostgresFollowerStore{db},
		Posts:     &PostgresPostStore{db},
		Users:     &PostgresUserStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
