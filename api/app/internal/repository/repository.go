package repository

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"gorm.io/gorm"
)

type User interface {
	InsertUser(ctx context.Context, user model.User) (*model.User, error)
	ProfileUser(ctx context.Context, userID string) (*model.User, error)
}

type Post interface {
	InsertItem(ctx context.Context, b model.Post) (*model.Post, error)
	DeleteItem(ctx context.Context, b model.Post) error
	AllItem(ctx context.Context) ([]*model.Post, error)
}

type Repository struct {
	User
	Post
}

func NewRepository(ctx context.Context, db *gorm.DB, log logging.Logger) *Repository {
	return &Repository{
		// User: NewUserRepository(ctx, db),
		Post: NewPostRepository(ctx, db, log),
	}
}
