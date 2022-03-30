package repository

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(ctx context.Context, c model.Category) (*model.Category, error)
	DeleteCategory(ctx context.Context, category model.Category, userTgId int) error
	AllCategory(ctx context.Context, userId int) ([]*model.Category, error)
	FindCategoryByID(ctx context.Context, categoryID uint64) (*model.Category, error)
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

func (db *categoryConnection) FindCategoryByID(ctx context.Context, categoryID uint64) (*model.Category, error) {
	tx := db.connection.WithContext(ctx)
	var category model.Category
	res := tx.Find(&category, categoryID)
	if res.Error != nil {
		db.log.Errorf("find category by id error %v", res.Error)
		return nil, res.Error
	}
	return &category, nil
}

func (db *categoryConnection) InsertCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	tx := db.connection.WithContext(ctx)
	tx.Save(&category)
	res := tx.Joins("User").Find(&category)
	if res.Error != nil {
		db.log.Errorf("insert category error: %v", res.Error)
		return nil, res.Error
	}
	return &category, nil
}

// Все category
func (db *categoryConnection) AllCategory(ctx context.Context, userTgId int) ([]*model.Category, error) {
	tx := db.connection.WithContext(ctx)
	var categories []*model.Category
	res := tx.Joins("Posts").Joins("User").Preload("User").Where(
		`"User"."user_tg_id" = ?`,
		userTgId,
	).Find(&categories)
	if res.Error != nil {
		db.log.Errorf("get all categories error %v", res.Error)
		return nil, res.Error
	}
	return categories, nil
}

// Удаление category
func (db *categoryConnection) DeleteCategory(ctx context.Context, category model.Category, userId int) error {
	tx := db.connection.WithContext(ctx)
	res := tx.Select("Posts").Delete(&category)
	if res.Error != nil {
		db.log.Errorf("delete category error %v", res.Error)
		return res.Error
	}
	return nil
}
