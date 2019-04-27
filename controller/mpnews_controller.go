package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type MpNewsController struct {
}

func (c *MpNewsController) ActionGetUrlByID(ginContext *gin.Context) {
	//var r map[string]interface{}
	//r = make(map[string]interface{})
	//r["code"] = 200
	//r["msg"] = ""
	//data := make(map[string]interface{})
	//defer func() {
	//	if p := recover(); p != nil {
	//		r["code"] = 500
	//		str, ok := p.(string)
	//		if ok {
	//			r["msg"] = str
	//		} else {
	//			r["msg"] = "抓取url失败"
	//		}
	//	}
	//	r["data"] = data
	//	ginContext.JSON(200, r)
	//}()
	//
	//mpNewsID := ginContext.Request.URL.Query().Get("id")
	//if mpNewsID == "" {
	//	panic("获取ID失败")
	//}
	//
	//o := orm.NewOrm()
	//o.Using((tongji.TjMpNewsSend{}).DBName())
	//
	////查询tj_mpnews_send
	//sql := ` SELECT * FROM tongji.tj_mpnews_send WHERE id =? `
	//
	//var one *tongji.TjMpNewsSend
	//err := o.Raw(sql, mpNewsID).QueryRow(&one)
	//if err != nil {
	//	panic("获取mpNews数据失败")
	//}
	//
	//var area *center.WxArea
	//sql = ` SELECT * FROM w_center.wx_area WHERE id =? `
	//err = o.Raw(sql, one.AreaID).QueryRow(&area)
	//if err != nil {
	//	panic("获取店铺信息失败")
	//} else if area.WeixinType != 4 {
	//	panic("目前仅支持认证订阅号抓取图文url")
	//}
	//
	////爬取对应的url
	//url := crawl.SougouMpnewsUrl(one.Title)
	//if url != "" {
	//	//更新tj_mpnews_send中的url
	//	o.Raw(" UPDATE tongji.tj_mpnews_send SET url=? WHERE area_id=? AND msgid=? ;", url, one.AreaID, one.MsgID).Exec()
	//} else {
	//	panic("目前仅支持认证订阅号抓取图文url")
	//}
	//
	//data["url"] = url
}

func (c *MpNewsController) ActionMpNews(ginContext *gin.Context) {
//	o := orm.NewOrm()
//
//	type SendNewsOrder struct {
//		ID          int32  `orm:"column(id)"`
//		AreaID      int32  `orm:"column(area_id)"`
//		Title       string `orm:"column(title)"`
//		MsgID       string `orm:"column(msgid)"`
//		SumReadUser int32  `orm:"column(c)"`
//	}
//
//	//查询阅读量排行最高的数据
//	sql := `
//SELECT
//  result.id,
//  result.area_id,
//  result.title,
//  result.msgid,
//  result.c
//FROM
//  (SELECT
//    *,
//    SUM(int_page_read_user) AS c
//  FROM
//    (SELECT
//      *
//    FROM
//      tongji.tj_mpnews_send
//    WHERE dateline >= (UNIX_TIMESTAMP() - (86400 * 7))
//      AND dateline <= UNIX_TIMESTAMP()
//    ORDER BY id DESC) AS b
//  GROUP BY b.area_id,
//    b.msgid
//  ORDER BY c DESC) AS result
//  LEFT JOIN w_center.wx_area
//    ON result.area_id = wx_area.id
//WHERE wx_area.weixinType = 4
//ORDER BY result.c DESC;
//	`
//
//	o.Using((tongji.TjMpNewsSend{}).DBName())
//	var list []*SendNewsOrder
//	o.Raw(sql).QueryRows(&list)
//
//	for _, v := range list {
//		//休息3秒
//		time.Sleep(time.Duration(time.Second * 3))
//
//		url := crawl.SougouMpnewsUrl(v.Title)
//		if url != "" {
//			//更新tj_mpnews_send中的url
//			o.Raw(" UPDATE tongji.tj_mpnews_send SET url=? WHERE area_id=? AND msgid=? ;", url, v.AreaID, v.MsgID).Exec()
//			fmt.Println(url, "mpnews_send ID=>", v.ID, "------------------------------\n")
//		}
//	}
}
