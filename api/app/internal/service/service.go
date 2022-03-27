package service

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/dto"
	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/internal/repository"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
)

type Item interface {
	Insert(ctx context.Context, b dto.ItemCreateDTO) (*model.Post, error)
	Delete(ctx context.Context, b model.Post) error
	All(ctx context.Context) ([]*model.Post, error)
}

type User interface {
	CreateUser(ctx context.Context, user dto.CreateUserDTO) (*model.User, error)
	Profile(ctx context.Context, userID string) (*model.User, error)
}

type Service struct {
	Item
	User
}

func NewService(ctx context.Context, r repository.Repository, log logging.Logger) *Service {
	return &Service{
		// Item: NewItemService(ctx, r.Item),
		// User: NewUserService(ctx, r.User),
	}
}
