package util

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	Sms SmsConfig `json:"sms"`
}

// 短信配置
type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
}

var _cfg *Config = nil

// 返回配置
func GetConfig() *Config {
	return _cfg
}

// 解析配置
func ParseConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&_cfg)
}
