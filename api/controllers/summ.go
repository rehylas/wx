package controllers

import (
	"github.com/astaxie/beego"

	"github.com/rehylas/wx/pkg/common"
	"github.com/rehylas/wx/pkg/models"
)

type SummController struct {
	BaseController
}

// @Title List
// @Description 统计数据
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /summ/list/ [get]
func (this *SummController) List() {

	var summs []models.Summ

	err := models.GetSumms(&summs)
	if err != nil {
		beego.Error("accDetail  GetList err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	resp.Data = summs
	this.Data["json"] = resp
	this.ServeJSON()
	return

}
