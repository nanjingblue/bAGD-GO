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
		apiv1.POST("/good", v1.Create)
		apiv1.GET("/good/:brand", v1.GetGood)
		apiv1.PUT("/good/:id", v1.Update)
		apiv1.DELETE("/good/:brand", v1.Delete)
	}
	return r
}
