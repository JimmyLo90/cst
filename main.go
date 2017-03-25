package main

import (
	"net/http"

	"log"

	"github.com/gorilla/websocket"
)

func handlerFunc() {

}

//pcPushMessage 推送的数据结构
type pcPushMessage struct {
	MsgCmd  string `json:"cmd"`
	MsgData struct {
		MsgAreaId string `json:"area_id"`
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
		log.Fatal("upgrade http request to ws err:", err)
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
			delete(clients, ws)
			ws.Close()
		} else {
			if msg.MsgCmd == "sign" {
				clients[ws] = msg
			} else if msg.MsgCmd == "push" {
				var responseMsg businessResponseMessage
				//查找当前area_id对应的业务信息
				responseMsg.Status = 200
				responseMsg.Message = "ok"
				responseMsg.Data.BusinessDrive = 0
				responseMsg.Data.BusinessEme = 0
				responseMsg.Data.BusinessIns = 0
				responseMsg.Data.BusinessMain = 0
				responseMsg.Data.BusinessPur = 0

				for key, val := range clients {
					if val.MsgData.MsgAreaId == msg.MsgData.MsgAreaId && val.MsgData.MsgType == msg.MsgData.MsgType {
						key.WriteJSON(responseMsg)
					}
				}
			}

		}
	}
}

//mian 程序主体
func main() {
	http.HandleFunc("/", handlerWs)

	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		panic("error" + err.Error())
	}
}
