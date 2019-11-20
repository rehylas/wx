package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

//RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}

func RandomInt(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(result, "")
}

//Struct2Map    map[string]interface{}
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

//Struct2Map   key 均转化为小写
func Struct2MapEx(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

// 参考函数
// func (this Stuff) MarshalJSON() ([]byte, error) {
// 	buffer := bytes.NewBufferString("{")
// 	length := len(this)
// 	count := 0
// 	for key, value := range this {
// 		jsonValue, err := json.Marshal(value)
// 		if err != nil {
// 			return nil, err
// 		}
// 		buffer.WriteString(fmt.Sprintf("\"%d\":%s", key, string(jsonValue)))
// 		count++
// 		if count < length {
// 			buffer.WriteString(",")
// 		}
// 	}
// 	buffer.WriteString("}")
// 	return buffer.Bytes(), nil
// }

//  未调通
// func PickJsonFields(obj interface{}, fields []string) ([]byte, error) {
// 	buffer := bytes.NewBufferString("{")
// 	length := len(this)
// 	count := 0
// 	for key, value := range this {
// 		jsonValue, err := json.Marshal(value)
// 		if err != nil {
// 			return nil, err
// 		}

// 		buffer.WriteString(fmt.Sprintf("\"%d\":%s", key, string(jsonValue)))
// 		count++
// 		if count < length {
// 			buffer.WriteString(",")
// 		}
// 	}
// 	buffer.WriteString("}")
// 	return buffer.Bytes(), nil
// }
