package worker

/*********************************************************************
该模块实现后台处理的工作线程，伴随着 api网关  一起启动, 工作线程需要注意防止多
次处理 api网关  负载均衡多实例的
该模块对外只输出  StartWork() ，其它业务均在内部完成
该模块现处理业务有：
1. 根据撮合引擎反馈的filled order 信息进行数据落地
2. 根据撮合引擎反馈的 trade, 信息进行k线形成，最新价形成
**********************************************************************/

// const ME_TRADE_LIST = "me_trade_list"
// const ME_ORDER_LIST = "me_order_list"

type WorkThread struct {
	Name string
}

func StartWork() {
	// workCCorderfilled := CCOrderFilledThread{}
	// workCCTrade2k := CCTrade2kThread{}
	summanyThread := SummanyThread{}

	// workCCorderfilled.Run()
	// workCCTrade2k.Run()
	summanyThread.Run()

}
