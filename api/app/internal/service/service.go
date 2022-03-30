package service

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/dto"
	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/internal/repository"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
)

type Post interface {
	Insert(ctx context.Context, b dto.PostCreateDTO) (*model.Post, error)
	Delete(ctx context.Context, b model.Post) error
	All(ctx context.Context) ([]*model.Post, error)
	IsAllowedToEdit(ctx context.Context, userID string, postID uint64) (bool, error)
}

type User interface {
	Insert(ctx context.Context, user dto.CreateUserDTO) (*model.User, error)
	Profile(ctx context.Context, userID string) (*model.User, error)
	IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error)
}

type Category interface {
	Insert(ctx context.Context, b dto.CreateCategoryDTO) (*model.Category, error)
	Delete(ctx context.Context, b model.Category, userTgId int) error
	All(ctx context.Context, userTgId int) ([]*model.Category, error)
	IsAllowedToEdit(ctx context.Context, userID string, categoryID uint64) (bool, error)
}

type Service struct {
	Post
	User
	Category
}

func NewService(ctx context.Context, r repository.Repository, log logging.Logger) *Service {
	return &Service{
		Post:     NewPostService(ctx, r.Post, log),
		User:     NewUserService(ctx, r.User, log),
		Category: NewCategoryService(ctx, r.Category, log),
	}
}
