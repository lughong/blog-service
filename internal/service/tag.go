package service

type Tag struct{}

type TagCountRequest struct {
	Name  string `form:"name" binding:"omitempty,min=3,max=100"`
	State uint8  `form:"state" binding:"required,oneof=0 1"`
}

type TagListRequest struct {
	Name     string `form:"name" binding:"omitempty,min=3,max=100"`
	State    uint8  `form:"state" binding:"required,oneof=0 1"`
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"page_size" binding:"required"`
}

type TagUpdateRequest struct {
	ID    uint32 `form:"id" binding:"required"`
	Name  string `form:"name" binding:"omitempty,minlength=3,maxlength=100"`
	State uint8  `form:"state" binding:"omitempty,oneof=0 1"`
}

type TagCreateRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	State    uint8  `form:"state,default=1" binding:"required,oneof=0 1"`
	CreateBy string `form:"create_by" binding:"required,min=3,max=100"`
}

type TagDeleteRequest struct {
	ID       uint32 `form:"id" binding:"required"`
	DeleteBy string `form:"delete_by" binding:"required,min=3,max=100"`
}
