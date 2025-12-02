package repository

import (
	"shortener-service/model"

	"gorm.io/gorm"
)

type ShortenerRepository struct {
	DB *gorm.DB
}

func NewShortenerRepository(db *gorm.DB) *ShortenerRepository {
	return &ShortenerRepository{DB: db}
}

func (r *ShortenerRepository) SaveCode(url *model.Url) (bool, error) {
	var exist model.Url
	res := r.DB.Where("short_code = ?", url.Short_code).First(&exist).Error
	return res == nil, r.DB.Create(url).Error
}

func (r *ShortenerRepository) GetLink(code string) (*model.Url, error) {
	var url model.Url
	err := r.DB.Where("short_code = ?", code).First(&url).Error
	return &url, err
}
