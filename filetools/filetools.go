package filetools

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
检测文件是否存在
*/
func CheckFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

/*
递归创建文件夹
*/
func MakeDir(dirPath string) (err error) {
	if 0 == len(dirPath) {
		err = errors.New("dir path not empty")
		return
	}
	dirPath = strings.Replace(dirPath, "\\", "/", -1)
	pathArr := strings.Split(dirPath, "/")
	pathLen := len(pathArr)
	num := strings.Index(dirPath, ".")
	if num > -1 {
		pathLen -= 1
	}
	for i := 1; i <= pathLen; i++ {
		nowPath := strings.Join(pathArr[:i], "/")
		if CheckFileExist(nowPath) {
			continue
		}
		err = os.Mkdir(nowPath, 0755) //系统默认文件夹权限，如果需要别的权限创建后可进行修改
		if nil != err {
			return
		}
	}
	return
}

/*
读取文件
*/
func ReadFile(filename string) (val string) {
	if !CheckFileExist(filename) {
		return
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666) //打开文件
	if nil != err {
		return
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	return string(fd)
}

/*
读取byte文件
*/
func ReadFileByte(filename string) (val []byte) {
	if !CheckFileExist(filename) {
		return
	}
	f, err := os.Open(filename)
	if nil != err {
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if nil != err && io.EOF != err {
			return
		}
		if 0 == n {
			break
		}
		val = append(val, buf[:n]...)
	}
	return
}

/*
写文件
*/
func WriteFile(filename string, value string) (err error) {
	return _writeFile(filename, value, os.O_RDWR|os.O_APPEND)
}

func WriteFileCover(filename string, value string) (err error) {
	os.Remove(filename)
	return _writeFile(filename, value, os.O_RDWR)
}

func _writeFile(filename string, value string, flag int) (err error) {
	if !CheckFileExist(filename) {
		//生成文件
		_, err = os.Create(filename)
	}
	f, err := os.OpenFile(filename, flag, 0666) //打开文件
	if nil != err {
		panic(err)
		return
	}
	defer f.Close()
	n, err := io.WriteString(f, value)
	if nil != err {
		panic(err)
		return
	}
	if 0 == n {
		err = errors.New("no byte write")
	}
	return
}

/*
写文件,字符类型
*/
func WriteFileByte(filename string, value []byte) (err error) {
	if !CheckFileExist(filename) {
		//生成文件
		_, err = os.Create(filename)
	}
	err = ioutil.WriteFile(filename, value, 0666)
	return
}
