package center

import "github.com/astaxie/beego/orm"

func init(){
	orm.RegisterModel(new(WxAppSubmitAudit))
	orm.RegisterModel(new(WxArea))
}
