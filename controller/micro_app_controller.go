package controller

import (
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"
	"time"
	"zhyq132/cst/config"
	"zhyq132/cst/model/mongo"
	"zhyq132/wechat7"

	"zhyq132/cst/model/center"
)

type MicroAppController struct {
}

func (c *MicroAppController) ActionTest(ginContext *gin.Context) {
	o := orm.NewOrm()

	//查询待发布的小程序
	o.Using((center.WxAppSubmitAudit{}).DBName())
	q := o.QueryTable((center.WxAppSubmitAudit{}).TableName())
	var holdReleaseList []*center.WxAppSubmitAudit
	//where条件 ：审核中的小程序的where条件
	w := orm.NewCondition()
	w.And("status", 0).And("dateline__lte", time.Now().Unix()-7200)
	q.SetCond(w).All(&holdReleaseList)

	var list = make(map[string]interface{})
	list["nowtime"] = time.Now().Unix()

	//获取mongo链接
	s := MongoSession.Copy()
	defer s.Close()

	//第三方平台的基础信息
	openPlatform := &mongo.WeixinOpenPlatform{}
	openCollec := s.DB(openPlatform.DBName()).C(openPlatform.TableName())
	openQ := bson.D{
		{"a_areaId", 0},
	}
	openCollec.Find(openQ).One(openPlatform)

	//实例化wechat7应用
	wechat7App := &wechat7.App{}
	wechat7App.SetBaseConf(config.Config.WxOpen)

	for _, v := range holdReleaseList {
		list[v.AuditID] = v
		//读取微信app授权信息
		wxOpenAppModel := &mongo.WeixinOpenApp{}
		c := s.DB(wxOpenAppModel.DBName()).C(wxOpenAppModel.TableName())
		q := bson.D{
			{"a_areaId", v.AreaID},
			{"oauth_status", 1},
		}

		c.Find(q).One(&wxOpenAppModel)
		if wxOpenAppModel.AreaId > 0 {
			//请求接口获取app的审核状态
			 wechat7App.AuditStatus(wxOpenAppModel,v.AuditID)
		}
	}

	ginContext.JSON(200, gin.H(list))
}
