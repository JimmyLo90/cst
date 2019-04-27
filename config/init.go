package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"zhyq132/wechat7/base"
)

//Config配置
var Config settings

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
type wssConf struct {
	Address    string `yaml:"service_listen_ip"`
	Port       string `yaml:"service_listen_port"`
	SecureCert string `yaml:"service_listen_securt_cert"`
	SecureKey  string `yaml:"service_listen_securt_key"`
}

//ListendAddr 获取推送服务监听的地址
func (w wssConf) ListendAddr() string {
	return w.Address + ":" + w.Port
}

type httpConf struct {
	Port string `yaml:"service_listen_port"`
}

type cstMysqlDB struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (db *cstMysqlDB) Dsn() string {
	return db.Username + ":" + db.Password + "@tcp(" + db.IP + ":" + db.Port + ")/4s_wx_db"
}

type cstMongoDB struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (db *cstMongoDB) Dsn() string {
	return "mongodb://" + db.IP + ":" + db.Port + "/"
}

type aimo struct {
	WssHost          string `yaml:"wss_host"`
	AppCode          string `yaml:"app_code"`
	PicSavePath      string `yaml:"pic_save_path"`
	PicSavePathForDB string `yaml:"pic_save_path_for_db"`
	NextServiceHost  string `yaml:"next_service_host"`
}

//settings 环境设置结构体
type settings struct {
	Wss        wssConf         `yaml:"wss_listen_service"`
	Http       httpConf        `yaml:"http_listen_service"`
	WxOpen     base.WxOpenConf `yaml:"wx_open_platform"`
	CstMysqlDB cstMysqlDB      `yaml:"cst_mysql_db"`
	CstMongoDB cstMongoDB      `yaml:"cst_mongo_db"`
	Aimo       aimo            `yaml:"aimo"`
}
