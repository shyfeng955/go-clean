package service

import (
	"context"
	"github.com/shyfeng955/go-clean/models"
)

type UserService interface {
	GetUserInfo(ctx context.Context, id int) (*models.User, error)
}
