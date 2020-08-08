package model

import "github.com/lughong/blog-service/pkg/app"

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
