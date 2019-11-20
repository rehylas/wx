package controllers

import (
	"encoding/json"
	"gtc/common"
	"gtc/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type PingController struct {
	beego.Controller
}

type ReqDataPing struct {
	UserName string `json:"username"`
	Password string `json:"Password"`
}

// @Title Login
// @Description Logs user into the system
// @Param	body    {"username":"name_zhangshan","age":1}
// @Success 0000 {string} ping  success
// @Failure 9999
// @router /ping  [post]
func (u *PingController) Ping() {
	var pingData ReqDataPing
	var user models.User

	json.Unmarshal(u.Ctx.Input.RequestBody, &pingData)
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	beego.Info("data:", pingData)
	beego.Info("user:", user)

	resp := common.MakeResp(common.ERROR_OK)
	resp.Data = pingData
	u.Data["json"] = resp

	u.ServeJSON()
}

func (u *PingController) DBDemo() {

	//增删改查
	user := &models.User{Username: "namezhang"}
	user.AddUser()

	models.DeleteUser("user_1564130052466842800")

	user.Username = "王"
	// user.UpdateUser("user_1564130115735887800")

	user.GetUser("user_1564130115735887800")
	beego.Info("user:", *user)

	users := []models.User{}
	models.GetAllUsers(&users)
	beego.Info("user:", users)

	coins := []models.Coin{}
	models.GetAllCoins(&coins)
	beego.Info("coins:", coins)

	// end

	resp := common.MakeResp(common.ERROR_OK)
	u.Data["json"] = resp
	u.ServeJSON()

}
