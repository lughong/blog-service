package global

import (
	"github.com/lughong/blog-service/pkg/logger"
	"github.com/lughong/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger
)
