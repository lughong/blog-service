package model

import "github.com/lughong/blog-service/pkg/app"

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
