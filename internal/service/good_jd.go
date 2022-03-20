package service

import (
	"Opendulum/internal/model"
	"Opendulum/internal/serializer"
)

type CreateJDGoodRequest struct {
	Brand string `json:"brand" form:"brand" binding:"required,min=1,max=50"`
	JDId  string `json:"jd_id" form:"jd_id" binding:"required,min=1,max=50"`
}

type UpdateJDGoodRequest struct {
	ID    uint   `json:"id" form:"id" binding:"required"`
	Brand string `json:"brand" form:"brand" binding:"min=1,max=50"`
	JDId  string `json:"jd_id" form:"jd_id" binding:"min=1,max=50"`
}

func (svc *Service) CreateJDGoodService(param *CreateJDGoodRequest) serializer.Response {
	jdGood := model.GoodJingDong{
		Brand: param.Brand,
		JDId:  param.JDId,
	}

	err := svc.db.Create(&jdGood).Error
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "create JDGood success",
		Data: serializer.BuildGoodJingDong(jdGood),
	}
}

func (svc Service) UpdateJDGoodService(param *UpdateJDGoodRequest) serializer.Response {
	var good model.GoodJingDong
	err := svc.db.First(&good, param.ID).Error
	if err != nil {
		return serializer.Response{
			Code: 500,
			Msg:  err.Error(),
		}
	}
	if param.Brand != "" {
		good.Brand = param.Brand
	}
	if param.JDId != "" {
		good.JDId = param.JDId
	}
	svc.db.Save(&good)
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildGoodJingDong(good),
		Msg:  "update JDGood success",
	}
}
