package dlog

import "github.com/caoshuyu/kit/gls"

const (
	defaultDir     = "/data/logs"
	LOG_TYPE_LOCAL = 1 // 本地文件存储
	LOG_TYPE_NET   = 2 //网络存储（暂不支持）
)

var (
	logType                                                                   = LOG_TYPE_LOCAL
	logPath                                                                   = defaultDir
	traceInfoFunc func() (sessionId, actionId, spanId string, colorID string) = traceFunc //获取用户标记信息
	prefix        string                                                                  //项目名称 类似 go_user
)

type SetLogConf struct {
	LogType       int
	LogPath       string
	TraceInfoFunc func() (sessionId, actionId, spanId, colorID string)
	Prefix        string
}

//初始化日志设置
func SetLog(conf SetLogConf) {
	switch conf.LogType {
	case LOG_TYPE_LOCAL, LOG_TYPE_NET:
		logType = conf.LogType
	default:
		logType = LOG_TYPE_LOCAL
	}
	if len(conf.LogPath) > 0 {
		logPath = conf.LogPath
	}
	if nil != conf.TraceInfoFunc {
		traceInfoFunc = conf.TraceInfoFunc
	} else {
		traceInfoFunc = traceFunc
	}
	prefix = conf.Prefix
}

func traceFunc() (sessionId, actionId, spanId string, colorID string) {
	sessionId, actionId, spanId, colorID = gls.GetTraceInfo()
	return
}
