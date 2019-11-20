package models

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

//{name:"", "heatbeat":"dt", "bizlog":"", "state":0 }
type SystemInfo struct {
	Name     string `json:"name"`     //日期时间
	Heatbeat string `json:"heatbeat"` //日期时间
	Bizlog   string `json:"bizlog"`   //日期时间
	State    int    `json:"state"`    //日期时间
	Desp     string `json:"desp"`     //日期时间
}

//
func (this *SystemInfo) GetSystemByName(name string) error {
	key := "system_" + name
	val, err := GetKeyRedis(key)
	if err != nil {
		return err
	}
	beego.Debug(val, err)
	err = json.Unmarshal([]byte(val), this)
	beego.Debug("unmarshal:", err)
	return err
}
