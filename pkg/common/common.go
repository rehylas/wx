package common

import "time"

const (
	VERSION = "0.9"
)

//数据库中的系统配置
const (
	CFG_TRADE_TAKEALL     = "trade_takeall"     //是否必须全匹配 1 全比配
	CFG_TRADE_BS_MYSELF   = "trade_bs_myself"   //是否允许自我交易 1 为允许
	CFG_TRADE_HAS_BALANCE = "trade_has_balance" //交易最小持币量
)

//发送短信相关的key
const (
	TMP_ID_MSG         = "193218"        // 【JUHOO】您的验证码是#code#。如非本人操作，请忽略本短信
	TMP_ID_PAYMENT_MSG = "111"           // 【JUHOO】您正在修改支付方式，如非本人操作，请登录app修改密码。
	SEND_MSG_KEY       = "SEND_MSG_KEY_" //注册发送短信验证码
	SEND_MSG_MAXCNT    = 20
	PR_SEND_MSG_MAXCNT = "PR_SEND_MSG_MAXCNT"
	// 短信发送当日最大次数20
)

const (
	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	TIME_TMP1 = "2006-01-02 15:04:05" //常规类型
	TIME_TMP2 = "2006/01/02 15:04:05" //其他类型
	TIME_TMP3 = "2006-01-02"          //日期标准化
	TIME_TMP4 = "15:04:05"            //时间标准化
)

// LOGIN_TYPE 登录类型定义
const (
	LOGIN_TYPE_ACCOUNT     = 1 //账号密码登录
	LOGIN_TYPE_MSG         = 2 //手机号短信验证码登录
	LOGIN_TYPE_FINGERPRINT = 3 //指纹密码登录
)

// UserRole.Rolecode 角色code
// const (
// 	PARTNER_GRADE         = "PARTNER_GRADE"         //合伙人等级配置code
// 	GROUP_MANAGER_GRADE   = "GROUP_MANAGER_GRADE"   //群管等级配置code
// 	COMPANY_MANAGER_CRADE = "COMPANY_MANAGER_GRADE" //高管等级配置code
// )

// 系统账户
const (
	SYSTEMACCOUNT       = "systemcommissionfee"  //系统手续费账户名
	SYS_DESTROY_ACCOUNT = "systemdestroyaccount" //公司销毁账户
	HANDLING_FEE        = 5.0                    //手续费
)

//转账类型
const (
	TRANSACTION_HANDLING_FEE = 1 //手续费转账
	TRANSACTION_TRANSFER_FEE = 2 //交易转账
	TRANSACTION_DIVIDEND_FEE = 3 //分红转账
	TRANSACTION_PERSON_FEE   = 4 //个人转账
)

//TimeTaskLog.Title 定时任务title
// const (
// 	BONUS_TITLE_PARTNER         = "BONUS_TITLE_PARTNER"         //合伙人分红title
// 	BONUS_TITLE_GROUP_MANAGER   = "BONUS_TITLE_GROUP_MANAGER"   //群管分红title
// 	BONUS_TITLE_COMPANY_MANAGER = "BONUS_TITLE_COMPANY_MANAGER" //高管分红title
// 	BONUS_TITLE_COMMEND         = "BONUS_TITLE_COMMEND"         //推荐人分红title
// 	BONUS_TITLE_DESTROY         = "BONUS_TITLE_DESTROY"         //销毁分红title
// 	BONUS_TITLE_MALL_BALANCE    = "BONUS_TITLE_MALL_BALANCE"    //商城账户冻结分红title
// )

const (
	BEGIN = "BEGIN" //开始
	END   = "END"   //结束
	ERROR = "ERROR" //异常
)

const (
	SESSION_NAME = "AutoFTSessionUser"
)

// 消息 数据类型定义
const (
	MSG_DT_NewOTCTrade = "newOtcTrade" // 新的交易形成
)

//OTC 数据周期定义
const (
	OTC_PERIOD_1MIN = "1min"
	OTC_PERIOD_1DAY = "1day"
)

//内容管理枚举ContentManage.Type
const (
	CM_BANNER    = "BANNER"    //banner
	CM_NOTICE    = "NOTICE"    //公告
	CM_DYNAMIC   = "DYNAMIC"   //动态
	CM_PROPAGATE = "PROPAGATE" //宣传
	CM_INVITE    = "INVITE"    //邀请海报
)

//用户收款方式类型：users表
const (
	PAYMENT_ALIPAY = "ALIPAY" //支付宝
	PAYMENT_BANK   = "BANK"   //银行卡
	PAYMENT_WX     = "WEIXIN" //微信
	PAYMENT_USDT   = "USDT"   //usdt
)

// 邀请链接
const INVITATION_LINK = "http://wwww.xxx.com:8888/register/index.html?inviteCode="
const PR_INVITATION_LINK = "PR_INVITATION_LINK" //redis前缀

func BuildMsgKey(key string, mobile string) string {
	return key + mobile
}

//redis 分布式锁key
const (
	Key    = "SYSTEMLOCK"
	EXPIRE = time.Second
)
