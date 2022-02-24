package routers

import (
	v1 "Opendulum/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Pong)
		apiv1.POST("/good", v1.Create)
		apiv1.GET("/good/:brand", v1.GetGood)
		apiv1.PUT("/good/:id", v1.Update)
	}
	return r
}
