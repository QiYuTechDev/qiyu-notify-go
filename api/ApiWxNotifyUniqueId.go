package api

import "encoding/json"
import "net/http"
import "bytes"
import "io"
import "github.com/QiYuTechDev/qiyu-notify-go/dt"

type ApiWxNotifyUniqueId struct {
}

// Get
// 发送通知消息
// pathArgs 字段有: unique_id
// queryArgs 字段有: text
func (d ApiWxNotifyUniqueId) Get(pathArgs map[string]string, queryArgs map[string]string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/wx/notify/{unique_id}"

	url, err := buildUrl(baseUrl, pathUrl, pathArgs, queryArgs)
	if err != nil {
		return nil, err
	}
	r.URL = url

	data, err := sendHttpRequest(r)
	if err != nil {
		return nil, err
	}

	return data, err
}

// Post
// 发送通知消息
// pathArgs 字段有: unique_id
func (d ApiWxNotifyUniqueId) Post(pathArgs map[string]string, body dt.NotifyArgs) (interface{}, error) {
	r := new(http.Request)
	r.Method = "POST"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/wx/notify/{unique_id}"

	queryArgs := map[string]string{}

	if t, err := json.Marshal(body); err == nil {
		r.Body = io.NopCloser(bytes.NewReader(t))
		r.Header.Set("Content-Type", "application/json")
	} else {
		return nil, err
	}

	url, err := buildUrl(baseUrl, pathUrl, pathArgs, queryArgs)
	if err != nil {
		return nil, err
	}
	r.URL = url

	data, err := sendHttpRequest(r)
	if err != nil {
		return nil, err
	}

	return data, err
}
