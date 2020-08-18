package service

import (
	"github.com/lughong/blog-service/internal/model"
	"github.com/lughong/blog-service/pkg/app"
)

type Tag struct{}

type TagCountRequest struct {
	Name string `form:"name" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"required,oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagCreateRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagUpdateRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=2,max=100"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type TagDeleteRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (src Service) CountTag(params *TagCountRequest) (int, error) {
	return src.dao.CountTag(params.Name, params.State)
}

func (src Service) GetTagList(params *TagListRequest, pager app.Pager) ([]*model.Tag, error) {
	return src.dao.GetTagList(params.Name, params.State, pager.Page, pager.PageSize)
}

func (src Service) CreateTag(params *TagCreateRequest) error {
	return src.dao.CreateTag(params.Name, params.State, params.CreatedBy)
}

func (src Service) UpdateTag(params *TagUpdateRequest) error {
	return src.dao.UpdateTag(params.ID, params.Name, params.State, params.ModifiedBy)
}

func (src Service) DeleteTag(params *TagDeleteRequest) error {
	return src.dao.DeleteTag(params.ID)
}