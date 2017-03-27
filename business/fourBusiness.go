package business

import (
	"fmt"
	"io/ioutil"
)

//DBConfig 配置文件路径
const DBConfig = "/q/project/Go/src/github.com/zhyq132/cst/config/db.yaml"

//SellAskCount 购车询价统计
func SellAskCount() {
	db, err := ioutil.ReadFile(DBConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
}
