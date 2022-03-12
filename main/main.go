package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github/yl/imchat/service"
	"net/http"
	"strconv"
	"text/template"
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
	defer func() {
		if r := recover();r != nil{
			fmt.Println(r)
		}
	}()

	router := gin.Default()

	router.GET("/register",register)
	//http.HandleFunc("/register",register)


	//db = model.Instance()
	//
	//rows,err := db.Query("select * from aa")
	//fmt.Println(rows,err)

	//这边要登录之后才长链接
	//如果有token 并且ok 才进入长链接
	router.Run(":8000")
}

func register(context *gin.Context) {
	mobile,err := strconv.ParseInt(context.Param("mobile"),10,64)
	if err != nil{
		panic("手机号码参数错误")
	}
	password := context.Param("password")
	service.Register(mobile,password,"nickname","avatar",1)

	context.JSON(200,gin.H{
		"code":200,
		"success":true,
	})
}


func view(){
	view,err := template.ParseGlob("../view/**/*") // ** 目录 *文件
	if err != nil{
		panic(err)
	}

	for _,v := range view.Templates(){
		viewName := v.Name()
		fmt.Println(viewName)
		http.HandleFunc(viewName, func(writer http.ResponseWriter, request *http.Request) {
			err = view.ExecuteTemplate(writer,viewName,nil)
			if err != nil{
				panic(err)
			}
		})
	}
}





//写一个 im golang