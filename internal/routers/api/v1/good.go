package v1

import (
	"Opendulum/internal/service"
	"Opendulum/pkg/convert"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func Create(ctx *gin.Context) {
	param := service.CreateGoodRequest{}
	svc := service.New(ctx.Request.Context())
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.CreateGoodService(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(500, gin.H{
			"msg":   "fail",
			"error": err,
		})
	}
}

func GetGood(ctx *gin.Context) {
	param := service.GetGoodRequest{
		Brand: convert.StrTo(ctx.Param("brand")).String(),
	}
	svc := service.New(ctx.Request.Context())
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.GetGoodService(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(500, gin.H{
			"msg":   "fail",
			"error": err,
		})
	}
}

func Update(ctx *gin.Context) {
	param := service.UpdateGoodRequest{}
	svc := service.New(ctx.Request.Context())
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.UpdateGoodService(&param, ctx.Param("id"))
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "fail",
			"error": err,
		})
	}
}
