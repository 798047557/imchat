package model
type User struct{
	Id int64 `xorm:" autoincr int(10)" json:"id"`
	Token string `xorm:" varchar(40)" json:"token"`
	Nickname string `xorm:" varchar(10)" json:"nickname"`
	Avatar string `xorm:" varchar(150)" json:"avatar"`
	Sex int8 `xorm:" tinyint(1)" json:"sex"`
	Mobile int64 `xorm:" bigint(20)" json:"mobile"`
	Password string `xorm:" varchar(20)" json:"password"`
	Online string `xorm:" tinyint(1)" json:"online"`
	CreatedTime int64 `xorm:" bigint(20)" json:"created_time"`
}


func (u *User) CreateInfo() (int64,error){
	//创建user模型
	return db.InsertOne(u)
}

func (u *User) FindInfoByMobile(mobile int64){
	_,err := db.Where("mobile = ? ",mobile).Get(u)
	if err != nil{
		panic(u)
	}
}

func UpdateUserByUid(u *User){

}