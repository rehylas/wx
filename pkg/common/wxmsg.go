package common

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/rehylas/wx/pkg/utils"
)

const SM_URL = "http://www.bangnikanzhe.com/DXBizGate/server?cmd=sendmsg4client"

func SendWxMsg(msg string) error {
	//     data = {'userid':'80010121','userpwd':'888889','txt': msg }
	//     res = stockMsgSend( data )
	msgData := fmt.Sprintf(`{ "userid":"80010121","userpwd":"888889","txt": "%s" }`, msg)
	ret, err := utils.HttpPostJson(SM_URL, msgData)
	if err != nil {
		beego.Error("SendWxMsg HttpPostJson err:", err)
	}
	beego.Info("SendWxMsg HttpPostJson ret:", ret)
	return err
}

func SendWxMsgThread(msg string) {
	go SendWxMsg(msg)
}

// {'userid':'80010121','userpwd':'888888','txt':'002230 科大讯飞要涨了' }

// def sendwxmsg( msg ):
//     data = {'userid':'80010121','userpwd':'888889','txt': msg }
//     res = stockMsgSend( data )
//     pass
