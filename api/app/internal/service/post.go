package service

import (
	"context"
	"fmt"

	"github.com/mashingan/smapping"
	"github.com/todd-sudo/blog_bot/api/internal/dto"
	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/internal/repository"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
)

type PostService interface {
	Insert(ctx context.Context, b dto.PostCreateDTO) (*model.Post, error)
	Delete(ctx context.Context, b model.Post) error
	All(ctx context.Context) ([]*model.Post, error)
	IsAllowedToEdit(ctx context.Context, userID string, postID uint64) (bool, error)
}

type postService struct {
	ctx            context.Context
	postRepository repository.PostRepository
	log            logging.Logger
}

func NewPostService(
	ctx context.Context,
	postRepository repository.PostRepository,
	log logging.Logger,
) PostService {
	return &postService{
		ctx:            ctx,
		postRepository: postRepository,
		log:            log,
	}
}

func (s *postService) Insert(ctx context.Context, p dto.PostCreateDTO) (*model.Post, error) {
	post := model.Post{}
	err := smapping.FillStruct(&post, smapping.MapFields(&p))
	if err != nil {
		s.log.Errorf("Failed map %v: ", err)
	}
	postM, errI := s.postRepository.InsertPost(ctx, post)
	if errI != nil {
		s.log.Errorf("post insert error: %v", errI)
		return nil, err
	}
	return postM, nil
}

func (s *postService) Delete(ctx context.Context, p model.Post) error {
	err := s.postRepository.DeletePost(ctx, p)
	if err != nil {
		s.log.Errorf("post delete error: %v", err)
		return err
	}
	return nil
}

func (s *postService) All(ctx context.Context) ([]*model.Post, error) {
	posts, err := s.postRepository.AllPost(ctx)
	if err != nil {
		s.log.Errorf("get all posts error: %v", err)
		return nil, err
	}
	return posts, nil
}

func (s *postService) IsAllowedToEdit(ctx context.Context, userID string, postID uint64) (bool, error) {
	post, err := s.postRepository.FindPostByID(ctx, postID)
	if err != nil {
		s.log.Errorf("is allowed to edit post error: %v", err)
		return false, err
	}
	id := fmt.Sprintf("%v", post.UserID)
	return userID == id, nil
}
