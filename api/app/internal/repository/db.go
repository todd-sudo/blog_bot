package repository

import (
	"fmt"

	"github.com/todd-sudo/blog_bot/api/internal/config"
	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config, log *logging.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection database: %v", err)
		return nil, err
	}

	err = migrations(db, log)
	if err != nil {
		return nil, err
	}
	log.Info("Migration Successfully")

	return db, nil
}

func migrations(db *gorm.DB, log *logging.Logger) error {
	err := db.AutoMigrate(&model.User{}, &model.Category{}, &model.Post{})
	if err != nil {
		log.Errorf("Migrate error: %v", err)
		return err
	}
	return nil
}
