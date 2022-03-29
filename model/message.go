package model

type Message struct{
	Id int64 `json:"id,omitempty"`
	SendUid int64 `json:"send_uid,omitempty"`
	ReceiveUid int64 `json:"receive_uid,omitempty"` //uid 或者 群id
	MessageType int8 `json:"message_type,omitempty"` //1纯文字 2图片 3文字图片 4红包
	Content string `json:"content,omitempty"` //消息内容
	Scope int8 `json:"scope,omitempty"` //消息范围 1个人 2 群聊
	CreatedTime int `json:"created_time,omitempty"` //创建时间
	IsRevoke int8 `json:"is_revoke,omitempty"` //是否撤回
}

func (m *Message) Send() (int64,error)  {
	return db.InsertOne(m)
}