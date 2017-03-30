package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zhyq132/cst/business"
	"github.com/zhyq132/cst/config"
)

//ConfigPath 配置路径，全局变量
var ConfigPath string

//pcPushMessage 推送的数据结构
type pcPushMessage struct {
	MsgCmd  string `json:"cmd"`
	MsgData struct {
		MsgAreaId int    `json:"area_id"`
		MsgType   string `json:"type"`
	} `json:"data"`
}

// businessResponseMessage 业务消息对应的json格式
type businessResponseMessage struct {
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

//clients map类型，存储所有ws链接
var clients = make(map[*websocket.Conn]pcPushMessage)

//upgrade http to ws
var upgrade = websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

//handlerWs websocket处理主体
func handlerWs(res http.ResponseWriter, req *http.Request) {
	ws, err := upgrade.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("upgrade http request to ws err:", err)
	}
	defer func() {
		delete(clients, ws)
		ws.Close()
	}()

	for {
		var msg pcPushMessage

		//接受ws传来的数据，不符合结构的，删除客户端,并关闭链接
		err := ws.ReadJSON(&msg)
		if err != nil {
			isWsClose := websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived)
			if isWsClose {
				delete(clients, ws)
				fmt.Println("*******************链接关闭,推送服务已停止\n\n\n\n ")
				break
			}
		} else {
			if msg.MsgCmd == "sign" {
				clients[ws] = msg
				fmt.Println("*******************身份标识确认,已接入推送服务\n\n\n\n ")
			} else if msg.MsgCmd == "push" {
				var responseMsg businessResponseMessage
				//查找当前area_id对应的业务信息
				responseMsg.Status = 200
				responseMsg.Message = "ok"
				responseMsg.Data.BusinessDrive = business.SellPromiseCount(msg.MsgData.MsgAreaId)
				responseMsg.Data.BusinessEme = business.SupportCount(msg.MsgData.MsgAreaId)
				responseMsg.Data.BusinessIns = business.XubaoCount(msg.MsgData.MsgAreaId)
				responseMsg.Data.BusinessMain = business.YangxiuCount(msg.MsgData.MsgAreaId)
				responseMsg.Data.BusinessPur = business.SellAskCount(msg.MsgData.MsgAreaId)

				for key, val := range clients {
					if val.MsgData.MsgAreaId == msg.MsgData.MsgAreaId && val.MsgData.MsgType == msg.MsgData.MsgType {
						key.WriteJSON(responseMsg)
						fmt.Println("*******************数据获取完毕确认,已推送\n\n\n\n ")
					}
				}
			}
		}
	}
}

//mian 程序主体
func main() {
	http.HandleFunc("/", handlerWs)

	if err := http.ListenAndServe(config.GetPcPushListendAddr(), nil); err != nil {
		panic("error" + err.Error())
	}
}
