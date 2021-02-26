package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/pkg/app"
	"github.com/lughong/blog-service/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallerFrames().Errorf("panic recover error. %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.InternalServerError)
				c.Abort()
			}
		}()

		c.Next()
	}
}
