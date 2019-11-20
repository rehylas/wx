package service

import (
	"github.com/astaxie/beego"
	"github.com/rehylas/wx/pkg/models"
)

type SystemInfo struct {
}

func GetSystemInfos(systems *[]models.SystemInfo) error {
	//var systems []models.SystemInfo

	sysinfo_aft := models.SystemInfo{Name: "aft"}
	sysinfo_api := models.SystemInfo{Name: "api"}
	sysinfo_market := models.SystemInfo{Name: "market"}
	sysinfo_trade := models.SystemInfo{Name: "trade"}

	sysinfo_aft.GetSystemByName("aft")
	sysinfo_api.GetSystemByName("api")
	sysinfo_market.GetSystemByName("market")
	sysinfo_trade.GetSystemByName("trade")

	*systems = append(*systems, sysinfo_aft)
	*systems = append(*systems, sysinfo_api)
	*systems = append(*systems, sysinfo_market)
	*systems = append(*systems, sysinfo_trade)

	beego.Debug("systems:", systems)

	//systems = &systems
	return nil
}
