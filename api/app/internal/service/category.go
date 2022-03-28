package service

import (
	"context"

	"github.com/mashingan/smapping"
	"github.com/todd-sudo/blog_bot/api/internal/dto"
	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/internal/repository"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
)

type CategoryService interface {
	Insert(ctx context.Context, b dto.CreateCategoryDTO) (*model.Category, error)
	Delete(ctx context.Context, b model.Category) error
	All(ctx context.Context) ([]*model.Category, error)
}

type categoryService struct {
	ctx                context.Context
	categoryRepository repository.CategoryRepository
	log                logging.Logger
}

func NewCategoryService(
	ctx context.Context,
	categoryRepository repository.CategoryRepository,
	log logging.Logger,
) CategoryService {
	return &categoryService{
		ctx:                ctx,
		categoryRepository: categoryRepository,
		log:                log,
	}
}

func (s *categoryService) Insert(ctx context.Context, p dto.CreateCategoryDTO) (*model.Category, error) {
	category := model.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&p))
	if err != nil {
		s.log.Errorf("Failed map %v: ", err)
	}
	categoryM, errI := s.categoryRepository.InsertCategory(ctx, category)
	if errI != nil {
		s.log.Errorf("category insert error: %v", errI)
		return nil, err
	}
	return categoryM, nil
}

func (s *categoryService) Delete(ctx context.Context, c model.Category) error {
	err := s.categoryRepository.DeleteCategory(ctx, c)
	if err != nil {
		s.log.Errorf("post delete error: %v", err)
		return err
	}
	return nil
}

func (s *categoryService) All(ctx context.Context) ([]*model.Category, error) {
	categories, err := s.categoryRepository.AllCategory(ctx)
	if err != nil {
		s.log.Errorf("get all categories error: %v", err)
		return nil, err
	}
	return categories, nil
}
