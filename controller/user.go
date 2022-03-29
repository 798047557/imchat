package controller

import (
	"github.com/gin-gonic/gin"
	"github/yl/imchat/lib"
	"github/yl/imchat/service"
	"strconv"
)





func Register(context *gin.Context) {
	defer lib.Throw(context)
	mobile,err := strconv.ParseInt(context.PostForm("mobile"),10,64)
	if err != nil{
		panic(err)
	}

	password := context.PostForm("password")
	if password == ""{
		panic("密码不能为空")
	}
	service.Register(mobile,password,"nickname","avatar",1)
	context.JSON(lib.SUCCESS_CODE,gin.H{
		"code":lib.SUCCESS_CODE,
		"msg":"注册成功",
	})
}


func Login(context *gin.Context) {
	defer lib.Throw(context)
	mobile,err := strconv.ParseInt(context.PostForm("mobile"),10,64)
	password := context.PostForm("password")
	if err != nil{
		panic(err)
	}

	token := service.Login(mobile,password)
	context.JSON(lib.SUCCESS_CODE,gin.H{
		"code":lib.SUCCESS_CODE,
		"msg":"登录成功",
		"data":token,
	})
}