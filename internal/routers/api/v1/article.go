package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/internal/dao"
	"github.com/lughong/blog-service/internal/service"
	"github.com/lughong/blog-service/pkg/app"
	"github.com/lughong/blog-service/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 新增文章
// @Produce json
// @Param title body string true "文章标题" minlength(3) maxlength(100)
// @Param desc body string true "文章描述"
// @Param content body string true "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	var params service.ArticleCreateRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	srv := service.New(c.Request.Context(), d)
	if err := srv.CreateArticle(&params); err != nil {
		global.Logger.Errorf("srv.CreateArticle error. %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(nil)
	return
}

// @Summary 更新文章
// @Produce json
// @Param id path int true "文章ID"
// @Param title body string false "文章标题" minlength(3) maxlength(100)
// @Param content body string false "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	var params service.ArticleUpdateRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	srv := service.New(c.Request.Context(), d)
	if err := srv.UpdateArticle(&params); err != nil {
		global.Logger.Errorf("srv.DeleteArticle error. %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(nil)
	return
}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	var params service.ArticleDeleteRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	srv := service.New(c.Request.Context(), d)
	if err := srv.DeleteArticle(&params); err != nil {
		global.Logger.Errorf("srv.DeleteArticle error. %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(nil)
	return
}

// @Summary 获取多篇文章
// @Produce json
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	var params service.ArticleListRequest
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	srv := service.New(c.Request.Context(), d)

	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	articles, count, err := srv.GetArticleList(&params, pager)
	if err != nil {
		global.Logger.Errorf("srv.GetArticleList error. %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}

	response.ToResponseList(articles, count)
	return
}

// @Summary 获取一篇文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	params := service.ArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	srv := service.New(c.Request.Context(), d)
	article, err := srv.GetArticle(&params)
	if err != nil {
		global.Logger.Errorf("service.GetArticle error. %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
	return
}
