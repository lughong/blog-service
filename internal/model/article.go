package model

type Article struct {
	*Model

	Title    string `json:"title"`
	Describe string `json:"describe"`
	Content  string `json:"content"`
}

func NewArticle() Article {
	return Article{}
}

func (a Article) TableName() string {
	return "article"
}
