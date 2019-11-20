package models

import (
	"encoding/json"

	"github.com/rehylas/wx/pkg/utils"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

//State    int    `json:"state"`  // 状态 0 等待   1 进入  2 出局
const MODE_STATE_WAITE = 0
const MODE_STATE_IN = 1
const MODE_STATE_OUT = 2

// 0  不执行   1 默认执行   2 人工执行；
const MODE_EXEC_0 = 0
const MODE_EXEC_DEF = 1
const MODE_EXEC_CUST = 2

// 交易模型
type Mode struct {
	ID       string  `json:"id"`
	Symbol   string  `json:"symbol"`
	ModeType string  `json:"modetype"`
	BsType   string  `json:"bstype"`
	Vol      int     `json:"vol"`
	VolDef   int     `json:"voldef"`
	Enable   int     `json:"enable"` // 0  不打开   1 打开
	Exec     int     `json:"exec"`   // 0  不执行   1 默认执行   2 人工执行； 如果不执行则不委托， 如果默认执行，则读取def下单， 如果人工执行则读取vol下单
	State    int     `json:"state"`  // 状态 0 等待   1 进入  2 出局
	Execdt   string  `json:"execdt"`
	Inprice  float64 `json:"inprice"`
}

/*
{
    "_id" : "5d945cd8443400757ce6ed95",
    "symbol" : "ru2001",
    "modetype" : "gytk",
    "bstype" : "buy",
    "vol" : 2,
    "voldef" : 1,
    "enable" : 1,
    "exec" : 0,
    "state" : 0,
    "execdt" : ""
}
*/

// 增加交易模型
func (this *Mode) AddRec() (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	this.ID = utils.RandomInt(8)
	collect, err := getCollect(COLLECTNAME_MODE)
	if err != nil {
		return err
	}
	err = collect.Insert(this)
	if err != nil {
		return err
	}
	return nil

}

// 删除交易模型
func (this *Mode) DelRec() (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	collect, err := getCollect(COLLECTNAME_MODE)
	if err != nil {
		return err
	}
	whereSql := bson.M{"id": this.ID}
	err = collect.Remove(whereSql)
	if err != nil {
		return err
	}
	return nil

}

// 更新模型信息
func (this *Mode) UpdateRecFields(fields []string) (err error) {

	collect, err := getCollect(COLLECTNAME_MODE)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	j, _ := json.Marshal(this)
	json.Unmarshal(j, &m)

	whereSql := bson.M{"id": this.ID}
	setSql := make(map[string]interface{})
	for _, field := range fields {
		setSql[field] = m[field]
	}

	return collect.Update(whereSql, bson.M{"$set": setSql})

}

func GetModes(modes *[]Mode) error {
	collect, err := getCollect(COLLECTNAME_MODE)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{"enable": 1}
	err = collect.Find(whereSql).All(modes)

	if err != nil {
		beego.Error("GetModes err:", err)
		return err
	}
	return err
}

func GetModesAll(modes *[]Mode) error {
	collect, err := getCollect(COLLECTNAME_MODE)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{}
	err = collect.Find(whereSql).All(modes)

	if err != nil {
		beego.Error("GetModes err:", err)
		return err
	}
	return err
}

func (this *Mode) GetById(id string) (err error) {
	collect, err := getCollect(COLLECTNAME_MODE)
	if err != nil {
		return err
	}
	whereSql := bson.M{"id": id}
	err = collect.Find(whereSql).One(this)
	if err != nil {
		return err
	}
	return nil
}
