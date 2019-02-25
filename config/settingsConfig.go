package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//settingsFile 环境设置文件名称
const settingsFile = "settings.yaml"

//Settings 环境设置结构体
type Settings struct {
	PcPushServiceSettings PcPushServiceSettingsDesc `yaml:"pc_push_service"`
}

// PcPushServiceSettingsDesc 环境设置细节
type PcPushServiceSettingsDesc struct {
	Address    string `yaml:"service_listen_ip"`
	Port       string `yaml:"service_listen_port"`
	SecureCert string `yaml:"service_listen_securt_cert"`
	SecureKey  string `yaml:"service_listen_securt_key"`
}

//GetPcPushServiceSettings 获取环境设置
func getSettings() Settings {
	settingsByte, err := ioutil.ReadFile(ConfigPath + settingsFile)
	if err != nil {
		fmt.Println(err)
	}

	var s Settings
	err = yaml.Unmarshal(settingsByte, &s)
	if err != nil {
		fmt.Println("unmarshal settings.yaml error:", err)
	}

	return s
}

//GetPcPushListendAddr 获取推送服务监听的地址
func GetPcPushListendAddr() string {
	s := getSettings()

	return s.PcPushServiceSettings.Address + ":" + s.PcPushServiceSettings.Port
}

//GetListendSecureCert 获取wss的cert
func GetListendSecureCert() string {
	s := getSettings()
	return s.PcPushServiceSettings.SecureCert
}

//GetListendSecureKey 获取wss的key
func GetListendSecureKey() string {
	s := getSettings()
	return s.PcPushServiceSettings.SecureKey
}
