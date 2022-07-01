package repo

import (
	"context"
	"fhz/model"
	"gorm.io/gorm"
	"time"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepo {
	return &ArticleRepository{db}
}

func (r ArticleRepository) Fetch(ctx context.Context, createData time.Time, num int) ([]model.Article, error) {

	var (
		record []model.Article
		err    error
	)

	err = r.DB.Model(&model.Article{}).
		Select("id,title,content,create_at,update_at").
		Where("create_at > ?", createData).Limit(num).Find(&record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}
