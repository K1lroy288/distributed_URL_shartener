package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Owner_id   int
	Short_code string `gorm:"unique"`
	Long_url   string
}
