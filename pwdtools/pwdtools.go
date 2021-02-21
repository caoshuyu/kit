package pwdtools

import (
	"github.com/caoshuyu/kit/stringtools"
	"strconv"
	"time"
)

/*
根据用户密码获取加密码和加密密码
*/
func MakePassword(word string) (pwd string, password string) {
	pwd = stringtools.Md5(strconv.Itoa(int(time.Now().Unix())))
	pwd = pwd[:6]
	password = stringtools.Md5(stringtools.Md5(word) + pwd)
	return
}

/*
根据加密码和用户密码获取加密后密码
*/
func MakePasswordByPWD(pwd string, password string) string {
	return stringtools.Md5(stringtools.Md5(password) + pwd)
}
