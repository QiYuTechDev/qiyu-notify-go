package api

import "net/http"

type ApiDdLinkUniqueId struct {
}

// Get
// 发送通知消息
func (d ApiDdLinkUniqueId) Get(pathArgs map[string]string, queryArgs map[string]string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/dd/link/{unique_id}"

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
