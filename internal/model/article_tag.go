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

func (at ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag

	err := db.Where("tag_id = ? AND is_del = ?", at.TagID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

func (at ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag

	if err := db.Where("tag_id = ? AND is_del = ?", at.TagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (at ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag

	err := db.Where("article_id IN (?) AND is_del = ?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleTags, nil
}

func (at ArticleTag) Create(db *gorm.DB) error {
	return db.Create(&at).Error
}

func (at ArticleTag) UpdateOne(db *gorm.DB, values ...interface{}) error {
	return db.Model(&at).Where("article_id=? AND is_del=?", at.ArticleID, 0).Limit(1).Updates(values).Error
}

func (at ArticleTag) Delete(db *gorm.DB) error {
	return db.Where("id=? AND is_del=?", at.ID, 0).Delete(&at).Error
}

func (at ArticleTag) DeleteOne(db *gorm.DB) error {
	return db.Where("article_id=? AND is_del=?", at.ArticleID, 0).Limit(1).Delete(&at).Error
}
