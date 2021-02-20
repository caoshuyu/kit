package echomiddleware

import (
	"bytes"
	"github.com/caoshuyu/kit/dlog"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"time"
)

//访问日志
func Access(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) (err error) {
		//保存body
		body, _ := ioutil.ReadAll(ectx.Request().Body)
		SetBody(ectx, body)
		ectx.Request().Body = ioutil.NopCloser(bytes.NewReader(body))
		start := time.Now()
		defer func() {
			accessLog(ectx, time.Since(start), GetBody(ectx), GetDataOut(ectx))
		}()
		err = next(ectx)
		return
	}
}

func SetHttpStatus(ectx echo.Context, httpStatus int) {
	ectx.Set("httpStatus", httpStatus)
}
func GetHttpStatus(ectx echo.Context) (httpStatus int) {
	b := ectx.Get("httpStatus")
	if b == nil {
		return http.StatusOK
	}
	return b.(int)
}

func SetDataIn(ectx echo.Context, dataIn interface{}) {
	ectx.Set("dataIn", dataIn)
}
func GetDataIn(ectx echo.Context) (dataIn interface{}) {
	return ectx.Get("dataIn")
}
func SetDataOut(ectx echo.Context, dataOut interface{}) {
	ectx.Set("dataOut", dataOut)
}
func GetDataOut(ectx echo.Context) (dataOut interface{}) {
	return ectx.Get("dataOut")
}

func SetBody(ectx echo.Context, body []byte) {
	ectx.Set("reqbody", body)
}
func GetBody(ectx echo.Context) (body []byte) {
	b := ectx.Get("reqbody")
	if b == nil {
		return nil
	}
	return b.([]byte)
}

func SetColor(ectx echo.Context, color string) {
	ectx.Set("color", color)
}

func GetColor(ectx echo.Context) (color string) {
	b := ectx.Get("color")
	if b == nil {
		return ""
	}
	return b.(string)
}

func accessLog(c echo.Context, dur time.Duration, body []byte, dataOut interface{}) {
	req := c.Request()
	color := GetColor(c)
	bodyStr := string(body)
	query := req.URL.RawQuery
	path := req.URL.Path
	dlog.INFO("type", "access",
		"color", color,
		"ip", c.RealIP(),
		"method", req.Method,
		"path", path,
		"query", query,
		"body", bodyStr,
		"output", dataOut,
		"time(ms)", int64(dur/time.Millisecond))
}
