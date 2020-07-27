package model

type Tag struct {
	*Model

	Name string `json:"name"`
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) TableName() string {
	return "tag"
}
