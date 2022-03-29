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
func chat(context *gin.Context){
	var err error
	uid,err := strconv.ParseInt(context.Query("id"),10,64)
	token,_ := context.Get("token")

	if token != "qwe"{
		panic("token不对")
	}

	webSocket := websocket.Upgrader{ //将http升级为websocket
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	userConn := &UserNode{}

	userConn.Conn,err = webSocket.Upgrade(context.Writer,context.Request,nil)
	if err != nil{
		panic(err.Error())
	}

	UserConnMap[uid] = userConn
	//绑定uid和链接的关系

	//从io链接里读取消息 存入数据库 对方是否在线 在线需要存入对方内容里 然后发送给这条消息给对方

	go sendMessage(userConn)

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

		if receiveNode,isset := UserConnMap[mesObj.ReceiveUid]; isset{
			receiveNode.Message <- message
		}
		//这里还要把这个消息返回给客户端
	}
}

func receiveMessage(node *UserNode)  {

	for {
		select {
			case v,_ := <- node.Message :
				fmt.Println(v)
		}
	}
}

