package echo_out_tools

import (
	"github.com/caoshuyu/kit/echomiddleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type ComResponse struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

const HTTP_CODE_OK = 0

func EchoErrorData(context echo.Context, err error, errCode int64) error {
	conData := ComResponse{
		Code: errCode,
		Msg:  err.Error(),
	}
	echomiddleware.SetDataOut(context, conData)
	return context.JSON(http.StatusOK, conData)
}

func EchoSuccessData(context echo.Context, data interface{}) error {
	var conData ComResponse
	switch data.(type) {
	case string:
		if strings.EqualFold(data.(string), "") {
			s := struct{}{}
			data = s
		}
	}
	conData = ComResponse{
		Code: HTTP_CODE_OK,
		Data: data,
	}
	echomiddleware.SetDataOut(context, conData)
	return context.JSON(http.StatusOK, conData)
}
