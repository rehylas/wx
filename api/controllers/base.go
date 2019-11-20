package controllers

import (
	"strings"

	"github.com/rehylas/wx/pkg/common"
	"github.com/rehylas/wx/pkg/models"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	userId         string
	userName       string
}

//const AUTH_ACTIONS = "ModeController.Del|ModeController.Add"

const AUTH_ACTIONS = "ModeController.Do|ModeController.Do2"

//前期准备
func (this *BaseController) Prepare() {

	beego.Info("RequestBody:", string(this.Ctx.Input.RequestBody))

	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)

	// noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start"
	// isNoAuth := strings.Contains(noAuth, self.actionName)
	beego.Info("controllerName, actionName:", controllerName, actionName)

	act := controllerName + "." + actionName
	if strings.Contains(AUTH_ACTIONS, act) && beego.BConfig.RunMode != "test" {
		this.auth()
	}
}

//检查是否已经登录，如果已经登录，则继续， 否则返回 session 过期
func (this *BaseController) auth() {
	//log.Println("request:", u.Ctx.Request)
	//根据session获取用户信息
	su := this.GetSession(common.SESSION_NAME)
	admin := su.(models.Admin)

	beego.Info("common查询用户session,sessionid=", this.CruSession.SessionID())
	if su == "" || su == nil || admin.Username == "" || admin.Id == "" {
		logs.Error("common查询用户session不存在,su=", su)
		resp := common.MakeResp(common.SESSION_ERROR)
		this.Data["json"] = resp
		this.ServeJSON()
		this.StopRun()
		return
	}

}
