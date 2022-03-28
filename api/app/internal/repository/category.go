package repository

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(ctx context.Context, c model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, category model.Category) error
	AllCategory(ctx context.Context) ([]*model.Category, error)
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

// func (db *categoryConnection) FindItemByID(ctx context.Context, categoryID uint64) (*model.Category, error) {
// 	tx := db.connection.WithContext(ctx)
// 	var category model.Category
// 	res := tx.Preload("Posts").Find(&category, categoryID)
// 	if res.Error != nil {
// 		db.log.Errorf("find category by id error %v", res.Error)
// 		return nil, res.Error
// 	}
// 	return &category, nil
// }

func (db *categoryConnection) InsertCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	tx := db.connection.WithContext(ctx)
	tx.Save(&category)
	res := tx.Find(&category)
	if res.Error != nil {
		db.log.Errorf("insert category error: %v", res.Error)
		return nil, res.Error
	}
	return &category, nil
}

// Все category
func (db *categoryConnection) AllCategory(ctx context.Context) ([]*model.Category, error) {
	tx := db.connection.WithContext(ctx)
	var categories []*model.Category
	res := tx.Preload("Posts").Find(&categories)
	if res.Error != nil {
		db.log.Errorf("get all categories error %v", res.Error)
		return nil, res.Error
	}
	return categories, nil
}

// Удаление category
func (db *categoryConnection) DeleteCategory(ctx context.Context, category model.Category) error {
	tx := db.connection.WithContext(ctx)
	res := tx.Delete(&category)
	if res.Error != nil {
		db.log.Errorf("delete category error %v", res.Error)
		return res.Error
	}
	return nil
}
