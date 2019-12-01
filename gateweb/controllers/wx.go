package controllers

import (
	"encoding/json"
	"fmt"

	"../biz"
	"../common"
	"../models"

	"github.com/astaxie/beego"
	"github.com/sidbusy/weixinmp"
)

var mp *weixinmp.Weixinmp

func init() {

	token := beego.AppConfig.String("token")   // 微信公众平台的Token "1234567890"
	appid := beego.AppConfig.String("appid")   // 微信公众平台的AppID "wx1cf0a81453f5bed2"
	secret := beego.AppConfig.String("secret") // 微信公众平台的AppSecret "efa1f0dcd9a0172924863349664d5e82"
	// 仅被动响应消息时可不填写appid、secret
	// 仅主动发送消息时可不填写token
	mp = weixinmp.New(token, appid, secret)
}

// Operations about Coins
type WxController struct {
	beego.Controller
}

// @Title Receiver
// @Description Receiver
func (u *WxController) Receiver() {

	w := u.Ctx.ResponseWriter.ResponseWriter
	r := u.Ctx.Request

	// 检查请求是否有效
	// 仅主动发送消息时不用检查
	if !mp.Request.IsValid(w, r) {
		beego.Debug("IsValid: false")
		return
	}

	// beego.Debug("", mp.Request)
	// //err := mp.SendTextMsg("og7lr1FI2yDdzDJVB2qA__SpM5NA", "txtmsg")
	// err := mp.SendTempMsg("og7lr1FI2yDdzDJVB2qA__SpM5NA", "txtmsg", "AqhideqORuAAxwGWZJpFjgy05AFROUQ8jngTiJ8aVqk")
	// beego.Debug("err:", err)

	beego.Info("Receiver:", mp.Request.FromUserName, mp.Request.MsgType, mp.Request.Content)
	// 判断消息类型
	if mp.Request.MsgType == weixinmp.MsgTypeText {

		// 回复消息
		mp.ReplyTextMsg(w, "欢迎使用帮您看着")
	}

	if mp.Request.MsgType == weixinmp.MsgTypeEvent {

		beego.Info("Receiver event:")
		fmt.Printf("Request data:%v", mp.Request)
		beego.Info("Request data:", mp.Request)
		if mp.Request.Event == weixinmp.EventSubscribe { //扫描关注

			user, err := biz.GetCreateCommuser(mp.Request.FromUserName)
			if err != nil {
				beego.Error("biz.GetCreateCommuser:", mp.Request.FromUserName, err)
			}

			msg := fmt.Sprintf("欢迎使用帮您看着,您的股讯账号是:%s 股讯密码:%s.请把账号发给客服开通,客服微信号:z1681234\n", user.Userid, user.Userpwd)
			mp.ReplyTextMsg(w, msg)
			ReportNewUser()

		}
		// 回复消息
		//mp.ReplyTextMsg(w, "欢迎使用帮您看着")
	}

}

/*
访问地址：
www.bangnikanzhe.com/DXBizGate/server?cmd=sendmsg4client
参数：
{'userid':'10011001','userpwd':'123456','txt':'002230 科大讯飞要涨了' }
*/

type ClientMsg struct {
	Userid  string `json:"userid"`
	Userpwd string `json:"userpwd"`
	Txt     string `json:"txt"`
}

func (u *WxController) Sendtempmsg() {
	//tempId := "AqhideqORuAAxwGWZJpFjgy05AFROUQ8jngTiJ8aVqk"
	tempId := "NQhyxhnM40-OGm2SPafHqCQ4yJFkY2BgEBV2DMJ7cUI"
	var msg ClientMsg
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &msg)
	if err != nil {
		beego.Error("Sendtempmsg json.Unmarshal error:", err, u.Ctx.Input.RequestBody)
	}

	beego.Debug("RequestBody:", string(u.Ctx.Input.RequestBody), msg.Userid)

	openid := "" //og7lr1FI2yDdzDJVB2qA__SpM5NA
	//openid := getOpenidByUserid(msg.Userid)
	var user models.CommUser
	(&user).GetCommUserById(msg.Userid)
	beego.Debug("user:", user)
	//if user.s
	if user.State != 1 {
		resp := common.MakeResp(common.ERROR_USER_STATE)
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}

	if user.Userpwd != msg.Userpwd {
		resp := common.MakeResp(common.ERROR_USERIDPWD_ERR)
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}

	if biz.CheckMsgSend(msg.Userid) == false {
		resp := common.MakeResp(common.ERROR_MAXMSG_SEND)
		u.Data["json"] = resp
		u.ServeJSON()
		return
	}

	openid = user.Openid
	senddata, err := mp.SendTempMsg(openid, msg.Txt, tempId)
	beego.Debug("senddata:", senddata)
	if err != nil {
		beego.Debug("err:", err)
	}

	resp := common.MakeResp(common.ERROR_OK)
	u.Data["json"] = resp
	u.ServeJSON()

}

//
func ReportNewUser() {
	tempId := "NQhyxhnM40-OGm2SPafHqCQ4yJFkY2BgEBV2DMJ7cUI"
	openid := "ot_s20be187ldnuWyJhUwUxo6DjA"

	mp.SendTempMsg(openid, "有新用户注册", tempId)

}
