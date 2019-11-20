package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/rehylas/wx/api/service"
	"github.com/rehylas/wx/pkg/common"
	"github.com/rehylas/wx/pkg/models"
)

//{name:"", "heatbeat":"dt", "bizlog":"", "state":0 }

// Operations about Mode
type SystemInfoController struct {
	BaseController
}

// @Title List
// @Description 获取系统信息
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /systeminfo/list/ [get]
func (this *SystemInfoController) List() {

	// GetSystemInfos(systems *[]models.SystemInfo)
	var systems []models.SystemInfo
	err := service.GetSystemInfos(&systems)

	if err != nil {
		beego.Error("mode.GetModes() err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	resp.Data = systems
	this.Data["json"] = resp
	this.ServeJSON()
	return

}

// @Title Update
// @Description 获取系统信息
// @Success 0000 {object} models.SystemInfo
// @Failure 9999 system is error
// @router /systeminfo/update/ [post]
func (this *SystemInfoController) Update() {

	// GetSystemInfos(systems *[]models.SystemInfo)
	var system models.SystemInfo

	json.Unmarshal(this.Ctx.Input.RequestBody, &system)
	if system.Name == "" {
		resp := common.MakeResp(common.ERROR_PARAM_ISNULL)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}
	strval, _ := json.Marshal(system)
	models.SetKeyRedis(system.Name, string(strval))

	resp := common.MakeResp(common.ERROR_OK)
	this.Data["json"] = resp
	this.ServeJSON()
	return

}
