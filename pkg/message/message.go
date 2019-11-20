package message

/***********************************************
该模块是系统内部的消息传递，消息推送，主要是把交易情况
k线， 聊天等信息，推送到用户APP, APP程序进行数据处理
或展示
************************************************/

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/rehylas/wx/pkg/database"
)

const DT_OTC_NEW_TRADE = "otc_new_trade"
const DT_OTC_ORDER_STATE_UPDATE = "otc_order_state_update"
const DT_CC_ORDER_STATE_UPDATE = "cc_order_state_update"
const DT_CC_NEW_TRADE = "cc_new_trade"

type MessageData struct {
	MsgType  string `json:"msgtype"`
	DataType string `json:"datatype"`
	ToUserId string `json:"touserid"`
	Subtitle string `json:"subtitle"`
	Data     interface{}
}

func init() {

}

// @Title SystemMessage
// @Description 系统信息推送， 会广播到所有APP
// @Param   datatype 数据分类比如 tradefilled , data,任意数据；  注意：此项内容需与app前端开发共同商议确定
// 广播会有较大压力，慎用
func AllSystemMessage(datatype string, data interface{}) {
	msg := MessageData{MsgType: "system", DataType: "datatype", Data: data}

	msgdata, err := json.Marshal(msg)
	if err != nil {
		beego.Error("Marshal error:", err)
		return
	}
	pubmsg(msg.MsgType, msgdata)

}

// 同 AllSystemMessage
func SystemMessage(datatype string, data interface{}) {
	AllSystemMessage(datatype, data)
}

// @Title SystemMessage
// @Description 系统信息推送， 会广播到所有APP
// @Param   datatype 数据分类比如 tradefilled , data,任意数据；  注意：此项内容需与app前端开发共同商议确定
// 子系统之间
func SubSystemMessage(datatype string, data interface{}) {
	msg := MessageData{MsgType: "subsystem", DataType: "datatype", Data: data}

	msgdata, err := json.Marshal(msg)
	if err != nil {
		beego.Error("Marshal error:", err)
		return
	}
	pubmsg(msg.MsgType, msgdata)
}

// @Title P2PMessage
// @Description p2p信息推送， 推送到具体的 userid app
// @Param   datatype 数据分类比如 tradefilled , data,任意数据； ToUserId  注意：此项内容需与app前端开发共同商议确定
func P2PMessage(datatype string, data interface{}, touserid string) {
	msg := MessageData{MsgType: "p2p", DataType: "datatype", ToUserId: touserid, Data: data}

	msgdata, err := json.Marshal(msg)
	if err != nil {
		beego.Error("Marshal error:", err)
		return
	}
	pubmsg(msg.MsgType, msgdata)
}

// @Title ChatMessage
// @Description 聊天功能，向用户发送聊天信息
// @Param   datatype 数据分类比如 txt  img , 注意：此项内容需与app前端开发共同商议确定 data,任意数据； ToUserId
func ChatMessage(datatype string, data interface{}, touserid string) {
	msg := MessageData{MsgType: "chat", DataType: datatype, ToUserId: touserid, Data: data}

	msgdata, err := json.Marshal(msg)
	if err != nil {
		beego.Error("Marshal error:", err)
		return
	}
	pubmsg(msg.MsgType, msgdata)
	beego.Debug("已向redis提交")
}

// @Title MarketMessage
// @Description 行情信息，会广播全推
// @Param   datatype 数据分类比如 tradefilled , data,任意数据；  注意：此项内容需与app前端开发共同商议确定
func MarketMessage(datatype string, data interface{}) {
	msg := MessageData{MsgType: "market", DataType: datatype, Data: data}

	msgdata, err := json.Marshal(msg)
	if err != nil {
		beego.Error("Marshal error:", err)
		return
	}
	pubmsg(msg.MsgType, msgdata)
}

// @Title MulticastMessage
// @Description  组播消息， 用于行情
// @Param
func MulticastMessage(datatype string, subtitle string, data interface{}) {
	//subtitle
	msg := MessageData{MsgType: "multicast", DataType: datatype, Subtitle: subtitle, Data: data}

	msgdata, err := json.Marshal(msg)
	if err != nil {
		beego.Error("Marshal error:", err)
		return
	}
	pubmsg(msg.MsgType, msgdata)
}

///////////////////////////////////////////////////////////////////////////////////////////////
//内部函数
// func encodemsg
func pubmsg(channel string, msgdata []byte) {
	client := database.GetRedisClient() //*redis.Client

	err := client.Publish(channel, msgdata).Err()
	if err != nil {
		beego.Error("Pubmsg error:", err)
	}

}
