package repository

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(ctx context.Context, c model.Category) (*model.Category, error)
}

type categoryConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

func NewCategoryRepository(ctx context.Context, dbConn *gorm.DB, log logging.Logger) CategoryRepository {
	return &categoryConnection{
		ctx:        ctx,
		connection: dbConn,
		log:        log,
	}
}

func (db *categoryConnection) InsertCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	tx := db.connection.WithContext(ctx)
	tx.Save(&category)
	res := tx.Preload("Post").Find(&category)
	if res.Error != nil {
		db.log.Errorf("insert category error: %v", res.Error)
		return nil, res.Error
	}
	return &category, nil
}
