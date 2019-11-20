package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/rehylas/wx/pkg/common"
	"github.com/rehylas/wx/pkg/models"
)

// Operations about Mode
type CommuserController struct {
	BaseController
}

type commuserReq struct {
	Userid    string `json:"userid"`    //用户编号
	Userpwd   string `json:"userpwd"`   //用户密码
	Openid    string `json:"openid"`    //微信信息
	State     int64  `json:"state"`     //用户状态
	Userlevel int64  `json:"userlevel"` //支付状态
	Daymsgmax int64  `json:"daymsgmax"`
	Insertdt  string `json:"insertdt"` //插入时间
	Paydt     string `json:"paydt"`    //支付时间
	Meo       string `json:"meo"`      //备注
}

// @Title List
// @Description 请求交易模型
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /commuser/list/ [get]
func (this *CommuserController) List() {

	var commusers []models.Commuser
	err := models.GetCommuserAll(&commusers)

	if err != nil {
		beego.Error("mode.GetCommuserAll() err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	resp.Data = commusers
	this.Data["json"] = resp
	this.ServeJSON()
	return

}

// @Title Add
// @Description  添加交易模型
// @Success 0000 {object} models.commuser
// @Failure 9999 system is error
// @router /commuser/add/ [post]
func (this *CommuserController) Add() {

	var reqData modeReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &reqData)
	if reqData.Symbol == "" {
		resp := common.MakeResp(common.ERROR_PARAM_ISNULL)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	mode := models.Mode{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &mode)
	err := mode.AddRec()
	if err != nil {
		beego.Error("mode.AddRec() err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	this.Data["json"] = resp
	this.ServeJSON()
	return

}

// @Title Update
// @Description 修改交易模型
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /commuser/update/ [post]
func (this *CommuserController) Update() {

	var reqData commuserReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &reqData)
	if reqData.Userid == "" {
		resp := common.MakeResp(common.ERROR_PARAM_ISNULL)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	commuser := models.Commuser{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &commuser)

	var fieds []string
	fieds = append(fieds, "userpwd")
	fieds = append(fieds, "state")
	fieds = append(fieds, "userlevel")
	fieds = append(fieds, "paydt")
	fieds = append(fieds, "meo")
	err := commuser.UpdateRecFields(fieds)

	if err != nil {
		beego.Error("commuser.UpdateRecFields() err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	this.Data["json"] = resp
	this.ServeJSON()
	return

}

// @Title Delete
// @Description 修改交易模型
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /commuser/del/ [post]
func (this *CommuserController) Del() {

	var reqData commuserReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &reqData)
	if reqData.Userid == "" {
		resp := common.MakeResp(common.ERROR_PARAM_ISNULL)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	commuser := models.Commuser{Userid: reqData.Userid}
	commuser.DelRec()

	resp := common.MakeResp(common.ERROR_OK)
	this.Data["json"] = resp
	this.ServeJSON()
	return

}
