package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"github.com/rehylas/wx/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
	mgo "gopkg.in/mgo.v2"
)

const DB_NAME = "aft"

const (
	COLLECTNAME_USER      = "users"
	COLLECTNAME_LOG       = "bizlog"
	COLLECTNAME_CHATMSG   = "chatmsg"
	COLLECTNAME_SYSCONFIG = "systemconfig"

	COLLECTNAME_MSG        = "message"
	COLLECTNAME_USERWALLET = "userwallet"
)

const (
	COLLECTNAME_ORDER     = "order"
	COLLECTNAME_MODE      = "mode"
	COLLECTNAME_FUTURE    = "futures"
	COLLECTNAME_ACCDETAIL = "accountdetail"
	COLLECTNAME_COMMUSER  = "commuser"
	COLLECTNAME_SUMM      = "summ" //统计表
)

/////////////////////////////////////////////////////////////////////////////////////////////////
//内部函数
//获取表指针
func getCollect(collectname string) (*mgo.Collection, error) {
	return database.GetCollect(collectname)
}

func getTransactionCollect(collectname string) *mongo.Collection {
	db := database.TransactionDB
	if db == nil {
		logs.Error("db初始化为空！")
		return nil
	}
	collection := db.Collection(collectname)
	return collection
}

func getRedisClient() *redis.Client {
	return database.GetRedisClient()
}
