package echomiddleware

import (
	"github.com/caoshuyu/kit/gls"
	"github.com/caoshuyu/kit/nettools"
	"github.com/labstack/echo/v4"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Gls(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		var sessionId = context.Request().Header.Get("header-sessionId")
		var actionId = context.Request().Header.Get("header-actionId")
		var spanID = GenerateSpanID(context.Request().RemoteAddr)
		var color = context.Request().Header.Get("color")
		gls.SetGls(sessionId, actionId, spanID, color, func() {
			handlerFunc(context)
		})
		return nil
	}
}

func GenerateSpanID(addr string) string {
	strAddr := strings.Split(addr, ":")
	ip := strAddr[0]
	ipLong := nettools.Ipv4ToInt64(ip)
	times := uint64(time.Now().UnixNano())
	spanId := ((times ^ uint64(ipLong)) << 32) | uint64(rand.Int31())
	return strconv.FormatUint(spanId, 16)
}
