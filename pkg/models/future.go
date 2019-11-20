package models

import "gopkg.in/mgo.v2/bson"

type Future struct {
	Symbol    string  `json:"symbol"`
	Exchange  string  `json:"exchange"`
	PriceTick float64 `bson:"priceTick"`
	Type      string  `json:"type"`
	Size      int     `json:"size"`
}

func (this *Future) GetFutureBySymbol(symbol string) error {
	collect, err := getCollect(COLLECTNAME_FUTURE)
	if err != nil {
		return err
	}
	whereSql := bson.M{"symbol": symbol}
	err = collect.Find(whereSql).One(this)
	if err != nil {
		return err
	}
	return nil
}
