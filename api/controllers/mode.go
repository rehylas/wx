package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/rehylas/wx/pkg/common"
	"github.com/rehylas/wx/pkg/models"
)

// Operations about Mode
type ModeController struct {
	BaseController
}

type modeReq struct {
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
	ModeType string `json:"modetype"`
	BsType   string `json:"bstype"`
	Vol      int    `json:"vol"`
	VolDef   int    `json:"voldef"`
	Enable   int    `json:"enable"` // 0  不打开   1 打开
	Exec     int    `json:"exec"`   // 0  不执行   1 默认执行   2 人工执行； 如果不执行则不委托， 如果默认执行，则读取def下单， 如果人工执行则读取vol下单
	State    int    `json:"state"`  // 状态 0 等待   1 进入  2 出局

}

// @Title List
// @Description 请求交易模型
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /mode/login/ [get]
func (this *ModeController) List() {

	var modes []models.Mode
	err := models.GetModesAll(&modes)

	if err != nil {
		beego.Error("mode.GetModes() err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	resp.Data = modes
	this.Data["json"] = resp
	this.ServeJSON()
	return

}

// @Title Add
// @Description  添加交易模型
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /mode/add/ [post]
func (this *ModeController) Add() {

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
// @router /mode/update/ [post]
func (this *ModeController) Update() {

	var reqData modeReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &reqData)
	if reqData.ID == "" {
		resp := common.MakeResp(common.ERROR_PARAM_ISNULL)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	mode := models.Mode{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &mode)

	var fieds []string
	fieds = append(fieds, "symbol")
	fieds = append(fieds, "enable")
	fieds = append(fieds, "exec")
	fieds = append(fieds, "state")
	fieds = append(fieds, "vol")
	fieds = append(fieds, "voldef")
	err := mode.UpdateRecFields(fieds)

	if err != nil {
		beego.Error("mode.UpdateRecFields() err:", err)
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
// @router /mode/del/ [post]
func (this *ModeController) Del() {

	var reqData modeReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &reqData)
	if reqData.ID == "" {
		resp := common.MakeResp(common.ERROR_PARAM_ISNULL)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	mode := models.Mode{ID: reqData.ID}
	mode.DelRec()

	resp := common.MakeResp(common.ERROR_OK)
	this.Data["json"] = resp
	this.ServeJSON()
	return

}
