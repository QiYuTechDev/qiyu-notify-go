package api

import "net/http"

type ApiDdRawUniqueId struct {
}

// Post
// 发送通知消息
func (d ApiDdRawUniqueId) Post(pathArgs map[string]string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "POST"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/dd/raw/{unique_id}"

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
