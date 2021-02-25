package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/internal/dao"
	"github.com/lughong/blog-service/internal/service"
	"github.com/lughong/blog-service/pkg/app"
	"github.com/lughong/blog-service/pkg/errcode"
)

type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error. %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	d := dao.New(global.DBEngine)
	srv := service.New(c.Request.Context(), d)
	if err := srv.CheckAuth(&param); err != nil {
		global.Logger.Errorf("srv.CheckAuth error. %v", err)
		response.ToErrorResponse(errcode.UnAuthorization)
		return
	}

	token, e := app.GenerateToken(param.AppKey, param.AppSecret)
	if e != nil {
		global.Logger.Errorf("app.GenerateToken error. %v", e)
		response.ToErrorResponse(errcode.ErrUnAuthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
