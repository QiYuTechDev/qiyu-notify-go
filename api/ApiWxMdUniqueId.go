package api

import "encoding/json"
import "net/http"
import "bytes"
import "io"
import "github.com/QiYuTechDev/qiyu-notify-go/dt"

type ApiWxMdUniqueId struct {
}

// Post
// 发送 Markdown 通知消息到企业微信
// pathArgs 字段有: unique_id
func (d ApiWxMdUniqueId) Post(pathArgs map[string]string, body dt.NotifyArgs) (interface{}, error) {
	r := new(http.Request)
	r.Method = "POST"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/wx/md/{unique_id}"

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
