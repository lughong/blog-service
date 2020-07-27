package v1

import "github.com/gin-gonic/gin"

type TagHandler struct{}

func NewTagHandler() *TagHandler {
	return &TagHandler{}
}
func (t *TagHandler) Create(c *gin.Context) {}
func (t *TagHandler) Update(c *gin.Context) {}
func (t *TagHandler) Delete(c *gin.Context) {}
func (t *TagHandler) List(c *gin.Context)   {}
func (t *TagHandler) Get(c *gin.Context)    {}
