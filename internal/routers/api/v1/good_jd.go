package v1

import (
	"Opendulum/internal/service"
	"Opendulum/pkg/convert"
	"github.com/gin-gonic/gin"
)

func CreateJD(ctx *gin.Context) {
	param := service.CreateJDGoodRequest{
		Brand: convert.StrTo(ctx.Param("brand")).String(),
		JDId:  convert.StrTo(ctx.Param("jd_id")).String(),
	}
	svc := service.New(ctx.Request.Context())
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.CreateJDGoodService(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "fail",
			"error": err.Error(),
		})
	}
}

func UPdateJD(ctx *gin.Context) {
	param := service.UpdateJDGoodRequest{}
	svc := service.New(ctx.Request.Context())
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.UpdateJDGoodService(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "fail",
			"error": err.Error(),
		})
	}
}
