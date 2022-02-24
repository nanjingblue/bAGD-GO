package service

import (
	"Opendulum/internal/model"
	"Opendulum/internal/serializer"
)

type CreateGoodRequest struct {
	Brand         string `json:"brand" form:"brand" binding:"required,min=1,max=50"`
	MapProduction string `json:"map_production" form:"map_production"`
	MapIngredient string `json:"map_ingredient" form:"map_ingredient"`
	MapTrends     string `json:"map_trends" form:"map_trends"`
	Energy        int32  `json:"energy" form:"energy"`
	Protein       int32  `json:"protein" form:"protein"`
	Fat           int32  `json:"fat" form:"fat"`
	Carbohydrates int32  `json:"carbohydrates" form:"carbohydrates"`
	Minerals      int32  `json:"minerals" form:"minerals"`
	Other         int32  `json:"other" form:"other"`
}

type GetGoodRequest struct {
	Brand string `form:"brand" json:"brand" binding:"required,min=1,max=50"`
}

func (svc *Service) CreateGoodService(param *CreateGoodRequest) serializer.Response {
	good := model.Good{
		Brand:         param.Brand,
		MapProduction: param.MapProduction,
		MapIngredient: param.MapIngredient,
		MapTrends:     param.MapTrends,
		Energy:        param.Energy,
		Protein:       param.Protein,
		Fat:           param.Fat,
		Carbohydrates: param.Carbohydrates,
		Minerals:      param.Minerals,
		Other:         param.Other,
	}
	err := svc.db.Create(&good).Error
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "internal server error",
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "create good success",
		Data: serializer.BuildGood(good),
	}
}

// GetGoodService 根据品牌名 获取商品信息
func (svc *Service) GetGoodService(param *GetGoodRequest) serializer.Response {
	var good model.Good
	err := svc.db.Where("brand = ?", param.Brand).First(&good).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "internal server error",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "success",
		Data: serializer.BuildGood(good),
	}
}
