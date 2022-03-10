package main

import (
	"Opendulum/global"
	"Opendulum/internal/model"
	"Opendulum/internal/routers"
	"Opendulum/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	err := setupSetting()
	if err != nil {
		panic(err)
	}
	err = model.SetupDBEngine()
	if err != nil {
		panic(err)
	}
}

// @title 基于AR的商品信息展示系统
// @version 1.0
// @description 基于AR的商品信息展示系统
// @termsOfService https://github.com/0pendulum/bAGD-GO
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:        ":" + global.ServerSetting.HttpPort,
		Handler:     router,
		ReadTimeout: global.ServerSetting.ReadTimeOut,
		//WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}

func setupSetting() interface{} {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
