package service

import (
	"github.com/lughong/blog-service/internal/dao"
	"github.com/lughong/blog-service/internal/model"
	"github.com/lughong/blog-service/pkg/app"
)

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state,default=1"`
	Tag           *model.Tag `json:"tag"`
}

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleCreateRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=100"`
	Content       string `form:"content" binding:"required,min=2,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=2,max=100"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleUpdateRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=100"`
	Content       string `form:"content" binding:"required,min=2,max=255"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=2,max=100"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleDeleteRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (src Service) GetArticle(params *ArticleRequest) (*Article, error) {
	article, err := src.dao.GetArticle(params.ID, params.State)
	if err != nil {
		return nil, err
	}

	articleTag, err := src.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}

	tag, err := src.dao.GetTag(articleTag.TagID, params.State)
	if err != nil {
		return nil, err
	}

	return &Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           &tag,
	}, nil
}

func (src Service) GetArticleList(params *ArticleListRequest, pager app.Pager) ([]*Article, int, error) {
	articleCount, err := src.dao.CountArticleListByTagID(params.TagID, params.State)
	if err != nil {
		return nil, 0, err
	}

	articles, err := src.dao.GetArticleListByTagID(params.TagID, params.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ID:            article.ArticleID,
			Title:         article.ArticleTitle,
			Desc:          article.ArticleDesc,
			Content:       article.Content,
			CoverImageUrl: article.CoverImageUrl,
			Tag: &model.Tag{
				Name: article.TagName,
				Model: &model.Model{
					ID: article.TagID,
				},
			},
		})
	}

	return articleList, articleCount, nil
}

func (src Service) CreateArticle(params *ArticleCreateRequest) error {
	article, err := src.dao.CreateArticle(&dao.Article{
		Title:         params.Title,
		Desc:          params.Desc,
		Content:       params.Content,
		CoverImageUrl: params.CoverImageUrl,
		State:         params.State,
		CreatedBy:     params.CreatedBy,
	})
	if err != nil {
		return err
	}

	return src.dao.CreateArticleTag(article.ID, params.TagID, params.CreatedBy)
}

func (src Service) UpdateArticle(params *ArticleUpdateRequest) error {
	err := src.dao.UpdateArticle(&dao.Article{
		ID:            params.ID,
		Title:         params.Title,
		Desc:          params.Desc,
		Content:       params.Content,
		CoverImageUrl: params.CoverImageUrl,
		State:         params.State,
		ModifiedBy:    params.ModifiedBy,
	})
	if err != nil {
		return err
	}

	return src.dao.UpdateArticleTag(params.ID, params.TagID, params.ModifiedBy)
}

func (src Service) DeleteArticle(params *ArticleDeleteRequest) error {
	err := src.dao.DeleteArticle(params.ID)
	if err != nil {
		return err
	}

	return src.dao.DeleteArticleTag(params.ID)
}
