package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zhyq132/cst/config"
	"github.com/zhyq132/cst/response"
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	ws, err := upgrade.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("upgrade http request to ws err:", err)
	} else {

	}
	defer func() {
		if _, exist := clients[ws]; exist {
			delete(clients, ws)
			ws.Close()
		}
	}()

	for {
		var msg pcPushMessage

		//接受ws传来的数据，不符合结构的，删除客户端,并关闭链接
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			panic(err.Error() + "*******************链接断开，停止参与推送服务\n\n\n\n ")
		} else {
			resBusinessMsg := response.ResponseBusiness(msg.MsgData.MsgAreaId)

			if msg.MsgCmd == "sign" {
				clients[ws] = msg
				fmt.Println("*******************身份标识确认,已接入推送服务\n\n\n\n ")
				ws.WriteJSON(resBusinessMsg)
			} else if msg.MsgCmd == "push" {
				for key, val := range clients {
					if val.MsgData.MsgAreaId == msg.MsgData.MsgAreaId && val.MsgData.MsgType == msg.MsgData.MsgType {
						key.WriteJSON(resBusinessMsg)
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
