package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github/yl/imchat/controller"
	"net/http"
)



type ConnUser struct{
	Conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func checkOrigin(r *http.Request) bool {
	return true
}

func main() {
	router := gin.Default()

	router.POST("/register",controller.Register)
	router.POST("/login",controller.Login)
	//router.POST("/hello",hello)
	//http.HandleFunc("/register",register)


	//db = model.Instance()
	//
	//rows,err := db.Query("select * from aa")
	//fmt.Println(rows,err)

	//这边要登录之后才长链接
	//如果有token 并且ok 才进入长链接
	router.Run(":8000")
}







//func view(){
//	view,err := template.ParseGlob("../view/**/*") // ** 目录 *文件
//	if err != nil{
//		panic(err)
//	}
//
//	for _,v := range view.Templates(){
//		viewName := v.Name()
//		fmt.Println(viewName)
//		http.HandleFunc(viewName, func(writer http.ResponseWriter, request *http.Request) {
//			err = view.ExecuteTemplate(writer,viewName,nil)
//			if err != nil{
//				panic(err)
//			}
//		})
//	}
//}





//写一个 im golang