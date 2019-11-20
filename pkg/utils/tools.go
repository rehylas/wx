package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	rd "math/rand"

	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/valyala/fasthttp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
// c    = fasthttp.Client{}
// req  = fasthttp.AcquireRequest()
// resp = fasthttp.AcquireResponse()
)

//RandString 生成随机字符串
func RandString() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//Str2Time 将"2016-02-15 12:00:00"或者"2016-04-18 09:33:56.694"等格式转化为time.Time
func Str2Time(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	return t, err
}

//Intface2Str 将interface{}类型转为string类型
func Intface2Str(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case float64:
		return strconv.FormatFloat(v, 'f', 6, 64), nil
	case bool:
		return strconv.FormatBool(v), nil
	default:
		return "", errors.New("invalid interface type")
	}
}

//测量一段代码执行时间
func TraceCode() func() {
	start := time.Now()
	return func() {
		t := time.Now().Sub(start).Nanoseconds()
		fmt.Printf("运行耗时:%d(纳秒)\n", t)
	}
}

//PrintStack 打印当前堆栈
func PrintStack(all bool) {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, all)
	os.Stdout.Write(buf[:n])
}

//GetStack 获取当前堆栈
func GetStack(all bool) []byte {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, all)
	return buf[:n]
}

// zero-copy, []byte转为string类型
// 注意，这种做法下，一旦[]byte变化，string也会变化
// 谨慎，黑科技！！除非性能瓶颈，否则请使用string(b)
func Bytes2String(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}

// zero-coy, string类型转为[]byte
// 注意，这种做法下，一旦string变化，程序立马崩溃且不能recover
// 谨慎，黑科技！！除非性能瓶颈，否则请使用[]byte(s)
func String2Bytes(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&s))
}

//IsEOF 判断一个error是否是io.EOF
func IsEOF(err error) bool {
	if err == nil {
		return false
	} else if err == io.EOF {
		return true
	} else if oerr, ok := err.(*net.OpError); ok {
		if oerr.Err.Error() == "use of closed network connection" {
			return true
		}
	} else {
		if err.Error() == "use of closed network connection" {
			return true
		}
	}
	return true
}

//LocalIP 获取本机ip
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if !strings.Contains(ipnet.IP.String(), "192.168") {
					return ipnet.IP.String()
				}
			}
		}
	}
	return ""
}

//获取远程ip
func RemoteIP(remoteAddr string) string {
	var ipAddr string
	if ip, _, err := net.SplitHostPort(remoteAddr); err == nil {
		ipAddr = ip
		fmt.Println("ip:", ip)
	} else {
		ipAddr = remoteAddr
	}
	return ipAddr
}

// Krand 随机字符串 , 0 数字 1 小写, 2 大写   3数字+大小写
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rd.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rd.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rd.Intn(scope))
	}
	bytestr := fmt.Sprintf("%d", time.Now().Year())[2:]
	return string(bytestr) + string(result)
}

//Cfg 读取配置文件
func Cfg(name string) (map[string]string, error) {
	dbCfg, err := beego.AppConfig.GetSection(name)
	if err != nil {
		beego.Error("Initdb error:", err)
		return nil, err
	}
	return dbCfg, nil
}

// get 网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		logs.Error("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		logs.Error("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		logs.Error("短信post请求失败:\r\n%v", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// // FastGet 网络请求
// func FastGet(apiURL string, params url.Values) (rs []byte, err error) {
// 	var Url *url.URL
// 	Url, err = url.Parse(apiURL)
// 	if err != nil {
// 		logs.Error("解析url错误:\r\n%v", err)
// 		return nil, err
// 	}
// 	//如果参数中有中文参数,这个方法会进行URLEncode
// 	Url.RawQuery = params.Encode()
// 	_, body, err := c.Get(nil, Url.String())
// 	if err != nil {
// 		logs.Error("err:", err)
// 		return nil, err
// 	}
// 	return body, nil
// }

// // FastPost 网络请求 ,params 是url.Values类型
// func FastPost(apiURL string, params url.Values) (rs []byte, err error) {
// 	arg := parseFormat(params)
// 	_, body, err := c.Post(nil, apiURL, &arg)
// 	if err != nil {
// 		logs.Error("短信post请求失败:\r\n%v", err)
// 		return nil, err
// 	}
// 	return body, nil
// }

