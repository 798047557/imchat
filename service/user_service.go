package service

import (
	"fmt"
	"github/yl/imchat/lib"
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
	user.Password = lib.Md5Password(pwd)
	//fmt.Println(user.Password)
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

func Login(mobile int64,pwd string) string{
	user := &model.User{}
	//fmt.Printf("%v",user)
	user.FindInfoByMobile(mobile)

	if user.Id == 0{
		panic("账户不存在")
	}

	fmt.Println(user.Password,lib.Md5Password(pwd))
	if user.Password != lib.Md5Password(pwd) {
		panic("密码不正确")
	}

	token := lib.CreateToken()
	user.Token = token
	user.UpdateUserByUid()

	return token
}

func checkToken(uid int64,token string) bool {
	return true
}