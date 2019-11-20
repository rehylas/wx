package biz

import (
	"fmt"
	"math/rand"
	"time"

	"../models"
)

func init() {

}

const DAY_MAX_COUNT = 100

var msgControl *MsgControl

type node struct {
	Openid    string
	SendCount int
}

type MsgControl struct {
	Nodes map[string]*node
	Date  string
}

func init() {
	rand.New(rand.NewSource(10))
}

func CheckMsgSend(openid string) bool {

	if msgControl == nil {
		msgControl = &MsgControl{Nodes: make(map[string]*node), Date: time.Now().Format("2006-01-02")}
	}

	if msgControl.Date != time.Now().Format("2006-01-02") {
		msgControl.Set0()
	}

	value, ok := msgControl.Nodes[openid]
	if ok {
		value.SendCount += 1
		if value.SendCount >= DAY_MAX_COUNT {
			return false
		} else {
			return true
		}
	} else {
		msgControl.Nodes[openid] = &node{Openid: openid, SendCount: 0}
		return true
	}

}

func (mc *MsgControl) Set0() {

	for _, v := range mc.Nodes {
		v.SendCount = 0
	}

}

// if exist  return ,  if not, back user info
func GetCreateCommuser(openid string) (*models.CommUser, error) {
	u := &models.CommUser{Openid: openid}
	u.GetCommUserByOpenId(openid)
	if u.Userid != "" {
		return u, nil
	}

	u = InitCommuser(openid)
	err := u.AddCommUser()
	if err != nil {
		return nil, err
	}
	return u, nil

}

func InitCommuser(openid string) *models.CommUser {

	dt := time.Now().Format("2006-01-02 15:04:05")
	u := &models.CommUser{Openid: openid, State: 0, Userlevel: 1, Daymsgmax: 20, Insertdt: dt}
	u.Userid = fmt.Sprintf("%d", int32((rand.Int() % 100000000)))
	u.Userpwd = fmt.Sprintf("%d", int32((rand.Int() % 1000000)))

	return u
}

// func Receiver(w http.ResponseWriter, r *http.Request) {
// 	token := beego.AppConfig.String("token")   // 微信公众平台的Token "1234567890"
// 	appid := beego.AppConfig.String("appid")   // 微信公众平台的AppID "wx1cf0a81453f5bed2"
// 	secret := beego.AppConfig.String("secret") // 微信公众平台的AppSecret "efa1f0dcd9a0172924863349664d5e82"
// 	// 仅被动响应消息时可不填写appid、secret
// 	// 仅主动发送消息时可不填写token
// 	mp := weixinmp.New(token, appid, secret)
// 	// 检查请求是否有效
// 	// 仅主动发送消息时不用检查
// 	if !mp.Request.IsValid(w, r) {
// 		return
// 	}

// 	beego.Debug("", mp.Request)

// 	//err := mp.SendTextMsg("og7lr1FI2yDdzDJVB2qA__SpM5NA", "txtmsg")
// 	err := mp.SendTempMsg("og7lr1FI2yDdzDJVB2qA__SpM5NA", "txtmsg", "AqhideqORuAAxwGWZJpFjgy05AFROUQ8jngTiJ8aVqk")
// 	beego.Debug("err:", err)

// 	// 判断消息类型
// 	if mp.Request.MsgType == weixinmp.MsgTypeText {

// 		// 回复消息
// 		mp.ReplyTextMsg(w, "Hello, 世界")
// 	}
// }
