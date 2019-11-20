package models

// 管理员数据结构
type Admin struct {
	Id        string `json:"id" bson:"id"`
	Username  string `json:"username" bson:"username"`
	Logintype string `json:"logintype" bson:"logintype"`
	Loginpwd  string `json:"loginpwd" bson:"loginpwd"`
}
