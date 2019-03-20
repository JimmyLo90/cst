package tongji

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(TjMpNewsSend))
}
