package worker

import (
	"time"

	"github.com/rehylas/wx/pkg/utils"
)

type WorkThread_demo struct {
	WorkThread
}

//工作线程
func (this *WorkThread_demo) Run() {

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
func (this *WorkThread_demo) ProTime() {
	timeStr := utils.TimeStr()
	if timeStr == "23:00:00" {
		// do something

	} //
}
