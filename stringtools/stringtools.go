package stringtools

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//首字母大写，有_的去掉后第一个字母大写
func InitialUpdateStr(str string) string {
	if 0 == len(str) {
		return ""
	}
	strList := strings.Split(str, "_")
	for k, v := range strList {
		strList[k] = strings.ToUpper(v[0:1]) + v[1:]
	}
	return strings.Join(strList, "")
}

//首字母小写，有_的去掉后第一个字母大写
func InitialLowStr(str string) string {
	s := InitialUpdateStr(str)
	return strings.ToLower(s[0:1]) + s[1:]
}

//变成大写字母，用_分割
func UpperToUnderlineToUpper(str string) string {
	if 0 == len(str) {
		return ""
	}
	return strings.ToUpper(UpperToUnderline(str))
}

//变成小写字母，用_分割
func UpperToUnderline(str string) string {
	if 0 == len(str) {
		return ""
	}
	rege := "[A-Z]*[^A-Z]+"
	reg, err := regexp.Compile(rege)
	if nil != err {
		return ""
	}
	value := reg.FindAllString(str, -1)
	info := make([]string, 0, len(value))
	for _, v := range value {
		info = append(info, strings.ToLower(v))
	}
	newDataType := strings.Join(info, "_")
	return newDataType
}

//制作问号
func MakeRoundabout(num int) string {
	str := strings.Repeat("?,", num)
	l := len(str)
	if l > 0 {
		str = str[:l-1]
	}
	return str
}

//MD5加密
func Md5(str string) string {
	md5Obj := md5.New()
	md5Obj.Write([]byte(str))
	cipherStr := md5Obj.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//Time33加密算法
func Time33(key string) int64 {
	md5Key := Md5(key)
	strKey := md5Key[0:10]
	strByteArr := []byte(strKey)
	hashNum := 0
	for _, val := range strByteArr {
		hashNum += hashNum*33 + int(val)
	}
	str := hashNum & 0x7FFFFFFF
	return int64(math.Abs(float64(str)))
}

//判断字符串是不是数字
func IsNum(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

//去除数组中重复的字符串
func RemoveSliceRepeatStr(arr1, arr2 []string) (result []string) {
	if len(arr1) == 0 || len(arr2) == 0 {
		result = append(arr1, arr2...)
		return
	}
	strFlag := make(map[string]bool)
	for _, item := range arr1 {
		if _, h := strFlag[item]; !h {
			result = append(result, item)
			strFlag[item] = true
		}
	}
	for _, item := range arr2 {
		if _, h := strFlag[item]; !h {
			result = append(result, item)
			strFlag[item] = true
		}
	}
	return
}
