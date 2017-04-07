package response

import "github.com/zhyq132/cst/business"

// BusinessResponseMessage 业务消息对应的json格式
type BusinessResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		BusinessPur   int `json:"business.purchase"`
		BusinessDrive int `json:"business.drive"`
		BusinessMain  int `json:"business.maintenance"`
		BusinessIns   int `json:"business.insurance" `
		BusinessEme   int `json:"business.emergency"`
	} `json:"data"`
}

//ResponseBusiness 获取业务统计的相应结果
func ResponseBusiness(areaId int) BusinessResponseMessage {
	var responseMsg BusinessResponseMessage
	//查找当前area_id对应的业务信息
	responseMsg.Status = 200
	responseMsg.Message = "ok"
	responseMsg.Data.BusinessDrive = business.SellPromiseCount(areaId)
	responseMsg.Data.BusinessEme = business.SupportCount(areaId)
	responseMsg.Data.BusinessIns = business.XubaoCount(areaId)
	responseMsg.Data.BusinessMain = business.YangxiuCount(areaId)
	responseMsg.Data.BusinessPur = business.SellAskCount(areaId)

	return responseMsg
}
