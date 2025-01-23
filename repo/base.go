package repo

import (
	"context"
	"gorm.io/gorm"
)

// BaseRepo DB通用逻辑接口
type BaseRepo[T any] interface {
	Get(ctx context.Context, id int) (*T, error)                                                              // 查询
	List(ctx context.Context, condition string, page, pageSize int, args ...interface{}) (int64, []*T, error) // 列表
	Insert(ctx context.Context, value *T) error                                                               // 新增
	Update(ctx context.Context, value *T) error                                                               // 修改
	Delete(ctx context.Context, value *T) error                                                               // 删除
	CustomUpdate(ctx context.Context, sql string, args ...interface{}) error                                  // 执行sql
}

// BaseRepoImpl DB通用操作结构体
type BaseRepoImpl[T any] struct {
	db *gorm.DB
}

// NewBaseRepo DB通用操作结构体初始化
func NewBaseRepo[T any](db *gorm.DB) BaseRepo[T] {
	return &BaseRepoImpl[T]{db: db}
}

// Get 根据ID查询
func (h *BaseRepoImpl[T]) Get(ctx context.Context, id int) (*T, error) {
	var result T
	err := h.db.WithContext(ctx).First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List 列表
func (h *BaseRepoImpl[T]) List(ctx context.Context, condition string, page, pageSize int, args ...interface{}) (int64, []*T, error) {
	var count int64
	var tmp *T
	err := h.db.WithContext(ctx).Model(tmp).Where(condition, args...).Count(&count).Error
	if err != nil {
		return 0, nil, err
	}
	var data []*T
	err = h.db.Limit(pageSize).Offset((page-1)*pageSize).Where(condition, args...).Find(&data).Error
	return count, data, err
}

// Insert 新增
func (h *BaseRepoImpl[T]) Insert(ctx context.Context, value *T) error {
	return h.db.WithContext(ctx).Create(value).Error
}

// Update 修改
func (h *BaseRepoImpl[T]) Update(ctx context.Context, value *T) error {
	return h.db.WithContext(ctx).Save(value).Error
}

// Delete 删除
func (h *BaseRepoImpl[T]) Delete(ctx context.Context, value *T) error {
	return h.db.WithContext(ctx).Delete(value).Error
}

// CustomUpdate 执行sql
func (h *BaseRepoImpl[T]) CustomUpdate(ctx context.Context, sql string, args ...interface{}) error {
	return h.db.WithContext(ctx).Exec(sql, args...).Error
}
