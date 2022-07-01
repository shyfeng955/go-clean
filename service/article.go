package service

import (
	"context"
	"fhz/model"
	"fhz/repo"
	"time"
)

type ArticleService struct {
	repo repo.IArticleRepo
}

func NewArticleService(repo repo.IArticleRepo) IArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (a ArticleService) Fetch(ctx context.Context, createData time.Time, num int) ([]model.Article, error) {
	if num == 0 {
		num = 10
	}

	res, err := a.repo.Fetch(ctx, createData, num)
	if err != nil {
		return nil, err
	}

	return res, nil
}
