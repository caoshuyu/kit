package dlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/caoshuyu/kit/filetools"
	"runtime"
	"strings"
	"time"
)

var writeLogChan = make(chan string, 10000) //10000长度缓冲

func addNewLog(logStr string) {
	writeLogChan <- logStr
}

func writeLogInfo() {
	//每秒进行一次写操作
	l := time.NewTicker(time.Duration(1) * time.Second)
	var writeArr []string
	for {
		select {
		case logInfo := <-writeLogChan:
			writeArr = append(writeArr, logInfo)
		case <-l.C:
			if len(writeArr) > 0 {
				writeStr := strings.Join(writeArr, "")
				writeArr = make([]string, 0)

				switch logType {
				case LOG_TYPE_LOCAL:
					d := dLog{}
					filetools.WriteFile(d.getFileSavePath(), writeStr)
				default:

				}
			}
		}
	}
}

type logVal struct {
	key string
	val interface{}
}

func marshalJSON(logArr []logVal) ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{")
	for i, kv := range logArr {
		if 0 != i {
			buf.WriteString(",")
		}
		valJson, err := json.Marshal(kv.val)
		if nil != err {
			buf.WriteString(`"` + kv.key + `"` + ":")
			continue
		}
		buf.WriteString(`"` + kv.key + `"` + ":" + string(valJson))
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

//kv 成对出现 类似 name:shuyu.s.cao
func (p *dLog) logStr(level string, kv ...interface{}) ([]byte, error) {
	_, file, line, _ := runtime.Caller(3)
	file = p.getFilePath(file)
	sessionId, actionId, spanId := traceInfoFunc()
	logValArr := make([]logVal, 0, 10)
	logValArr = append(logValArr,
		logVal{key: "session_id", val: sessionId},
		logVal{key: "action_id", val: actionId},
		logVal{key: "span_id", val: spanId},
		logVal{key: "prefix", val: prefix},
		logVal{key: "file", val: file},
		logVal{key: "level", val: level},
		logVal{key: "line", val: line},
		logVal{key: "cre_time", val: time.Now().Format(time.RFC3339Nano)},
		logVal{key: "cre_unix_time", val: time.Now().Unix()},
		logVal{key: "__tag_path__", val: p.getFileSavePath()},
	)
	if 0 != len(kv)%2 {
		kv = append(kv, "unknown")
	}
	for i := 0; i < len(kv); i += 2 {
		v := kv[i+1]
		if err, ok := v.(error); ok {
			logValArr = append(logValArr, logVal{key: fmt.Sprintf("%v", kv[i]), val: err.Error()})
		} else {
			logValArr = append(logValArr, logVal{key: fmt.Sprintf("%v", kv[i]), val: kv[i+1]})
		}
	}
	return marshalJSON(logValArr)
}


