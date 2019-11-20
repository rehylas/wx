package service

import (
	"github.com/astaxie/beego"
	"github.com/rehylas/wx/pkg/models"
)

type AccDetail struct {
}

//
func (this *AccDetail) GetList(accdetails *[]models.AccDetail) error {

	amount := models.GetConfigByDBFloat("myamount", 10000.0)
	beego.Debug("amount:", amount)
	models.FixAccDetails(amount)

	err := models.GetAccDetails(accdetails)
	if err != nil {
		beego.Error("mode.GetModes() err:", err)
		return err
	}
	return nil
}
