package global

import (
	"Opendulum/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettings
	DBEngine        *gorm.DB
)
