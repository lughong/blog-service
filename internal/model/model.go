package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/lughong/blog-service/global"
)

type Model struct {
	Id         uint32 `json:"id"`
	CreatedOn  uint32 `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn uint32 `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
	DeletedOn  uint32 `json:"deleted_on"`
}

func NewDBEngine() (*gorm.DB, error) {
	db, err := gorm.Open(
		global.DatabaseSetting.DBType,
		fmt.Sprintf(
			"%s:%s@/%s?charset=%s&parseTime=%t&loc=%s",
			global.DatabaseSetting.Username,
			global.DatabaseSetting.Password,
			global.DatabaseSetting.DBName,
			global.DatabaseSetting.Charset,
			global.DatabaseSetting.ParseTime,
			global.DatabaseSetting.Loc,
		),
	)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	return db, nil
}
