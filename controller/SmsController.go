package controller

import (
	"github.com/gin-gonic/gin"
	"go-sms/service"
	"log"
	"net/http"
)

type SmsController struct {
}

func (sc *SmsController) Router(e *gin.Engine) {
	e.GET("/sms/api/send-code", sc.SendCode)
}

func (sc *SmsController) SendCode(c *gin.Context) {

	phone, _ := c.GetQuery("phone")
	code, _ := c.GetQuery("code")

	s := service.SmsService{}
	result := s.SendCode(phone, code)
	log.Println("短信发送结果：", result)
	if result {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    0,
			"message": "验证码发送成功",
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": "验证码发送失败",
		})
	}
}
