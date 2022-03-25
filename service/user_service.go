package service

import (
	"github/yl/imchat/model"
	"time"
)

func Register(mobile int64,pwd,nickname,avatar string,sex int8) bool{
	user := &model.User{}
	//fmt.Printf("%v",user)
	user.FindInfoByMobile(mobile)
	if user.Id > 0 {
		panic("该手机已经被注册")
	}

	user.Mobile = mobile
	user.Password = model.Md5Password(pwd)
	user.Nickname = nickname
	user.Avatar = avatar
	user.Sex = sex
	user.CreateTime = time.Now().Unix()
	_,err := user.CreateInfo()
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	return true
}