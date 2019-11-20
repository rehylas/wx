package worker

import (
	"time"

	"github.com/astaxie/beego"

	"github.com/rehylas/wx/pkg/models"
	"github.com/rehylas/wx/pkg/utils"
)

type WorkThread_summ struct {
	WorkThread
}

//工作线程
//每天 23:00:00 统计一下当天用户数
func (this *WorkThread_summ) Run() {

	for true {

		select {

		case <-this.done_ch:
			{
				return
			}
		case <-time.After(time.Millisecond * 1000):
			{
				this.ProTime()
				//beego.Info("time out 500 ms")
			}
		}
	}
}

// 业务逻辑：
func (this *WorkThread_summ) ProTime() {
	// beego.Debug("protick ")
	timeStr := utils.TimeStr()
	if timeStr == "23:50:00" {
		totals, pays := models.SummUser()
		beego.Debug("totals, pays ", totals, pays)
		sum := models.Summ{Totalusers: totals, Payusers: pays}
		sum.AddRec()

	}

}
