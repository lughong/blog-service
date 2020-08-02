package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/lughong/blog-service/global"
	"github.com/lughong/blog-service/pkg/str"
)

func GetPage(c *gin.Context) uint8 {
	page := c.Query("page")
	if page == "" {
		return 1
	}

	return str.MustToUint8(page)
}

func GetPageSize(c *gin.Context) uint8 {
	pz := c.Query("page_size")
	if pz == "" {
		return global.AppSetting.DefaultPageSize
	}

	pageSize := str.MustToUint8(pz)
	if pageSize < 0 {
		return global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize uint8) uint8 {
	if page < 1 {
		page = 1
	}

	return (page - 1) * pageSize
}
