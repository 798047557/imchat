package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"main"
)

func chat(context *gin.Context){
	//query := context.URL.Query()
	//id := query.Get("id")
	//token := query.Get("token")


	webSocket := websocket.Upgrader{ //将http升级为websocket
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	userConn := main.

	conn,err := webSocket.Upgrade(context.Writer,context.Request,nil)

	if err != nil{
		panic(err.Error())
	}

	fmt.Println(conn.ReadMessage())


}

