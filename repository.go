package main

import "gorm.io/gorm"

type SortOrder string

const (
	Asc  SortOrder = "ASC"
	Desc SortOrder = "DESC"
)

type PaginatedResponse struct {
	TotalCount int64       `json:"total_count"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	Data       interface{} `json:"data"`
}

type Repository[M any] struct {
	DB *gorm.DB
}

func NewRepository[M any](db *gorm.DB) *Repository[M] {
	return &Repository[M]{DB: db}
}

func (r *Repository[M]) Create(model M) error {
	return r.DB.Create(&model).Error
}

func (r *Repository[M]) FindOrCreate(model M) error {
	return r.DB.FirstOrCreate(&model).Error
}

func (r *Repository[M]) FindByID(id uint) (M, error) {
	var model M
	err := r.DB.First(&model, id).Error
	return model, err
}

func (r *Repository[M]) Update(model M) error {
	return r.DB.Save(&model).Error
}

func (r *Repository[M]) Delete(model M) error {
	return r.DB.Delete(&model).Error
}

type QueryOption func(db *gorm.DB) *gorm.DB

func WithCondition(condition string, args ...interface{}) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(condition, args...)
	}
}

func WithPagination(page, pageSize int) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

func WithOrder(order string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func WithSort(field string, order SortOrder) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(field + " " + string(order))
	}
}

func WithCount(count *int64) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Count(count)
	}
}

func (r *Repository[M]) Find(options ...QueryOption) ([]M, error) {
	var models []M
	db := r.DB.Model(new(M))
	for _, option := range options {
		db = option(db)
	}
	err := db.Find(&models).Error
	return models, err
}

func (r *Repository[M]) FindOne(options ...QueryOption) (M, error) {
	var model M
	db := r.DB.Model(new(M))
	for _, option := range options {
		db = option(db)
	}
	err := db.First(&model).Error
	return model, err
}

func (r *Repository[M]) UpdateWithCondition(model M, condition string, args ...interface{}) error {
	return r.DB.Model(&model).Where(condition, args...).Updates(&model).Error
}
