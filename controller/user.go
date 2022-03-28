package controller

import (
	"github.com/gin-gonic/gin"
	"github/yl/imchat/service"
	"strconv"
)


const SUCCESS_CODE = 200
const FAIL_CODE = 500

func throw(context *gin.Context){
	if r := recover();r != nil{
		context.JSON(FAIL_CODE,gin.H{
			"code":FAIL_CODE,
			"msg":r,
		})
	}
}

func Register(context *gin.Context) {
	defer throw(context)
	mobile,err := strconv.ParseInt(context.PostForm("mobile"),10,64)
	if err != nil{
		panic(err)
	}

	password := context.PostForm("password")
	if password == ""{
		panic("密码不能为空")
	}
	service.Register(mobile,password,"nickname","avatar",1)
	context.JSON(SUCCESS_CODE,gin.H{
		"code":SUCCESS_CODE,
		"msg":"注册成功",
	})
}


func Login(context *gin.Context) {
	defer throw(context)
	mobile,err := strconv.ParseInt(context.PostForm("mobile"),10,64)
	password := context.PostForm("password")
	if err != nil{
		panic(err)
	}

	token := service.Login(mobile,password)
	context.JSON(SUCCESS_CODE,gin.H{
		"code":SUCCESS_CODE,
		"msg":"登录成功",
		"data":token,
	})
}