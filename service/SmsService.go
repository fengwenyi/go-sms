package service

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"go-sms/util"
	"log"
)

type SmsService struct {
}

func (s *SmsService) SendCode(phone string, code string) bool {
	smsConfig := util.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(smsConfig.RegionId, smsConfig.AppKey, smsConfig.AppSecret)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = smsConfig.SignName
	request.TemplateCode = smsConfig.TemplateCode
	request.PhoneNumbers = phone

	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})

	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	// 3.判断发送结果：接收返回结果，并判断发送状态
	return response.Code == "OK"
}
