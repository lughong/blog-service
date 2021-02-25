package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lughong/blog-service/pkg/app"
	"github.com/lughong/blog-service/pkg/errcode"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token = ""
			code  = errcode.Success
		)

		if s, exists := c.GetQuery("token"); exists {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			code = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errcode.ErrUnAuthorizedTokenExpire
				default:
					code = errcode.ErrUnAuthorizedTokenFail
				}
			}
		}

		if code != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(code)
			c.Abort()
			return
		}

		c.Next()
	}
}
