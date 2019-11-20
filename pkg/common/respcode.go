package common

type RespCode struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ERROR_OK                          = "0000"
	ERROR_USERISEXIST                 = "1001"
	MSG_CODE_ERROR                    = "1002"
	ERROR_PARAM_ISNULL                = "1003"
	USERNAME_OR_PWD_ERROR             = "1004"
	USER_STATUS_ERROR                 = "1005"
	PWD_IS_ERROR                      = "1006"
	BANK_AUTH_ERROR                   = "1007"
	PARTNER_GARDE_NOTEXIST            = "1008"
	BALANCE_NOT_ENOUGH                = "1009"
	SESSION_ERROR                     = "1010"
	SENDMSG_OUT_OF_LIMIT              = "1011"
	ERROR_USERISNOTEXIST              = "1012"
	UPDATE_NICKNAME_ERROR             = "1013"
	ERROR_SYSTEM_ERR                  = "9999"
	MSG_OTC_SUCCESS                   = "2000"
	ERROR_OTC_PRICE                   = "2001"
	ERROR_OTC_INPUT                   = "2002"
	ERROR_CC_INPUT                    = "2002"
	ERROR_OTC_SYMBOLEXIST             = "2003"
	ERROR_OTC_COIN_NOT_ENGHONE        = "2004" //币量不足
	ERROR_OTC_MENU_FINISH             = "2005" //订单已经完成
	ERROR_OTC_MENU_PAY                = "2006" //订单已经支付
	ERROR_OTC_MENU_RECEIVABLES        = "2007" //订单已收款
	ERROR_OTC_MENU_NOT_PAY            = "2008" //订单待支付
	ERROR_OTC_MENU_NOT_COINMONEY      = "2009" //订单待打币
	ERROR_OTC_MENU_ERROR              = "2010" //订单错误
	ERROR_OTC_MENU_SUCCESS            = "2011" //订单取消成功
	ERROR_OTC_MENU_ERROR1             = "2015" //取消订单输入错误
	ERROR_OTC_MENU_BUY_CONFIRM_ERROR  = "2012" //卖单持有者不能确认已付款
	ERROR_OTC_MENU_SELL_CONFIRM_ERROR = "2013" //买单持有者不能确认收款
	ERROR_OTC_MENU_COIN_LIMIT         = "2014" //币量小于限额
	ERR_COMMON_INPUT                  = "8001" //输出参数错误
	ERROR_REPORT_ERROR                = "3001" //用户撤销举报订单超时
	ERROR_CC_BLANCE_NOT_ENGHONE       = "4101" //余额（币量）不足
	ERROR_CC_ORDER_NOTEXIST           = "4102" //订单不存在
	ERROR_REDIS_LOCK_EXIST            = "9001" //锁已存在
)

var (
	error_map map[string]string
)

func init() {
	initErrMap()
}

func GetErrMsg(errorCode string) string {

	strRet, ok := error_map[errorCode]
	if ok == true {
		return strRet
	}
	return "unknown error"
}

//MakeResp ...
func MakeResp(errCode string) RespCode {
	respCode := RespCode{Code: errCode, Msg: GetErrMsg(errCode)}
	return respCode
}

//MakeResp2 ...
func MakeResp2(errCode string, data interface{}) RespCode {
	respCode := RespCode{Code: errCode, Msg: GetErrMsg(errCode), Data: data}
	return respCode
}

//MakeResp3 ...
func MakeResp3(errCode string, errMsg string) RespCode {
	respCode := RespCode{Code: errCode, Msg: errMsg}
	return respCode
}
func initErrMap() {
	if error_map == nil {
		error_map = make(map[string]string)
	}

	error_map[ERROR_OK] = "ok"
	error_map[ERROR_USERISEXIST] = "用户已存在"
	error_map[ERROR_USERISNOTEXIST] = "用户不存在"
	error_map[ERROR_SYSTEM_ERR] = "系统错误"
	error_map[MSG_CODE_ERROR] = "消息编码错"
	error_map[ERROR_PARAM_ISNULL] = "请求参数不能为空"

	error_map[USERNAME_OR_PWD_ERROR] = "用户名密码错"
	error_map[USER_STATUS_ERROR] = "用户状态错"
	error_map[PWD_IS_ERROR] = "密码错"
	error_map[BANK_AUTH_ERROR] = "银行验证错"
	error_map[PARTNER_GARDE_NOTEXIST] = "会员等级配置不存在"
	error_map[BALANCE_NOT_ENOUGH] = "用户钱包余额不足"

	error_map[ERROR_OTC_SYMBOLEXIST] = "don't exsit symbol"

	//2xxx
	error_map[MSG_OTC_SUCCESS] = "Suspension of success"
	error_map[ERROR_OTC_PRICE] = "get price is error"
	error_map[ERROR_OTC_INPUT] = "input params is null"
	error_map[ERROR_OTC_COIN_NOT_ENGHONE] = "insufficient currency"                          //币量不足
	error_map[ERROR_OTC_MENU_FINISH] = "menu finish"                                         //订单已经完成
	error_map[ERROR_OTC_MENU_PAY] = "payment of orders"                                      //订单已支付
	error_map[ERROR_OTC_MENU_RECEIVABLES] = "order receivable"                               //订单已收款
	error_map[ERROR_OTC_MENU_NOT_PAY] = "order pending"                                      //订单待支付
	error_map[ERROR_OTC_MENU_NOT_COINMONEY] = "order to be coined"                           //订单待打币
	error_map[SESSION_ERROR] = "session异常"                                                   //session异常
	error_map[SENDMSG_OUT_OF_LIMIT] = "短信发送已达当日上限"                                           //短信发送已达当日上限
	error_map[ERROR_OTC_MENU_ERROR] = "Not matching orders"                                  //订单类型错误
	error_map[ERROR_OTC_MENU_SUCCESS] = "Order cancelled successfully"                       //订单取消成功
	error_map[ERROR_OTC_MENU_BUY_CONFIRM_ERROR] = "Sell order holder cannot confirm payment" //卖单持有者不能确认已付款
	error_map[ERROR_OTC_MENU_SELL_CONFIRM_ERROR] = "The buyer cannot confirm the receipt"    //卖家确认订单错误
	error_map[ERROR_OTC_MENU_COIN_LIMIT] = "The input currency is less than the limit"       //输入币量小于限额
	error_map[ERROR_REPORT_ERROR] = "User cancels report order over time"                    //用户撤销举报订单超时
	error_map[ERROR_OTC_MENU_ERROR1] = "order has been cancelled"

	//4xxx
	error_map[ERROR_CC_BLANCE_NOT_ENGHONE] = "币币账户余额不足"
	error_map[ERROR_CC_ORDER_NOTEXIST] = "币币交易订单不存在"

	// 8xxx
	error_map[ERR_COMMON_INPUT] = "输出参数错误"
	error_map[UPDATE_NICKNAME_ERROR] = "修改昵称失败，距离上次修改时间还不到一个月"

	// 9xxx
	error_map[ERROR_REDIS_LOCK_EXIST] = "锁已存在"

}
