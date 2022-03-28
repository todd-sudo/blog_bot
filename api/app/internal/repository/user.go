package repository

import (
	"context"

	"github.com/todd-sudo/blog_bot/api/internal/model"
	"github.com/todd-sudo/blog_bot/api/pkg/logging"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(ctx context.Context, user model.User) (*model.User, error)
	ProfileUser(ctx context.Context, userID string) (*model.User, error)
	IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error)
}

type userConnection struct {
	ctx        context.Context
	connection *gorm.DB
	log        logging.Logger
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(ctx context.Context, db *gorm.DB, log logging.Logger) UserRepository {
	return &userConnection{
		ctx:        ctx,
		connection: db,
		log:        log,
	}
}

func (db *userConnection) IsDuplicateUserTGID(ctx context.Context, tgID int) (bool, error) {
	var user *model.User
	res := db.connection.WithContext(ctx).Where("user_tg_id = ?", tgID).Take(&user)
	db.log.Debug(res.Error)
	if res.Error != nil {
		db.log.Error(res.Error)
		return true, res.Error
	}
	return false, nil
}

// Добавление пользователя
func (db *userConnection) InsertUser(ctx context.Context, user model.User) (*model.User, error) {
	tx := db.connection.WithContext(ctx)
	res := tx.Save(&user)
	if res.Error != nil {
		db.log.Errorf("insert user error %v", res.Error)
		return nil, res.Error
	}
	return &user, nil
}

// Вывод профиля пользователя
func (db *userConnection) ProfileUser(ctx context.Context, userID string) (*model.User, error) {
	tx := db.connection.WithContext(ctx)
	var user model.User
	res := tx.Preload("Posts").Preload("Posts.User").Find(&user, userID)
	if res.Error != nil {
		db.log.Errorf("profile user error %v", res.Error)
		return nil, res.Error
	}
	return &user, nil
}
