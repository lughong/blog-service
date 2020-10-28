package model

import (
	"github.com/jinzhu/gorm"

	"github.com/lughong/blog-service/pkg/app"
)

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

type Tag struct {
	*Model

	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "tag"
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag

	if err := db.Where("id = ? AND state = ? AND is_del = ?", t.ID, t.State, 0).First(&tag).Error; err != nil {
		return tag, err
	}

	return tag, nil
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}

	db = db.Where("state=?", t.State)
	if err := db.Model(&t).Where("is_del=?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}

	db = db.Where("state=?", t.State)
	if err := db.Model(&t).Where("is_del=?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values ...interface{}) error {
	return db.Model(&t).Where("id=? AND is_del=?", t.ID, 0).Update(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id=? AND is_del=?", t.ID, 0).Delete(&t).Error
}
