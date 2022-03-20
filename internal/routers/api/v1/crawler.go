package v1

import (
	"Opendulum/internal/service"
	"Opendulum/pkg/convert"
	"github.com/gin-gonic/gin"
)

func GetComments(ctx *gin.Context) {
	param := service.GetCommentsRequest{
		ItemId: convert.StrTo(ctx.Param("product_id")).String(),
	}
	svc := service.New(ctx.Request.Context())
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.GetCommentsService(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(200, gin.H{
			"msg":   "fail",
			"error": err,
		})
	}
}
