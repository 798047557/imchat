package lib

import "github.com/gin-gonic/gin"

const SUCCESS_CODE = 200
const FAIL_CODE = 500

func Throw(context *gin.Context){
	if r := recover();r != nil{
		context.JSON(FAIL_CODE,gin.H{
			"code":FAIL_CODE,
			"msg":r,
		})
	}
}
