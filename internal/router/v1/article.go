package v1

import "github.com/gin-gonic/gin"

type ArticleHandler struct{}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}
func (a *ArticleHandler) Create(c *gin.Context) {}
func (a *ArticleHandler) Update(c *gin.Context) {}
func (a *ArticleHandler) Delete(c *gin.Context) {}
func (a *ArticleHandler) List(c *gin.Context)   {}
func (a *ArticleHandler) Get(c *gin.Context)    {}
