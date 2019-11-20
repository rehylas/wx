package models

import (
	"../database"

	"gopkg.in/mgo.v2"
)

const COLLECTNAME_COMMUSER = "commuser"

/////////////////////////////////////////////////////////////////////////////////////////////////
//内部函数
//获取表指针
func getCollect(collectname string) (*mgo.Collection, error) {
	return database.GetCollect(collectname)
}
