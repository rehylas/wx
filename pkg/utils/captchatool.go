package utils

import (
	"encoding/json"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

// 图片验证码
type Captcha struct {
	CaptchaId string
	//CaptchaType     string
	VerifyValue string
	//ConfigAudio     base64Captcha.ConfigAudio
	//ConfigCharacter base64Captcha.ConfigCharacter
	//ConfigDigit     base64Captcha.ConfigDigit
}

//数字验证码配置
var configA = base64Captcha.ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 5,
}

//字符,公式,验证码配置
var configB = base64Captcha.ConfigCharacter{
	Height:             60,
	Width:              240,
	Mode:               0,
	ComplexOfNoiseText: 0,
	ComplexOfNoiseDot:  0,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         6,
}

// base64Captcha create http handler
func CreateCaptchaHandler(w http.ResponseWriter, r *http.Request) *Captcha {
	//parse request parameters
	//接收客户端发送来的请求参数
	decoder := json.NewDecoder(r.Body)
	var postParameters Captcha
	err := decoder.Decode(&postParameters)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha(postParameters.CaptchaId, configA)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	postParameters.CaptchaId = captchaId
	postParameters.VerifyValue = base64Png

	return &postParameters

	//or you can do this
	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//set json response
	//设置json响应

	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//body := map[string]interface{}{"code": 1, "data": base64Png, "captchaId": captchaId, "msg": "success"}
	//json.NewEncoder(w).Encode(body)
}

// base64Captcha verify http handler
func VerifyCaptchaHandle(w http.ResponseWriter, r *http.Request) bool {
	//parse request parameters
	//接收客户端发送来的请求参数
	decoder := json.NewDecoder(r.Body)
	var postParameters Captcha
	err := decoder.Decode(&postParameters)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	//比较图像验证码
	verifyResult := base64Captcha.VerifyCaptcha(postParameters.CaptchaId, postParameters.VerifyValue)

	//set json response
	//设置json响应
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//body := map[string]interface{}{"code": "error", "data": "验证失败", "msg": "captcha failed"}
	//if verifyResult {
	//	body = map[string]interface{}{"code": "success", "data": "验证通过", "msg": "captcha verified"}
	//}
	//json.NewEncoder(w).Encode(body)
	return verifyResult
}
