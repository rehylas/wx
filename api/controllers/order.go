package controllers

import (
	"github.com/astaxie/beego"

	"github.com/rehylas/wx/pkg/common"
	"github.com/rehylas/wx/pkg/models"
)

// Operations about Mode
type OrderController struct {
	BaseController
}

// @Title List
// @Description 交易
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /order/list/ [get]
func (this *OrderController) List() {

	var orders []models.Order
	err := models.GetOrdersAll(&orders)

	if err != nil {
		beego.Error("mode.GetModes() err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	resp.Data = orders
	this.Data["json"] = resp
	this.ServeJSON()
	return

}

// @Title List
// @Description 交易
// @Success 0000 {object} models.User
// @Failure 9999 system is error
// @router /order/listtoday/ [get]
func (this *OrderController) Listtoday() {

	var orders []models.Order
	err := models.GetOrderToday(&orders)

	if err != nil {
		beego.Error("mode.GetModes() err:", err)
		resp := common.MakeResp(common.ERROR_SYSTEM_ERR)
		this.Data["json"] = resp
		this.ServeJSON()
		return
	}

	resp := common.MakeResp(common.ERROR_OK)
	resp.Data = orders
	this.Data["json"] = resp
	this.ServeJSON()
	return

}
