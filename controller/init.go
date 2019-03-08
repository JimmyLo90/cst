package controller

import (
	"github.com/astaxie/beego/orm"
	"gopkg.in/mgo.v2"
)

var MongoSession *mgo.Session

func init() {
	//初始化mysql
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:allen123hjmHJM@tcp(192.168.2.240:5306)/w_center")

	//初始化mongo

	s, err := mgo.Dial("mongodb://192.168.2.240:11708/")
	if err != nil {
		panic("mongo连接失败")
	}

	MongoSession = s.Copy()
}
