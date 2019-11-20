package controllers

import (
	"encoding/json"

	"github.com/rehylas/wx/pkg/common"
	"github.com/rehylas/wx/pkg/models"
)

// Operations about Users
type AdminController struct {
	BaseController
}

type adminReq struct {
	Username  string `json:"username" bson:"username"`
	Logintype string `json:"logintype" bson:"logintype"`
	Loginpwd  string `json:"loginpwd" bson:"loginpwd"`
}

// @Title Login
// @Description 登录接口
// @Param  username  body  body    true        "用户名"
// @Param  logintype  body  body    true        "登录类型：1账号密码登录，2手机号短信验证码登录， 3指纹密码登录"
// @Param  loginpwd  body  body    true       "登录密码"
// @Param  msgcode  body  body    true        "短信验证码"
// @Param  fingerprint  body  body    false        "指纹"
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /admin/login/ [post]
func (this *AdminController) Login() {
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", this.Ctx.Request.Header.Get("Origin"))

	var reqData adminReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &reqData)
	if reqData.Username == "" {
		resp := common.MakeResp(common.ERROR_PARAM_ISNULL)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	if reqData.Username == "admin" && reqData.Loginpwd == "a1234567" {
		resp := common.MakeResp(common.ERROR_OK)
		//resp.Data = validuser
		this.Data["json"] = resp
		this.ServeJSON()
		user := models.Admin{}
		this.SetSession(common.SESSION_NAME, user)
		return
	} else {

		resp := common.MakeResp(common.ERROR_USERISNOTEXIST)
		this.Data["json"] = resp
		this.ServeJSON()

	}

	// u.SetSession(common.SESSION_NAME, *validuser)

	// cache.RefreshUserById(validuser.Id)
	// logs.Info("login SessionID:", u.CruSession.SessionID())
	// // resp := common.MakeResp2(common.ERROR_OK, &UserIncome{
	// // 	UserID:         validuser.Id,
	// // 	Inviteesnumber: validuser.InviteesNumber,
	// // 	Myincome:       validuser.Myincome,
	// // })
	// resp := common.MakeResp(common.ERROR_OK)
	// resp.Data = validuser
	// u.Data["json"] = resp
	// u.ServeJSON()

}
