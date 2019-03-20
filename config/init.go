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

//ListendAddr 获取推送服务监听的地址
func (w WssConf) ListendAddr() string {
	return w.Address + ":" + w.Port
}

type HttpConf struct {
	Port string `yaml:"service_listen_port"`
}

// WxOpenConf 开放平台配置明细
type WxOpenConf struct {
	WxOpenAppID          string `yaml:"WX_OPEN_APPID"`
	WxOpenSecret         string `yaml:"WX_OPEN_SECRET"`
	WxOpenReceiveToken   string `yaml:"WX_OPEN_RECEIVE_TOKEN"`
	WxOpenEncodingAeskey string `yaml:"WX_OPEN_ENCODING_AESKEY"`
	WxOpenAppUrl         string `yaml:"WX_OPEN_APP_URL"`
}

type CstMysqlDB struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (db *CstMysqlDB) Dsn() string {
	return db.Username + ":" + db.Password + "@tcp(" + db.IP + ":" + db.Port + ")/w_center"
}

type CstMongoDB struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (db *CstMongoDB) Dsn() string {
	return "mongodb://" + db.IP + ":" + db.Port + "/"
}

//Settings 环境设置结构体
type Settings struct {
	Wss        WssConf    `yaml:"wss_listen_service"`
	Http       HttpConf   `yaml:"http_listen_service"`
	WxOpen     WxOpenConf `yaml:"wx_open_platform"`
	CstMysqlDB CstMysqlDB `yaml:"cst_mysql_db"`
	CstMongoDB CstMongoDB `yaml:"cst_mongo_db"`
}
