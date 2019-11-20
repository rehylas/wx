package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

// ctp 交易指令， 只要单子下到该表， 会被ctp执行程序执行
type Order struct {
	Insertdt   string  `json:"insertdt"`
	Symbol     string  `json:"symbol"`
	Type       string  `json:"type"`
	Offset     string  `json:"offset"`
	Price      float64 `json:"price"`
	Vol        int     `json:"vol"`
	Date       string  `json:"date"`
	State      int     `json:"state"`
	Execdt     string  `json:"execdt"`
	Orderidsys string  `json:"orderidsys"`

	//CTP撤单需要
	Exchange  string `json:"exchange"`
	FrontID   string `bson:"frontID"`
	SessionID string `bson:"sessionID"`
	OrderID   string `bson:"orderID"`
}

// 下单委托
// type : buy/sell/cancel  offset: open/close/close_today
func (this *Order) AddRec() (err error) {

	//this.ID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	collect, err := getCollect(COLLECTNAME_ORDER)
	if err != nil {
		return err
	}
	err = collect.Insert(this)
	if err != nil {
		return err
	}
	return nil

}

func GetOrdersAll(orders *[]Order) error {
	collect, err := getCollect(COLLECTNAME_ORDER)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{}
	err = collect.Find(whereSql).All(orders)

	if err != nil {
		beego.Error("GetOrdersAll err:", err)
		return err
	}
	return err
}

func GetOrderToday(orders *[]Order) error {
	collect, err := getCollect(COLLECTNAME_ORDER)
	if err != nil {
		beego.Error("getCollection", err)
		return err
	}
	whereSql := bson.M{}
	err = collect.Find(whereSql).Sort("-insertdt").Limit(20).All(orders)
	if err != nil {
		beego.Error("GetOrderToday err:", err)
		return err
	}
	return err
}

func GetOrderidBySysOrderid(SysOrderid string) (string, error) {
	order := &Order{}
	collect, err := getCollect(COLLECTNAME_ORDER)
	if err != nil {
		return "", err
	}
	whereSql := bson.M{"orderidsys": SysOrderid}
	err = collect.Find(whereSql).One(order)
	if err != nil {
		return "", err
	}
	return order.OrderID, nil
}

func GetOrderBySysOrderid(SysOrderid string) (*Order, error) {
	order := &Order{}
	collect, err := getCollect(COLLECTNAME_ORDER)
	if err != nil {
		return nil, err
	}
	whereSql := bson.M{"orderidsys": SysOrderid}
	err = collect.Find(whereSql).One(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

/*
order
{ "_id" : ObjectId("5d1df768824a62f67874434c"), "vol" : 1, "symbol" : "ru1909",
"insertdt" : "2019-07-04 20:56:08", "offset" : "close",
"date" : "2019-07-04", "orderid" : "", "state" : 2, "execdt" : "", "type" : "sell", "price" : 11260 }

codemode
{ "_id" : ObjectId("5d4190736d4201e73984656d"), "indt" : "", "outprice" : 0, "type" : "", "Exec" : 1,
 "outdt" : "", "vol" : 1, "signaldt" : "", "symbol" : "zn1909", "mustWinOut" : 0,
 "insertdt" : "2019-07-31 20:58:27", "inrate" : 0, "state" : 0, "mustOut" : 0,
"income" : 0, "date" : "2019-07-31", "mode_name" : "tk", "maxfail" : 0, "maxwin" : 0, "inprice" : 0 }

#4. 获取执行指令
def db_exec_getBSOrder_One( ):

    collectionname = 'order'
    data_Sql = { 'state':0 }
    return db_select( DB_NAME, collectionname, sql = data_Sql , nlimit = 1  )

    pass

#5. 更新执行指令
def db_exec_updateBSOrder_One( rec ):
    collectionname = 'order'
    sql = { u"_id": rec['_id']  }
    return db_update( DB_NAME, collectionname, rec, sql,  upsert=False  )

    pass

#6. 插入执行指令
def db_exec_insertBSOrder_One(type,  offset,  symbol, price, vol,  exchange = ''  ): # type : buy/sell/cancel  offset: open/close/close_today
    orderInfo ={ 'date':getDate(), 'insertdt':getDt(), 'type':type,  'offset':offset,  'symbol':symbol, 'price':price, 'vol':vol, 'orderid':'', 'execdt':'', 'state':0 }
    collectionname = 'order'
    db_insertOne( DB_NAME, collectionname, orderInfo )
	pass

*/
