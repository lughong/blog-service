package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lughong/blog-service/internal/router/v1"
)

type Router struct {
	g *gin.Engine
}

func NewRouter() *Router {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	tagGroup := g.Group("v1")
	{
		tag := v1.NewTagHandler()
		tagGroup.GET("/tags/{:id}", tag.Get)
		tagGroup.GET("/tags", tag.List)
		tagGroup.POST("/tags", tag.Create)
		tagGroup.DELETE("/tags/{:id}", tag.Delete)
		tagGroup.PUT("/tags/{:id}", tag.Update)
	}

	articleGroup := g.Group("v1")
	{
		article := v1.NewArticleHandler()
		articleGroup.GET("/articles/{:id}", article.Get)
		articleGroup.GET("/articles", article.List)
		articleGroup.POST("/articles", article.Create)
		articleGroup.DELETE("/articles/{:id}", article.Delete)
		articleGroup.PUT("/articles/{:id}", article.Update)
	}

	return &Router{
		g: g,
	}
}

func (r *Router) Run(addr string) error {

	if err := r.g.Run(addr); err != nil {
		return err
	}

	return nil
}
