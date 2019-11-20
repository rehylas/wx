package utils

import (
	"fmt"
	"math"
)

const SHA256_RAND_DATA = ""

/******************************************
此文件定义了业务相关的 加解密操作
输出：
CreateBlanceCheckVal   //生成余额加密值
ChkeckBlanceCheckVal   //校验余额加密值
*******************************************/

func CreateBlanceCheckVal(username string, blance float64) string {

	val := math.Trunc(blance*1e2+0.5) * 1e-8 //小数点后取8位
	data := fmt.Sprintf("%%s%.8f", SHA256_RAND_DATA, username, val)
	return Sha256(data)
}

func ChkeckBlanceCheckVal(username string, blance float64, checkval string) bool {
	if CreateBlanceCheckVal(username, blance) == checkval {
		return true
	} else {
		return false
	}

}
