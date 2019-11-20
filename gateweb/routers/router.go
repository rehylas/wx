// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"

	"../controllers"
)

func init() {
	beego.Info("route.go init")
	//ping
	beego.Router("/ping", &controllers.PingController{}, "post:Ping")
	beego.Router("/dbdemo", &controllers.PingController{}, "post:DBDemo")

	// wx服务  WxController
	beego.Router("/receiver", &controllers.WxController{}, "get:Receiver")
	beego.Router("/receiver", &controllers.WxController{}, "post:Receiver")
	beego.Router("/TestWx2/ServletApi", &controllers.WxController{}, "post:Receiver")
	beego.Router("/wx/receiver", &controllers.WxController{}, "post:Receiver")
	beego.Router("/wx/sendtempmsg", &controllers.WxController{}, "post:Sendtempmsg")
	beego.Router("/DXBizGate/server", &controllers.WxController{}, "post:Sendtempmsg") //兼容老服务

}
