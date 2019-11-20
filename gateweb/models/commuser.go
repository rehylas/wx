package models

import (
	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

func init() {

}

type CommUser struct {
	Userid    string `json:"userid"`
	Userpwd   string `json:"userpwd"`
	Openid    string `json:"openid"`
	State     int64  `json:"state"`
	Userlevel int64  `json:"userlevel"`
	Daymsgmax int64  `json:"daymsgmax"`
	Insertdt  string `json:"insertdt"`
}

func GetAllCommUsers(result *[]CommUser) error {

	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		return err
	}
	whereSql := bson.M{"enable": 1}
	err = collect.Find(whereSql).All(result)
	beego.Debug("recs:", result)
	return err

}

func (user *CommUser) GetCommUserById(userid string) error {
	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		return err
	}
	whereSql := bson.M{"userid": userid}
	err = collect.Find(whereSql).One(user)
	beego.Debug("user:", user)
	return err
}

func (user *CommUser) GetCommUserByOpenId(openid string) error {
	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		return err
	}
	whereSql := bson.M{"openid": openid}
	err = collect.Find(whereSql).One(user)
	beego.Debug("recs:", user)
	return err
}

func (user *CommUser) AddCommUser() error {
	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		return err
	}
	err = collect.Insert(user)
	return err
}
