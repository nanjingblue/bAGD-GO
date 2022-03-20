package v1

import (
	"Opendulum/internal/service"
	"Opendulum/pkg/convert"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{})
}

// @Summary 新增商品
// @Produce json
// @Param brand body string true "品牌名称" minlength(1) maxlenght(50)
// @Param map_production body string false "产地分布图" minlength(1) maxlenght(100)
// @Param map_Ingredient body string false "配料比例图"
// @Param map_trends body string false "品牌销量趋势图"
// @Param energy body int false "能量值"
// @Param protein body int false "蛋白质"
// @Param fat body int false "脂肪"
//@Param carbohydrates body int false "碳水化合物"
// @param minerals body int false "矿物质"
// @Param other body int false "其他"
// @Success 200 {object} model.Good "成功"
// @Failure 400 "请求错误"
// @Router /api/v1/good [post]
func CreateGood(ctx *gin.Context) {
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

func UpdateGood(ctx *gin.Context) {
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

func DeleteGood(ctx *gin.Context) {
	param := service.DeleteGoodRequest{
		Brand: convert.StrTo(ctx.Param("brand")).String(),
	}
	svc := service.New(ctx.Request.Context())
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.DeleteGoodService(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "fail",
			"error": err,
		})
	}
}
