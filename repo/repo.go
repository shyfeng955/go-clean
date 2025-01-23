package repo

import (
	"context"
	"github.com/shyfeng955/go-clean/models"
)

type UserRepo interface {
	BaseRepo[models.User] // 继承基础的crud函数
	GetUserById(ctx context.Context, id int, fields []string) (*models.User, error)
}
