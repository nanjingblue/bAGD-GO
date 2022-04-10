package model

import (
	"Opendulum/global"
	"github.com/jinzhu/gorm"
)

type GoodJingDong struct {
	gorm.Model
	Brand         string `gorm:"type:varchar(50);unique"`
	JDId          string
	Concentration string // 净含量
	ShelfLife     string // 保质期
}

// GetJDId 由brand从数据库中获取JDId
func (d *GoodJingDong) GetJDId() string {
	var good GoodJingDong
	err := global.DBEngine.Where("brand = ?", d.Brand).First(&good).Error
	if err != nil {
		return ""
	}
	return good.JDId
}
