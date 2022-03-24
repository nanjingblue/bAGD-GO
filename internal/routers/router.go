package routers

import (
	_ "Opendulum/docs"
	v1 "Opendulum/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Pong)
		apiv1.POST("/good", v1.CreateGood)
		apiv1.GET("/good/:brand", v1.GetGood)
		apiv1.PUT("/good/:id", v1.UpdateGood)
		apiv1.DELETE("/good/:brand", v1.DeleteGood)
		apiv1.GET("/good/comments/:brand", v1.GetGoodComments)

		apiv1.GET("/comments/:product_id", v1.GetComments)

		apiv1.POST("/jdgood", v1.CreateJD)
		apiv1.PUT("/jdgood", v1.UPdateJD)
	}
	return r
}
