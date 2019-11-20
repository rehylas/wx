package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//http client  get url  ,  return body
func HttpGet(url string, params map[string]string) (string, error) {

	paramsStr := ""
	if len(params) > 0 {
		paramsStr = "?"
		for key, val := range params {
			paramsStr = paramsStr + fmt.Sprintf("%s=%s", key, val) + "&"
		}
		paramsStr = paramsStr[0 : len(paramsStr)-1]
	}
	resp, err := http.Get(url + paramsStr)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(body))
	return string(body), nil

}

//http client post json , return body
func HttpPostJson(url string, data string) (string, error) {

	reader := bytes.NewReader([]byte(data))
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err

	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err

	}
	//byte数组直接转成string，优化内存  有必要 ???
	// str := (*string)(unsafe.Pointer(&respBytes))
	// return *str, nil

	return string(respBytes), nil

}
