package repository

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"gorm.io/gorm"
)

//BookRepository is a ....
type PostRepository interface {
	InsertItem(ctx context.Context, b model.Post) (*model.Post, error)
	AllItem(ctx context.Context) ([]*model.Post, error)
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
func (db *postConnection) InsertItem(ctx context.Context, post model.Post) (*model.Post, error) {
	tx := db.connection.WithContext(ctx)
	tx.Save(&post)
	res := tx.Preload("User").Find(&post)
	if res.Error != nil {
		db.log.Errorf("inset post error: %v", res.Error)
		return nil, res.Error
	}
	return &post, nil
}

// Все посты
func (db *postConnection) AllItem(ctx context.Context) ([]*model.Post, error) {
	tx := db.connection.WithContext(ctx)
	var posts []*model.Post
	res := tx.Preload("User").Find(&posts)
	if res.Error != nil {
		db.log.Errorf("get all posts error %v", res.Error)
		return nil, res.Error
	}
	return posts, nil
}

// Удаление поста
func (db *postConnection) DeleteItem(ctx context.Context, post model.Post) error {
	tx := db.connection.WithContext(ctx)
	res := tx.Delete(&post)
	if res.Error != nil {
		db.log.Errorf("delete post error %v", res.Error)
		return res.Error
	}
	return nil
}
