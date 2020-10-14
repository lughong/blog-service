package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lughong/blog-service/pkg/app"
)

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

type Article struct {
	*Model

	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	State   uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "article"
}

func (a Article) Update(db *gorm.DB, values ...interface{}) error {
	return db.Model(&a).Where("id=? AND is_del=?", a.ID, 0).Update(values).Error
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id=? AND is_del=?", a.ID, 0).Delete(&a).Error
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article

	if err := db.Model(&a).Where("is_del = ?", 0).First(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	db = db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int

	db.Where("state = ?", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
