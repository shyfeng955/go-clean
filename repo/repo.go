package repo

import (
	"context"
	"github.com/shyfeng955/go-clean/models"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id int, fields []string) (*models.User, error)
}
