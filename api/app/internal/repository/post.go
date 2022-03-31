package repository

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"gorm.io/gorm"
)

type PostRepository interface {
	InsertPost(ctx context.Context, b model.Post) (*model.Post, error)
	AllPost(ctx context.Context, userId int) ([]*model.Post, error)
	DeletePost(ctx context.Context, post model.Post) error
	FindPostByID(ctx context.Context, postID uint64) (*model.Post, error)
}

type postConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewPostRepository(ctx context.Context, dbConn *gorm.DB, log logging.Logger) PostRepository {
	return &postConnection{
		ctx:        ctx,
		connection: dbConn,
		log:        log,
	}
}

// Добавление item
func (db *postConnection) InsertPost(ctx context.Context, post model.Post) (*model.Post, error) {
	tx := db.connection.WithContext(ctx)
	tx.Save(&post)
	// .Preload("User").Preload("Category")
	res := tx.Find(&post)
	if res.Error != nil {
		db.log.Errorf("insert post error: %v", res.Error)
		return nil, res.Error
	}
	return &post, nil
}

// Все посты
func (db *postConnection) AllPost(ctx context.Context, userId int) ([]*model.Post, error) {
	tx := db.connection.WithContext(ctx)
	var posts []*model.Post
	res := tx.Preload("User").Where(
		`"user_id" = ?`,
		userId,
	).Preload("Category").Find(&posts)
	if res.Error != nil {
		db.log.Errorf("get all posts error %v", res.Error)
		return nil, res.Error
	}
	return posts, nil
}

// Удаление поста
func (db *postConnection) DeletePost(ctx context.Context, post model.Post) error {
	tx := db.connection.WithContext(ctx)
	res := tx.Delete(&post)
	if res.Error != nil {
		db.log.Errorf("delete post error %v", res.Error)
		return res.Error
	}
	return nil
}

func (db *postConnection) FindPostByID(ctx context.Context, postID uint64) (*model.Post, error) {
	tx := db.connection.WithContext(ctx)
	var post model.Post
	res := tx.Preload("User").Preload("Category").Find(&post, postID)
	if res.Error != nil {
		db.log.Errorf("find post by id error %v", res.Error)
		return nil, res.Error
	}
	return &post, nil
}
