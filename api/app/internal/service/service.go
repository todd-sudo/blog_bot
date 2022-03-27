package service

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/repository"
)

// type Item interface {
// 	Insert(ctx context.Context, b dto.ItemCreateDTO) (*model.Item, error)
// 	Delete(ctx context.Context, b model.Item) error
// 	All(ctx context.Context) ([]*model.Item, error)
// }

// type User interface {
// 	Update(ctx context.Context, user dto.UserUpdateDTO) (*model.User, error)
// 	Profile(ctx context.Context, userID string) (*model.User, error)
// }

type Service struct {
}

func NewService(ctx context.Context, r repository.Repository) *Service {
	return &Service{
		// Item: NewItemService(ctx, r.Item),
		// User: NewUserService(ctx, r.User),
	}
}
