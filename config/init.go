package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Config配置
var Config Settings

func init() {
	c := flag.String("conf", "./config/", "config files path")
	flag.Parse()
	cPath := *c

	readTimes := 0
readFile:
	f, err := ioutil.ReadFile(cPath + "settings.yaml")
	if err != nil {
		if readTimes >= 3 {
			panic("读取配置文件错误，启动失败")
		} else {
			readTimes ++
			goto readFile
		}
	}

	yaml.Unmarshal(f, &Config)
}

// WssConf 环境设置细节
type WssConf struct {
	Address    string `yaml:"service_listen_ip"`
	Port       string `yaml:"service_listen_port"`
	SecureCert string `yaml:"service_listen_securt_cert"`
	SecureKey  string `yaml:"service_listen_securt_key"`
}

// WxOpenConf 开放平台配置明细
type WxOpenConf struct {
	WxOpenAppID          string `yaml:"WX_OPEN_APPID"`
	WxOpenSecret         string `yaml:"WX_OPEN_SECRET"`
	WxOpenReceiveToken   string `yaml:"WX_OPEN_RECEIVE_TOKEN"`
	WxOpenEncodingAeskey string `yaml:"WX_OPEN_ENCODING_AESKEY"`
	WxOpenAppUrl         string `yaml:"WX_OPEN_APP_URL"`
}

//Settings 环境设置结构体
type Settings struct {
	Wss    WssConf    `yaml:"wss_listen_service"`
	WxOpen WxOpenConf `yaml:"wx_open_platform"`
}

//ListendAddr 获取推送服务监听的地址
func (w WssConf) ListendAddr() string {
	return w.Address + ":" + w.Port
}
