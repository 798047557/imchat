package service

import (
	"github/yl/imchat/model"
	"time"
)

func Register(mobile int64,pwd,nickname,avatar string,sex int8) (){
	user := &model.User{}
	user.FindInfoByMobile(mobile)
	if user.Id > 0 {
		panic("已经被注册")
	}

	user.Mobile = mobile
	user.Password = model.Md5Password(pwd)
	user.Nickname = nickname
	user.Avatar = avatar
	user.Sex = sex
	user.CreatedTime = time.Now().Unix()
	_,err := user.CreateInfo()
	if err != nil {
		panic(err)
	}
}