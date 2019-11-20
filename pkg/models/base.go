package models

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Base struct {
	ID  string `json:"_id"`
	Key string `json:"key"`
	Val string `json:"val"`
}

func (this *Base) GetById(collectName string, id string) (err error) {
	collect, err := getCollect(collectName)
	if err != nil {
		return err
	}
	whereSql := bson.M{"_id": id}
	err = collect.Find(whereSql).One(this)
	if err != nil {
		return err
	}
	return nil
}

func (this *Base) DelById(collectName string, id string) (err error) {
	collect, err := getCollect(collectName)
	if err != nil {
		return err
	}

	err = collect.Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil

}

func (this *Base) AddRec(collectName string) (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	collect, err := getCollect(collectName)
	if err != nil {
		return err
	}
	err = collect.Insert(this)
	if err != nil {
		return err
	}
	return nil

}

func (this *Base) UpdateRecFields(collectName string, fields []string) (err error) {

	collect, err := getCollect(collectName)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	j, _ := json.Marshal(this)
	json.Unmarshal(j, &m)

	whereSql := bson.M{"_id": this.ID}
	setSql := make(map[string]interface{})
	for _, field := range fields {
		setSql[field] = m[field]
	}

	return collect.Update(whereSql, bson.M{"$set": setSql})

}

func (this *Base) UpdateRec(collectName string) (err error) {

	collect, err := getCollect(collectName)
	if err != nil {
		return err
	}
	return collect.Update(bson.M{"_id": this.ID}, this)

}

func (this *Base) SetKey() error {
	client := getRedisClient()
	err := client.Set(DB_NAME+":"+this.Key, this.Val, 0).Err()
	return err
}

func (this *Base) GetKey() error {
	var err error
	client := getRedisClient()
	this.Val, err = client.Get(DB_NAME + ":" + this.Key).Result()
	return err
}

func SetKeyRedis(key, val string) error {
	client := getRedisClient()

	err := client.Set(DB_NAME+":"+key, val, 0).Err()
	return err
}

func GetKeyRedis(key string) (string, error) {

	client := getRedisClient()
	val, err := client.Get(DB_NAME + ":" + key).Result()
	return val, err
}
