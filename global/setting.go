package global

import (
	"Opendulum/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettings
	CrawlerSetting  *setting.CrawlerSettings
	DBEngine        *gorm.DB
)
