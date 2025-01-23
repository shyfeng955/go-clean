package service

import (
	"context"
	"github.com/shyfeng955/go-clean/models"
	"github.com/shyfeng955/go-clean/repo"
)

type userServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (u *userServiceImpl) GetUserInfo(ctx context.Context, id int) (*models.User, error) {
	// 进行你的业务逻辑处理
	info, err := u.userRepo.GetUserById(ctx, id, []string{"username"})
	if err != nil {
		return nil, err
	}
	return info, nil
}
