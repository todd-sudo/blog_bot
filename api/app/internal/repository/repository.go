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
	IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error)
}

type Post interface {
	InsertPost(ctx context.Context, b model.Post) (*model.Post, error)
	DeletePost(ctx context.Context, b model.Post) error
	AllPost(ctx context.Context, userId int) ([]*model.Post, error)
	FindPostByID(ctx context.Context, postID uint64) (*model.Post, error)
}

type Category interface {
	InsertCategory(ctx context.Context, c model.Category) (*model.Category, error)
	AllCategory(ctx context.Context, userTgId int) ([]*model.Category, error)
	DeleteCategory(ctx context.Context, category model.Category) error
	FindCategoryByID(ctx context.Context, categoryID uint64) (*model.Category, error)
}

type Repository struct {
	User
	Post
	Category
}

func NewRepository(ctx context.Context, db *gorm.DB, log logging.Logger) *Repository {
	return &Repository{
		User:     NewUserRepository(ctx, db, log),
		Post:     NewPostRepository(ctx, db, log),
		Category: NewCategoryRepository(ctx, db, log),
	}
}
