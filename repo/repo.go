package repo

import (
	"context"
	"fhz/model"
	"time"
)

type IArticleRepo interface {
	Fetch(ctx context.Context, createData time.Time, num int) ([]model.Article, error)
}
