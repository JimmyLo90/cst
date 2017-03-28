package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//DBConfigPath 配置文件路径
const DBConfigFile = "db.yaml"

//DbConfig 数据库配置结构体
type DbConfig struct {
	Db  DbDesc `yaml:"db"`
	Wdb DbDesc `yaml:"wdb"`
}

// DbDesc 数据库配置细节
type DbDesc struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

//getDbConfig 获取dbconfig
func getDbConfig() DbConfig {
	db, err := ioutil.ReadFile(ConfigPath + DBConfigFile)
	if err != nil {
		fmt.Println(err)
	}
	var dbConfig DbConfig
	err = yaml.Unmarshal(db, &dbConfig)

	return dbConfig
}

//GetDbDSN 后去db的数据库DSN
func GetDbDSN() string {
	var s string

	dbConf := getDbConfig()

	db := dbConf.Db

	// example   id:password@tcp(your-amazonaws-uri.com:3306)/dbname
	s = db.Username + ":" + db.Password + "@tcp(" + db.Address + ":" + db.Port + ")/" + db.Dbname

	return s
}
