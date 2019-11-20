package common

type RespCode struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ERROR_OK            = "0000"
	ERROR_USERISEXIST   = "1001"
	ERROR_USERIDPWD_ERR = "1002"
	ERROR_USER_STATE    = "1003"

	ERROR_MAXMSG_SEND = "7001"
	ERR_SYSTEM_ERR    = "9999"
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

func MakeResp(errCode string) RespCode {
	respCode := RespCode{Code: errCode, Msg: GetErrMsg(errCode)}
	return respCode
}

func initErrMap() {
	if error_map == nil {
		error_map = make(map[string]string)
	}

	error_map[ERROR_OK] = "ok"
	error_map[ERROR_USERISEXIST] = "user is exist"
	error_map[ERR_SYSTEM_ERR] = "system is error"
	error_map[ERROR_USERIDPWD_ERR] = "userid or user pwd error"
	error_map[ERROR_USER_STATE] = "user state is error"

	error_map[ERROR_MAXMSG_SEND] = "send msgs number  than day max "

}
