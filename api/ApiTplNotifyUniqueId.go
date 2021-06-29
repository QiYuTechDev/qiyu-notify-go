package api

import "net/http"

type ApiTplNotifyUniqueId struct {
}

// Post
// 发送自定义模版通知消息
// pathArgs 字段有: unique_id
func (d ApiTplNotifyUniqueId) Post(pathArgs map[string]string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "POST"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/tpl/notify/{unique_id}"

	queryArgs := map[string]string{}

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
