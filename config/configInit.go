package config

import "flag"

//ConfigPath 配置文件路径
var ConfigPath string

func init() {
	conf := flag.String("conf", "./", "config files path")
	flag.Parse()
	ConfigPath = *conf
}
