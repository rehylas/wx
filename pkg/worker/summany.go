package worker

import (
	"time"

	"github.com/rehylas/wx/pkg/models"
)

type SummanyThread struct {
	WorkThread
}

func (this *SummanyThread) Run() {
	go this.work()
}

//////////////////////////////////////////////////////////////
// 该单元统计准实时数据  ， 注册用户数， 当日交易量等
func (this *SummanyThread) work() {
	for {

		this.summanydo()
		time.Sleep(time.Second * 60)

	}
}

func (this *SummanyThread) summanydo() {
	//
	totals, pays := models.SummUser()
	sum := models.Summ{Totalusers: totals, Payusers: pays}
	sum.AddRec()
}
