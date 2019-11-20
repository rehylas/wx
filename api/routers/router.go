// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/rehylas/wx/api/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func initSwagger() {

}

func init() {

	if beego.BConfig.RunMode == "dev" {
		beego.Info("initSwagger ")
		initSwagger()
	}

	//解决跨域问题
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	//后台管理
	beego.Router("/v1/admin/login", &controllers.AdminController{}, "post:Login")

	//添加Mode管理
	beego.Router("/v1/mode/list", &controllers.ModeController{}, "get:List")
	beego.Router("/v1/mode/add", &controllers.ModeController{}, "post:Add")
	beego.Router("/v1/mode/del", &controllers.ModeController{}, "post:Del")
	beego.Router("/v1/mode/update", &controllers.ModeController{}, "post:Update")

	//添加commuser管理
	beego.Router("/v1/commuser/list", &controllers.CommuserController{}, "get:List")
	beego.Router("/v1/commuser/add", &controllers.CommuserController{}, "post:Add")
	beego.Router("/v1/commuser/del", &controllers.CommuserController{}, "post:Del")
	beego.Router("/v1/commuser/update", &controllers.CommuserController{}, "post:Update")

	//添加Order管理
	beego.Router("/v1/order/list", &controllers.OrderController{}, "get:List")
	beego.Router("/v1/order/listtoday", &controllers.OrderController{}, "get:Listtoday")

	//添加账户曲线
	beego.Router("/v1/accdetail/list", &controllers.AccountDetailController{}, "get:List")

	//添加用户统计曲线
	beego.Router("/v1/summ/list", &controllers.SummController{}, "get:List")

	//系统信息
	beego.Router("/v1/systeminfo/list", &controllers.SystemInfoController{}, "get:List")
	beego.Router("/v1/systeminfo/update", &controllers.SystemInfoController{}, "get:Update")

}
