package controller

import (
	"github.com/astaxie/beego/orm"
	"gopkg.in/mgo.v2"
	"zhyq132/cst/config"
)

var MongoSession *mgo.Session

func init() {
	//初始化mysql
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", config.Config.CstMysqlDB.Dsn())

	//初始化mongo
	s, err := mgo.Dial(config.Config.CstMongoDB.Dsn())
	if err != nil {
		panic("mongo连接失败")
	}

	MongoSession = s.Copy()
}
