package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lughong/blog-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	v1 "github.com/lughong/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiv1 := r.Group("/api/v1")
	{
		t := v1.NewTag()
		apiv1.POST("/tags", t.Create)
		apiv1.PUT("/tags/:id", t.Update)
		apiv1.DELETE("/tags/:id", t.Delete)
		apiv1.GET("/tags", t.List)
		apiv1.GET("/tags/:id", t.Get)

		a := v1.NewArticle()
		apiv1.POST("/articles", a.Create)
		apiv1.PUT("/articles/:id", a.Update)
		apiv1.DELETE("/articles/:id", a.Delete)
		apiv1.GET("/articles", a.List)
		apiv1.GET("/articles/:id", a.Get)
	}

	return r
}
