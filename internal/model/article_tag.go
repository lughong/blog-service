package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model

	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (at ArticleTag) TableName() string {
	return "article_tag"
}

func (at ArticleTag) Count(db *gorm.DB) (int, error)          {}
func (at ArticleTag) Get(db *gorm.DB) (ArticleTag, error)     {}
func (at ArticleTag) List(db *gorm.DB) ([]*ArticleTag, error) {}
func (at ArticleTag) Create(db *gorm.DB) error {
	return db.Create(&at).Error
}
func (at ArticleTag) Update(db *gorm.DB, values ...interface{}) error {
	return db.Model(&at).Where("id=? AND is_del=?", at.ID, 0).Update(values).Error
}
func (at ArticleTag) Delete(db *gorm.DB) error {
	return db.Where("id=? AND is_del=?", at.ID, 0).Delete(&at).Error
}
