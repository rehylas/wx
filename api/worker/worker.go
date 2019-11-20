package worker

import (
	"time"

	"github.com/rehylas/autoft/aft/market"
	"github.com/rehylas/autoft/pkg/models"

	"github.com/astaxie/beego"
)

var (
	Modelist []models.Mode
	Market   *market.MarketRedis
)

/*********************************************************************

该模块对外只输出  StartWork() ，其它业务均在内部完成
该模块现处理业务有：
1. 读取数据库模型配置，根据不同的模型类型，加载不同的工作实例

**********************************************************************/

type WorkThread struct {
	Name    string
	Mode    models.Mode
	tick_ch chan models.Tick
	done_ch chan int
}

type WorkThreadInterface interface {
	Run()
	Stop()
	GetTickChan() chan models.Tick
}

// ID       string `json:"_id"`
// Symbol   string `json:"symbol"`
// ModeType string `json:"modetype"`
// BsType   string `json:"bstype"`
// Vol      int    `json:"vol"`
// VolDef   int    `json:"voldef"`
// Enable   int    `json:"enable"` // 0  不打开   1 打开
// Exec     int    `json:"exec"`   // 0  不执行   1 默认执行   2 人工执行； 如果不执行则不委托， 如果默认执行，则读取def下单， 如果人工执行则读取vol下单
// State    int    `json:"state"`  // 状态 0 等待   1 进入  2 出局
// Execdt   string `json:"execdt"`

var works []WorkThreadInterface

func StartWork() {
	// wt_demo := WorkThread_demo{}

	// wt_demo.Run()

	wt_summ := WorkThread_summ{}

	go wt_summ.Run()

}

func StopWork() {

}

//基本工作线程
func (this *WorkThread) Run() {
	for true {
		// if eng.runState == RUN_STATE_PAUSE {
		// 	time.Sleep(time.Microsecond * 100)
		// 	continue
		// }

		select {
		case tick := <-this.tick_ch:
			{
				beego.Info("WorkThread:", tick)
			}
		case <-this.done_ch:
			{
				return
			}
		case <-time.After(time.Microsecond * 500):
			{
				//beego.Info("time out 500 ms")
			}
		}
	}
}

func (this *WorkThread) Stop() {
	this.done_ch <- 0
}

func (this *WorkThread) GetTickChan() chan models.Tick {
	return this.tick_ch
}
