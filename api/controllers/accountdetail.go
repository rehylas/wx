package controllers

import (
	"github.com/astaxie/beego"

	"github.com/rehylas/wx/api/service"
	"github.com/rehylas/wx/pkg/common"
	"github.com/rehylas/wx/pkg/models"
)

// Operations about Mode
type AccountDetailController struct {
	BaseController
}

// @Title List
// @Description 账户记录
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /accdetail/list/ [get]
func (this *AccountDetailController) List() {

	var accDetails []models.AccDetail
	var accDetail service.AccDetail

	err := accDetail.GetList(&accDetails)
	if err != nil {
		beego.Error("accDetail  GetList err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	resp.Data = accDetails
	this.Data["json"] = resp
	this.ServeJSON()
	return

}
