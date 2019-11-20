package main

import (
	"fmt"

	"github.com/astaxie/beego"

	"./database"
	"./models"
	_ "./routers"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/wx.log"}`)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Informational()
	server := fmt.Sprintf(":%s", beego.AppConfig.String("httpport"))

	// 方案一
	// 注册处理函数
	// beego.Informational("start in ", server)
	// http.HandleFunc("/receiver", biz.Receiver)
	// log.Fatal(http.ListenAndServe(server, nil))

	// 方案二
	dbCfg, err := beego.AppConfig.GetSection("database")
	if err != nil {
		beego.Error("Initdb error:", err)
		return
	}

	beego.Info("", dbCfg["mongohost"], dbCfg["redishost"])
	err = database.Initdb(dbCfg["mongohost"], dbCfg["redishost"])
	if err != nil {
		beego.Error("Initdb error:", err)
		return
	}

	beego.SetStaticPath("/", "web")

	test()
	beego.Informational("start in ", server)
	beego.Run()

}

func test() {
	var user models.CommUser
	beego.Debug("recs:", user.Userid)
	(&user).GetCommUserById("80010121")

	//biz.CheckMsgSend("testopenid")
}
