package models

import (
	"encoding/json"

	"github.com/rehylas/wx/pkg/utils"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

// 账户记录
type AccDetail struct {
	ID             string  `json:"id"`
	Date           string  `json:"date"`
	Balance        float64 `json:"balance"`
	Available      float64 `json:"available"`
	Income         float64 `json:"income"`
	Prebalance     float64 `json:"prebalance"`
	Commission     float64 `json:"commission"`
	Positionprofit float64 `json:"positionprofit"`
	Closeprofit    float64 `json:"closeprofit"`
	Accountid      string  `json:"accountid"`
	Vtaccountid    string  `json:"vtaccountid"`
}

/*
{
    "_id" : ObjectId("5dcee7a0c291b313b0d125bf"),
    "available" : 0.77,
    "gatewayname" : "ctp",
    "positionprofit" : 0.0,
    "prebalance" : 0.77,
    "commission" : 0.0,
    "vtaccountid" : "ctp.900703691",
    "date" : "2019-11-16",
    "closeprofit" : 0.0,
    "balance" : 0.77,
    "margin" : 0.0,
    "id" : "954854",
    "accountid" : "900703691"
}
*/

// 增加交易模型
func (this *AccDetail) AddRec() (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	this.ID = utils.RandomInt(6)
	collect, err := getCollect(COLLECTNAME_ACCDETAIL)
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
func (this *AccDetail) DelRec() (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	collect, err := getCollect(COLLECTNAME_ACCDETAIL)
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
func (this *AccDetail) UpdateRecFields(fields []string) (err error) {

	collect, err := getCollect(COLLECTNAME_ACCDETAIL)
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

func GetAccDetails(accdetails *[]AccDetail) error {
	collect, err := getCollect(COLLECTNAME_ACCDETAIL)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{}
	err = collect.Find(whereSql).Sort("date").All(accdetails)

	if err != nil {
		beego.Error("GetModes err:", err)
		return err
	}
	return err
}

func FixAccDetails(amount float64) error {
	var accdetails []AccDetail
	err := GetAccDetails(&accdetails)
	if err != nil {
		return err
	}

	for _, rec := range accdetails {
		if rec.Income == 0 {
			rec.FixIncome(amount)
		}

	}

	return nil

}

func (this *AccDetail) FixIncome(amount float64) error {
	collect, err := getCollect(COLLECTNAME_ACCDETAIL)
	if err != nil {
		return err
	}

	whereSql := bson.M{"id": this.ID} //"income": 0
	setSql := bson.M{"income": this.Balance - amount}
	return collect.Update(whereSql, bson.M{"$set": setSql})

}

func (this *AccDetail) GetAccDetailNow() error {

	collect, err := getCollect(COLLECTNAME_ACCDETAIL)
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

func (this *AccDetail) GetById(id string) (err error) {
	collect, err := getCollect(COLLECTNAME_ACCDETAIL)
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
