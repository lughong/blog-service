package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/internal/service"
	"github.com/lughong/blog-service/pkg/app"
	"github.com/lughong/blog-service/pkg/convert"
	"github.com/lughong/blog-service/pkg/errcode"
	"github.com/lughong/blog-service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	f, fileHeader, err := c.Request.FormFile("file")
	filtType := convert.StrTo(c.PostForm("type")).MustToInt()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	if fileHeader == nil || filtType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	srv := service.New(c.Request.Context(), nil)
	fileInfo, err := srv.UploadFile(upload.FileType(filtType), fileHeader, f)
	if err != nil {
		global.Logger.Errorf("service.UploadFile error. %v", err)
		response.ToErrorResponse(errcode.ErrUploadFail)
		return
	}

	response.ToResponse(gin.H{"file_access_url": fileInfo.AccessURL})
	return
}
