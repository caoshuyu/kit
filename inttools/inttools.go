package inttools

import (
	"strconv"
	"strings"
)

//数字数组转字符串
//input [3306,4450,3,12],-
//out 3306-4450-3-12
func Int64ArrToString(data []int64, split string) string {
	strArr := make([]string, 0, len(data))
	for _, v := range data {
		strArr = append(strArr, strconv.FormatInt(v, 10))
	}
	return strings.Join(strArr, split)
}

//字符串分割为数字数组
//input 3306-4450-3-12,-
//out [3306,4450,3,12]
func StringToInt64Arr(data string, split string) (val []int64) {
	strArr := strings.Split(data, split)
	if 0 == len(strArr) {
		return
	}
	for _, v := range strArr {
		one, err := strconv.ParseInt(v, 10, 64)
		if nil == err {
			val = append(val, one)
		}
	}
	return
}

//通过页码和每页数量计算起始行数
func GetPageNum(page int64, pageContext int64) (newPage, newPageContext, startLine int64) {
	if 0 >= page {
		page = 1
	}
	if 0 >= pageContext {
		pageContext = 10
	}
	startLine = (page - 1) * pageContext
	newPage = page
	newPageContext = pageContext
	return
}

// Intersect 取两个数组交集
func Intersect(arr1 []int64, arr2 []int64) []int64 {
	var arr []int64

	m1 := make(map[int64]int64)
	for _, v := range arr1 {
		m1[v]++
	}

	for _, v := range arr2 {
		if _, ok := m1[v]; ok {
			arr = append(arr, v)
		}
	}
	return arr
}

