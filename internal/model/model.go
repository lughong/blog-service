package model

type Model struct {
	ID        uint32 `json:"id"`
	CreateOn  uint32 `json:"create_on"`
	CreateBy  string `json:"create_by"`
	ModifyOn  uint32 `json:"modify_on"`
	ModifyBy  string `json:"modify_by"`
	IsDel     uint8  `json:"is_del"`
	DeletedBy string `json:"delete_by"`
}
