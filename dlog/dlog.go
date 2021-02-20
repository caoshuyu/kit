package dlog

import (
	"path"
	"time"
)

type dLog struct {
}

func (p *dLog) getFileSavePath() string {
	return logPath + prefix + ".log_json." + time.Now().Format("2006-01-02")
}

func (p *dLog) getFilePath(file string) string {
	dir, base := path.Dir(file), path.Base(file)
	return path.Join(path.Base(dir), base)
}

func DEBUG(kv ...interface{}) {
	writeLog("DEBUG", kv...)

}
func INFO(kv ...interface{}) {
	writeLog("INFO", kv...)
}

func WARN(kv ...interface{}) {
	writeLog("WARN", kv...)
}

func ERROR(kv ...interface{}) {
	writeLog("ERROR", kv...)
}

func writeLog(level string, kv ...interface{}) {
	d := dLog{}
	js, err := d.logStr(level, kv...)
	if nil != err {
		return
	}
	addNewLog(string(js) + "\n")
}
