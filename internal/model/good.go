package model

import "github.com/jinzhu/gorm"

// Good 商品具体信息
type Good struct {
	gorm.Model
	Brand         string `gorm:"type:varchar(50);unique"`
	MapProduction string // 产地图
	MapIngredient string // 配料表
	MapTrends     string // 品牌趋势
	Energy        int32  // 能量
	Protein       int32  // 蛋白质
	Fat           int32  // 脂肪
	Carbohydrates int32  // 碳水化合物
	Minerals      int32  // 矿物质
	Other         int32  // 其他
	VideoURL      string // 商品广告视频
}
