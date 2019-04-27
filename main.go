package main

import (
	"github.com/gin-gonic/gin"
	"zhyq132/cst/config"
	"zhyq132/cst/controller"

	_ "zhyq132/cst/config"
)

//mian 程序主体
func main() {
	r := gin.Default()

	//microApp := &controller.MicroAppController{}
	//r.GET("release-app", microApp.ActionReleaseApp)
	//
	//mpNews := &controller.MpNewsController{}
	//r.GET("get-url-by-id", mpNews.ActionGetUrlByID)

	aimo := &controller.AimoController{}
	r.GET("aimo/wss-client", aimo.WssClient)

	r.Run(":" + config.Config.Http.Port)
}
