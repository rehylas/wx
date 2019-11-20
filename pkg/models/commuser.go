package models

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/rehylas/wx/pkg/utils"
	"gopkg.in/mgo.v2/bson"
)

type Commuser struct {
	Userid    string `json:"userid"`
	Userpwd   string `json:"userpwd"`
	Openid    string `json:"openid"`
	State     int64  `json:"state"`
	Userlevel int64  `json:"userlevel"`
	Daymsgmax int64  `json:"daymsgmax"`
	Insertdt  string `json:"insertdt"`
	Paydt     string `json:"paydt"`
	Meo       string `json:"meo"`
}

// {
//     "_id" : ObjectId("5d4ec66dc15f077db7f15c44"),
//     "userid" : "80010121",
//     "userpwd" : "888889",
//     "userlevel" : 1.0,
//     "openid" : "ot_s20be187ldnuWyJhUwUxo6DjA",
//     "state" : NumberInt(1),
//     "insertdt" : "2017-06-30 17:55:21",
//     "mobile" : "139",
//     "tempid" : "40010121",
//     "daymsgmax" : 20.0,
//     "vaildt" : "2019-01-01 00:00:00"
// }

// 增加通信用户
func (this *Commuser) AddRec() (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	this.Userid = utils.RandomInt(8)
	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		return err
	}
	err = collect.Insert(this)
	if err != nil {
		return err
	}
	return nil

}

// 删除通信用户
func (this *Commuser) DelRec() (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		return err
	}
	whereSql := bson.M{"userid": this}
	err = collect.Remove(whereSql)
	if err != nil {
		return err
	}
	return nil

}

// 更新通信用户
func (this *Commuser) UpdateRecFields(fields []string) (err error) {

	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	j, _ := json.Marshal(this)
	json.Unmarshal(j, &m)

	whereSql := bson.M{"userid": this.Userid}
	setSql := make(map[string]interface{})
	for _, field := range fields {
		setSql[field] = m[field]
	}

	return collect.Update(whereSql, bson.M{"$set": setSql})

}

//获取开启的通信用户
func GetCommusers(commusers *[]Commuser) error {
	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{"state": 1}
	err = collect.Find(whereSql).All(commusers)

	if err != nil {
		beego.Error("GetModes err:", err)
		return err
	}
	return err
}

//获取所有的通信用户
func GetCommuserAll(commusers *[]Commuser) error {
	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{}
	err = collect.Find(whereSql).All(commusers)

	if err != nil {
		beego.Error("GetModes err:", err)
		return err
	}
	return err
}

//获取某一通信用户
func (this *Commuser) GetById(id string) (err error) {
	collect, err := getCollect(COLLECTNAME_COMMUSER)
	if err != nil {
		return err
	}
	whereSql := bson.M{"userid": id}
	err = collect.Find(whereSql).One(this)
	if err != nil {
		return err
	}
	return nil
}
