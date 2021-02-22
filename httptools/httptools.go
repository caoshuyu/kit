package httptools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

const (
	HTTP_DEFAULT_TIME_OUT = 120
)

type ExternalHttpRequest struct {
	Jar          *cookiejar.Jar
	Method       string
	Header       map[string]string
	TimeOut      int
	Domain       string
	Uri          string
	Body         string
	ReqUriParams string
}

type ExternalHttpResponse struct {
	StatusCode int
	Header     map[string][]string
	Body       []byte
	UseTime    int64
}

//Making HTTP Requests
func HTTPRequest(request *ExternalHttpRequest) (response *ExternalHttpResponse, err error) {
	reqBody := bytes.NewBuffer([]byte(request.Body))
	tOut := request.TimeOut
	if tOut <= 0 {
		tOut = HTTP_DEFAULT_TIME_OUT
	}
	var c http.Client
	if nil == request.Jar {
		c = http.Client{
			Timeout: time.Second * time.Duration(tOut),
		}
	} else {
		c = http.Client{
			Timeout: time.Second * time.Duration(tOut),
			Jar:     request.Jar,
		}
	}
	reqUrl := request.Domain + request.Uri
	if len(request.ReqUriParams) > 0 {
		if strings.Index(reqUrl, "?") > -1 {
			reqUrl += "&" + request.ReqUriParams
		} else {
			reqUrl += "?" + request.ReqUriParams
		}
	}
	reqUrl = urlEncode(reqUrl)

	startTime := time.Now()
	var req *http.Request
	req, err = http.NewRequest(request.Method, reqUrl, reqBody)
	if nil != err {
		return
	}
	for k, v := range request.Header {
		req.Header.Add(k, v)
	}
	resp, err := c.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if nil != err {
		return
	}
	respBodyByte, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return
	}
	statusCode := resp.StatusCode
	respHeader := make(map[string][]string)
	for k, v := range resp.Header {
		respHeader[k] = v
	}
	endTime := time.Now()
	response = &ExternalHttpResponse{
		StatusCode: statusCode,
		Header:     respHeader,
		Body:       respBodyByte,
		UseTime:    startTime.UnixNano() - endTime.UnixNano(),
	}
	return
}

func urlEncode(url string) (str string) {
	str = strings.Replace(url, " ", "%20", -1)
	str = strings.Replace(str, "'", "%27", -1)
	return
}
