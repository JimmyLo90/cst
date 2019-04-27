package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"
	"zhyq132/cst/config"
	"zhyq132/cst/service/aimo"

	"github.com/gorilla/websocket"
)

type AimoController struct {
}

func (c *AimoController) WssClient(ctx *gin.Context) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: config.Config.Aimo.WssHost, Path: "/snap/notification"}
	log.Printf("connecting to %s", u.String())

	requestHeader := http.Header{}
	requestHeader.Add("Authorization", "APPCODE "+config.Config.Aimo.AppCode)
	wssClient, _, err := websocket.DefaultDialer.Dial(u.String(), requestHeader)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer wssClient.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := wssClient.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			//验证websockets的数据
			go func(msg []byte) {
				m := aimo.WssMessage{}
				id, err := m.Stored(msg)
				if err != nil {
					log.Printf("stored message err:%s,%v \n", err, id)
				} else {
					log.Printf("新入库aimo_pic数据，ID: %v", id)
				}

				//@todo 通知aimo智能客流项目组，有新的照片数据进来
				{
					body := url.Values{}
					body.Add("id", strconv.FormatInt(int64(id), 10))
					if _, err := http.PostForm(config.Config.Aimo.NextServiceHost, body); err != nil {
						log.Printf("stored message err:%s,%v \n", err, id)
					}
				}
			}(message)
		}
	}()

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := wssClient.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Println("write:", err)
				return
			} else {
				log.Println("write pingMessage to wss")
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			wssClient.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			wssClient.Close()
			close(done)
		}
	}
}
