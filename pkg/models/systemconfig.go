package models

import (
	"strconv"

	"github.com/rehylas/wx/pkg/utils"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2/bson"
)

// 系统配置表
type SystemConfig struct {
	Id         string `json:"id"`
	ConfigName string `json:"configname"` //配置名称
	ConfigCode string `json:"configcode"`

	AndroidVersion       string  `json:"androidversion"`       //安卓版本号
	IosVersion           string  `json:"iosversion"`           //ios版本号
	AndroidLowestVersion string  `json:"androidlowestversion"` //安卓最低版本号
	IosLowestVersion     string  `json:"ioslowestversion"`     //ios最低版本号
	AndroidUrl           string  `json:"androidurl"`           //安卓下载链接
	IosUrl               string  `json:"iosurl"`               //ios下载链接
	UnfreezeFeeV1        float64 `json:"unfreezefeev1"`        //v1等级解封费用
	UnfreezeFeeV2        float64 `json:"unfreezefeev2"`        //v2等级解封费用
	UnfreezeFeeV3        float64 `json:"unfreezefeev3"`        //v3等级解封费用
	UnlockCount          int64   `json:"unlockcount"`          //次 每月封号超出次数永久封号
	UnlockHour           int64   `json:"unlockhour"`           //小时 临时封号需要等待时间后自动解封
	InvitationLink       string  `json:"invitationlink"`       //邀请链接
	SendMsgMaxcnt        int64   `json:"sendmsgmaxcnt"`        //当日短信发送最大条数

	TradeWeekday   string `json:"tradeweekday"`   //交易周期 星期
	TradeStartTime string `json:"tradestarttime"` //交易开始时间
	TradeEndTime   string `json:"tradeendtime"`   //交易结束时间
	TradeTimeout   int64  `json:"tradetimeout"`   //交易超时取消时间(分钟),匹配后必须在指定时间内完成交易
	TradeMin       int64  `json:"trademin"`       //交易数量区间,交易最小值
	TradeMax       int64  `json:"trademax"`       //交易数量区间,交易最大值
	Val            string `json:"val"`            //参数值

	Modifier   string `json:"modifier" `   //修改人
	UpdateTime string `json:"updatetime" ` //修改时间
}

func (sc *SystemConfig) Insert() error {
	sc.Id = utils.Krand(18, 0)
	collect, err := getCollect(COLLECTNAME_SYSCONFIG)
	if err != nil {
		return err
	}
	err = collect.Insert(sc)
	if err != nil {
		return err
	}
	return nil
}

func (sc *SystemConfig) GetConfigByCode() error {
	collect, err := getCollect(COLLECTNAME_SYSCONFIG)
	if err != nil {
		return err
	}
	whereSql := bson.M{"configcode": sc.ConfigCode}
	err = collect.Find(whereSql).One(sc)
	return err
}

func (sc *SystemConfig) UpdateConfig() (err error) {
	collect, err := getCollect(COLLECTNAME_SYSCONFIG)
	if err != nil {
		return err
	}
	setSql := make(map[string]interface{})
	setSql["id"] = sc.Id
	err = collect.Update(setSql, sc)
	return err
}

func GetConfigByDBString(configName string, defval string) string {
	sc := SystemConfig{ConfigCode: configName}
	err := sc.GetConfigByCode()
	if err != nil {
		beego.Error("GetConfigByCode err:", err)
		return defval
	}
	return sc.Val

}

func GetConfigByDBInt(configName string, defval int) int {
	sc := SystemConfig{ConfigCode: configName}
	err := sc.GetConfigByCode()
	if err != nil {
		beego.Error("GetConfigByCode err:", err)
		return defval
	}
	nval, err := strconv.Atoi(sc.Val)
	if err != nil {
		return defval
	}
	return nval
}

func GetConfigByDBFloat(configName string, defval float64) float64 {
	sc := SystemConfig{ConfigCode: configName}
	err := sc.GetConfigByCode()
	if err != nil {
		beego.Error("GetConfigByCode err:", err)
		return defval
	}

	beego.Debug(sc)

	fval, err := strconv.ParseFloat(sc.Val, 64)
	if err != nil {
		return defval
	}
	return fval
}
