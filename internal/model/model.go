package model

type Model struct {
	ID         uint32 `json:"id"`
	CreatedOn  uint32 `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn uint32 `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
	DeletedOn  uint32 `json:"deleted_on"`
}
