package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github/yl/imchat/model"
	"net/http"
	"strconv"
)

type UserNode struct{
	Conn *websocket.Conn
	Message chan []byte //作为线上用户私聊的消息
}

var UserConnMap map[int64]*UserNode = make(map[int64]*UserNode) //uid 和 用户的映射关系

//链接上ws
func Chat(context *gin.Context){
	defer func() {
		if r := recover();r != nil{
			fmt.Println(r)
		}
	}()
	var err error
	uid,err := strconv.ParseInt(context.Query("id"),10,64)
	token := context.Query("token")
	//fmt.Println(token)
	if token != "qwe"{
		panic("token不对")
	}

	webSocket := websocket.Upgrader{ //将http升级为websocket
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	userConn := &UserNode{
		Message: make(chan []byte,50), //50句话的容量
	}

	userConn.Conn,err = webSocket.Upgrade(context.Writer,context.Request,nil)
	if err != nil{
		panic(err.Error())
	}

	UserConnMap[uid] = userConn
	//绑定uid和链接的关系

	//从io链接里读取消息 存入数据库 对方是否在线 在线需要存入对方内容里 然后发送给这条消息给对方

	//fmt.Println(userConn)
	go sendMessage(userConn)
	//
	go receiveMessage(userConn)

}

//这里是不停从前端接受消息 发送出去
func sendMessage(node *UserNode)  {
	for {
		_,message,err := node.Conn.ReadMessage()
		if err != nil{
			panic(err)
		}

		mesObj := &model.Message{}
		err = json.Unmarshal(message,mesObj) //这个msg只负责入库
		if err != nil{
			panic(err)
		}

		_,err = mesObj.Send()
		if err != nil{
			panic(err)
		}

		//fmt.Println(UserConnMap[mesObj.ReceiveUid])
		if receiveNode,isset := UserConnMap[mesObj.ReceiveUid]; isset{
			receiveNode.Message <- message
		}
	}
}

func receiveMessage(node *UserNode)  {
	for {
		select {
			case <- node.Message :
				node.Conn.WriteMessage(websocket.TextMessage,[]byte("收到了"))
		}
	}
}

