package main

import (
	"github.com/gin-gonic/gin"
	"go-sms/controller"
	"go-sms/util"
)

func main() {

	util.ParseConfig("./config/app.json")

	r := gin.Default()

	Routers(r)

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func Routers(r *gin.Engine) {
	new(controller.SmsController).Router(r)
}
