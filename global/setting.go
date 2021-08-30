package global

import (
	"github.com/go-programming-tour/blog-service/pkg/logger"
	"github.com/go-programming-tour/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSetting
	EmailSetting    *setting.EmailSetting
)
