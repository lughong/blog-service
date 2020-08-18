package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		engine: db,
	}
}
