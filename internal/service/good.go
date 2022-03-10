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
	VideoURL      string `json:"video_url" form:"video_url"`
}

type GetGoodRequest struct {
	Brand string `form:"brand" json:"brand" binding:"required,min=1,max=50"`
}

type UpdateGoodRequest struct {
	ID            uint
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
	VideoURL      string `json:"video_url" form:"video_url"`
}

type DeleteGoodRequest struct {
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
		VideoURL:      param.VideoURL,
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
func (svc *Service) GetGoodService(param *GetGoodRequest) serializer.Good {
	var good model.Good
	err := svc.db.Where("brand = ?", param.Brand).First(&good).Error
	if err != nil {
		panic(err)
	}
	return serializer.BuildGood(good)
}

// UpdateGoodService 根据商品 ID 更新good信息
func (svc *Service) UpdateGoodService(param *UpdateGoodRequest, id string) serializer.Response {
	var good model.Good
	err := svc.db.First(&good, id).Error
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "internal server error",
			Error: err.Error(),
		}
	}
	if param.Brand != "" {
		good.Brand = param.Brand
	}
	if param.MapProduction != "" {
		good.MapProduction = param.MapProduction
	}
	if param.MapIngredient != "" {
		good.MapIngredient = param.MapIngredient
	}
	if param.MapTrends != "" {
		good.MapTrends = param.MapTrends
	}
	if param.Energy != 0 {
		good.Energy = param.Energy
	}
	if param.Protein != 0 {
		good.Protein = param.Protein
	}
	if param.Fat != 0 {
		good.Fat = param.Fat
	}
	if param.Carbohydrates != 0 {
		good.Carbohydrates = param.Carbohydrates
	}
	if param.Minerals != 0 {
		good.Minerals = param.Minerals
	}
	if param.Other != 0 {
		good.Other = param.Other
	}
	if param.VideoURL != "" {
		good.VideoURL = param.VideoURL
	}
	svc.db.Save(&good)
	return serializer.Response{
		Code: 200,
		Msg:  "update good success",
		Data: serializer.BuildGood(good),
	}
}

// DeleteGoodService 根据商品品牌删除商品
func (svc *Service) DeleteGoodService(param *DeleteGoodRequest) serializer.Response {
	var good model.Good
	err := svc.db.Where("brand = ?", param.Brand).First(&good).Error
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "internal server error",
		}
	}
	svc.db.Delete(&good)
	return serializer.Response{
		Code: 200,
		Msg:  "delete good success",
		Data: serializer.BuildGood(good),
	}
}
