package main

import (
	"encoding/gob"
	"flag"
	"log"
	"net/http"

	_ "github.com/rehylas/wx/api/routers"
	"github.com/rehylas/wx/api/worker"

	// "github.com/rehylas/wx/pkg/cache"
	"github.com/rehylas/wx/pkg/database"
	"github.com/rehylas/wx/pkg/models"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

var (
	serverIsInit bool
)

func init() {

	//解决跨域问题

	serverIsInit = false
	beego.SetLogger("file", `{"filename":"logs/api.log"}`)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		Pprof()
	} else {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.SetLevel(beego.LevelWarning)
	}

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
	err = database.InitTransactionDB(dbCfg["mongohost"])
	if err != nil {
		beego.Error("InitTransactionDB error:", err)
		return
	}

	//设置错误页
	beego.ErrorHandler("404", page_not_found)

	initSession()
	//设置静态资源访问
	setFileStaticPath()

	serverIsInit = true
}

func initSession() { //beego的session序列号是用gob的方式，因此需要将注册models.User

	gob.Register(models.Admin{}) //https://beego.me/docs/mvc/controller/session.md
}

func setFileStaticPath() {
	beego.SetStaticPath("/ui", "ui")

	// //文件上传路径
	// beego.SetStaticPath("/fileupload", "fileupload")
	// beego.SetStaticPath("/register", "register")

	// //注册页
	// beego.BConfig.WebConfig.StaticDir["/register"] = "register"
	// beego.BConfig.WebConfig.StaticDir["/static"] = "register/static"

}

func page_not_found(rw http.ResponseWriter, r *http.Request) {

	rw.Write([]byte("访问地址不存在"))

}

// Pprof pprof
func Pprof() {
	flag.Parse()
	//这里实现了远程获取pprof数据的接口
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
}

func test() {
	// models.FixAccDetails(100)
	// amount := models.GetConfigByDBFloat("myamount", 10000.0)
	// beego.Debug("amount:", amount)
	v1, v2 := models.SummUser()
	beego.Debug("", v1, v2)
}

func main() {

	if serverIsInit == false {
		beego.Error("server init error")
		return
	}
	test()

	// 启动一些工作线程
	worker.StartWork()

	beego.Run()

}
