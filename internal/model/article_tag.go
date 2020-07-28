package model

type ArticleTag struct {
	*Model

	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (at ArticleTag) TableName() string {
	return "article_tag"
}
