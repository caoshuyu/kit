package gls

import "github.com/jtolds/gls"

var (
	mgr          = gls.NewContextManager()
	sessionIdKey = "session_id"
	actionIdKey  = "action_id"
	spanIDKey    = "span_id"
	colorKey     = "color" //流量染色
)

func SetGls(sessionId, actionId, spanID string, color string, cb func()) {
	mgr.SetValues(gls.Values{sessionIdKey: sessionId, actionIdKey: actionId, spanIDKey: spanID, colorKey: color}, cb)
}
func GetTraceInfo() (sessionId, actionId string, spanID string, colorID string) {
	session, ok := mgr.GetValue(sessionIdKey)
	if ok {
		sessionId = session.(string)
	}
	action, ok := mgr.GetValue(actionIdKey)
	if ok {
		actionId = action.(string)
	}
	span, ok := mgr.GetValue(spanIDKey)
	if ok {
		spanID = span.(string)
	}
	color, ok := mgr.GetValue(colorKey)
	if ok {
		colorID = color.(string)
	}
	return
}
