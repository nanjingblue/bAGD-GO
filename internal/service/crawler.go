package service

import (
	"Opendulum/internal/serializer"
	"Opendulum/pkg/crawler"
	"github.com/gin-gonic/gin"
)

type GetCommentsRequest struct {
	ItemId string `form:"item_id" json:"item_id" binding:"required,min=1,max=50"`
}

func (svc *Service) GetCommentsService(param *GetCommentsRequest) serializer.Response {
	jd := crawler.NewJingDongCommentsRes(param.ItemId)
	comments := jd.Fetch()
	if comments != nil {
		return serializer.Response{
			Code: 200,
			Data: comments,
			Msg:  "get json ok",
		}
	}
	return serializer.Response{
		Code: 500,
		Data: gin.H{
			"url": jd.Url,
		},
		Msg: "get json fail",
	}
}