func parseFormat(params url.Values) fasthttp.Args {
	var arg = fasthttp.Args{}
	str := params.Encode()
	for _, v := range strings.Split(str, "&") {
		args := strings.Split(v, "=")
		arg.Set(args[0], args[1])
	}
	return arg
}

// 获取一天的开始时间 2006-01-02 00:00:00
func GetDayStartTime(businessTime string) string {
	if businessTime == "" {
		return time.Now().Format("2006-01-02") + " 00:00:00"
	} else {
		return businessTime + " 00:00:00"
	}
}

// 获取一天的结束时间 2006-01-02 23:59:59
func GetDayEndTime(businessTime string) string {
	if businessTime == "" {
		return time.Now().Format("15:04:05") + " 23:59:59"
	} else {
		return businessTime + " 23:59:59"
	}
}

// GetTime 获取当前时间 格式：xx-xx-xx h:m:s
func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//SetTime 挂单设置时间 3d
func SetTime(n int, layout string) string {
	dd, _ := time.ParseDuration(fmt.Sprintf("%d", n*12) + "h")
	dd1 := time.Now().Add(dd)
	return dd1.Format(layout)
}

//GetBeforeDy 获取前一天时间
func GetBeforeDy() (string, string) {
	k := time.Now()
	d, _ := time.ParseDuration("-24h")
	start := fmt.Sprintf("%s%s", k.Add(d).Format("2006-01-02"), " 00:00:00")
	end := fmt.Sprintf("%s%s", k.Add(d).Format("2006-01-02"), " 23:59:59")
	return start, end
}

//

// 判断结构体是否为空
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

//http client  get url  ,  return body
// func HttpGet(url string, params map[string]string) (string, error) {

// 	paramsStr := ""
// 	if len(params) > 0 {
// 		paramsStr = "?"
// 		for key, val := range params {
// 			paramsStr = paramsStr + fmt.Sprintf("%s=%s", key, val) + "&"
// 		}
// 		paramsStr = paramsStr[0 : len(paramsStr)-1]
// 	}
// 	resp, err := http.Get(url + paramsStr)
// 	if err != nil {
// 		return "", err
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	//fmt.Println(string(body))
// 	return string(body), nil

// }

//http client post json , return body
// func HttpPostJson(url string, data string) (string, error) {

// 	reader := bytes.NewReader([]byte(data))
// 	request, err := http.NewRequest("POST", url, reader)
// 	if err != nil {
// 		return "", err
// 	}
// 	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
// 	client := http.Client{}
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		return "", err

// 	}
// 	defer resp.Body.Close()

// 	respBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err

// 	}
// 	//byte数组直接转成string，优化内存  有必要 ???
// 	// str := (*string)(unsafe.Pointer(&respBytes))
// 	// return *str, nil

// 	return string(respBytes), nil

// }

//return  now datetime  2019-12-30 16:00:01
func NowStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeStr() string {
	return time.Now().Format("15:04:05")
}

func DateStr() string {
	return time.Now().Format("2006-01-02")
}

//return  today date  2019-12-30
func TodayStr() string {
	return time.Now().Format("2006-01-02")
}

//小数点后取位数, n 是取的位数 0～8
func Round(val float64, n int) float64 {
	if n == 1 {
		return math.Round(val*1e1) * 1e-1
	}
	if n == 2 {
		return math.Round(val*1e2) * 1e-2
	}
	if n == 3 {
		return math.Round(val*1e3) * 1e-3
	}
	if n == 4 {
		return math.Round(val*1e4) * 1e-4
	}
	if n == 5 {
		return math.Round(val*1e5) * 1e-5
	}
	if n == 6 {
		return math.Round(val*1e6) * 1e-6
	}
	if n == 7 {
		return math.Round(val*1e7) * 1e-7
	}
	if n == 8 {
		return math.Round(val*1e8) * 1e-8
	}
	return math.Round(val)
}

//流水号
func SeqID() string {
	return func() string { return fmt.Sprintf("%d", time.Now().Year())[2:] }() + string(Krand(10, 0))
}

func RantInt() string {
	return fmt.Sprintf("%d", rd.Int())
}

func RantInt8() string {
	return fmt.Sprintf("%08d", rd.Intn(100000000))
}

func RantInt10() string {
	return fmt.Sprintf("%010d", rd.Intn(10000000000))
}
