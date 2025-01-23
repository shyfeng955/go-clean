package repo

import (
	"context"
	"github.com/shyfeng955/go-clean/models"
	"gorm.io/gorm"
)

type userRepoImpl struct {
	mysql *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepoImpl{mysql: db}
}

func (u *userRepoImpl) GetUserById(ctx context.Context, id int, fields []string) (*models.User, error) {
	var user *models.User
	err := u.mysql.WithContext(ctx).
		Select(fields).
		Where(&models.User{Id: id}).
		First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
