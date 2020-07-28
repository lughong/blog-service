package model

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
