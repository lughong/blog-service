package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/internal/dao"
	"github.com/lughong/blog-service/internal/service"
	"github.com/lughong/blog-service/pkg/app"
	"github.com/lughong/blog-service/pkg/convert"
	"github.com/lughong/blog-service/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	params := service.TagCreateRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvaildParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	src := service.New(c.Request.Context(), d)
	if err := src.CreateTag(&params); err != nil {
		global.Logger.Errorf("src.CreateTag error. %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	params := service.TagUpdateRequest{ID: convert.StrTo(c.Param("id")).MustToUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvaildParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	src := service.New(c.Request.Context(), d)
	if err := src.UpdateTag(&params); err != nil {
		global.Logger.Errorf("src.UpdateTag error. %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	params := service.TagDeleteRequest{ID: convert.StrTo(c.Param("id")).MustToUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvaildParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	src := service.New(c.Request.Context(), d)
	if err := src.DeleteTag(&params); err != nil {
		global.Logger.Errorf("src.DeleteTag error. %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	params := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvaildParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	src := service.New(c.Request.Context(), d)

	totalRows, err := src.CountTag(&service.TagCountRequest{Name: params.Name, State: params.State})
	if err != nil {
		global.Logger.Errorf("src.CountTag error. %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
	}

	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	tags, err := src.GetTagList(&params, pager)
	if err != nil {
		global.Logger.Errorf("src.GetTagList error. %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

// @Summary 获取一个标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context) {}
