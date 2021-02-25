package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth

	if err := db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0).
		First(&auth).Error; err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
