package main

import (
	"github.com/gin-gonic/gin"
	"zhyq132/cst/controller"

	_ "zhyq132/cst/config"
)

//mian 程序主体
func main() {
	r := gin.Default()

	microApp:= &controller.MicroAppController{}
	r.GET("test", microApp.ActionTest)

	r.Run(":80")
}