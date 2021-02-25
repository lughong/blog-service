package global

import (
	"github.com/lughong/blog-service/pkg/logger"
	"github.com/lughong/blog-service/pkg/setting"
)

var (
	Logger *logger.Logger

	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JwtSetting      *setting.JwtSettingS
)
