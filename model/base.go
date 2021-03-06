package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const(
	host = "127.0.0.1"
	port = "3306"
	user = "root"
	dbname = "im"
	password = "123456"
)

var db *xorm.Engine

func init() () {
	var err error
	db,err = xorm.NewEngine("mysql",user+":"+password+"@("+host+":"+port+")/"+dbname+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	//db.ShowSQL(true)//打印sql
	db.SetMaxOpenConns(10)//最大连接数
}



