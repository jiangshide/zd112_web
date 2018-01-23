package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/astaxie/beego"
	"time"
)

type WebSocketController struct {
	BaseController
}

type Message struct {
	Message string `json:"message"`
}

var (
	upgrader  = websocket.Upgrader{}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
)

func (this *WebSocketController) Get() {
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
	}
	clients[ws]=true
	for{
		time.Sleep(time.Second*3)
		msg := Message{Message:"这是向页面发送的数据:"+time.Now().Format("2006-01-02 15:04:05")}
		broadcast <- msg
	}
}

func init() {
	go handleMessages()
}

func handleMessages() {
	for {
		msg := <-broadcast
		beego.Info("client len: ", len(clients))
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				beego.Info("client.WriteJSON error:%v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
