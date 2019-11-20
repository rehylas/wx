package models

import (
	"encoding/json"

	"github.com/rehylas/wx/pkg/utils"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

// 统计数据
type Summ struct {
	Date       string `json:"date"`
	Totalusers int    `json:"totalusers"`
	Payusers   int    `json:"payusers"`
}

/*
{
    "_id" : ObjectId("5dd4acbfa1da650b7c481bee"),
    "date" : "2019-11-20",
    "totalusers" : 100.0,
    "payusers" : 5.0
}
*/

func (this *Summ) AddRec() (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	this.Date = utils.DateStr()
	collect, err := getCollect(COLLECTNAME_SUMM)
	if err != nil {
		return err
	}
	err = collect.Insert(this)
	if err != nil {
		return err
	}
	return nil

}

func (this *Summ) UpdateRecFields(fields []string) (err error) {

	collect, err := getCollect(COLLECTNAME_SUMM)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	j, _ := json.Marshal(this)
	json.Unmarshal(j, &m)

	whereSql := bson.M{"id": this.Date}
	setSql := make(map[string]interface{})
	for _, field := range fields {
		setSql[field] = m[field]
	}

	return collect.Update(whereSql, bson.M{"$set": setSql})

}

func GetSumms(summs *[]Summ) error {
	collect, err := getCollect(COLLECTNAME_SUMM)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{}
	err = collect.Find(whereSql).Sort("date").All(summs)

	if err != nil {
		beego.Error("GetModes err:", err)
		return err
	}
	return err
}

func (this *AccDetail) GetSummNow() error {

	collect, err := getCollect(COLLECTNAME_SUMM)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{}
	err = collect.Find(whereSql).Sort("-date").One(&this)

	if err != nil {
		beego.Error("GetModes err:", err)
		return err
	}
	return err
}

//统计用户总数，支付用户数
func SummUser() (totals, pays int) {
	collect, err := getCollect( COLLECTNAME_COMMUSER)
	if err != nil {
		beego.Error("getCollection", err)
		return 0, 0
	}
	whereSql := bson.M{}
	totals, err = collect.Find(whereSql).Count()

	whereSql = bson.M{"userlevel": 2}
	pays, err = collect.Find(whereSql).Count()

	return totals, pays

}
