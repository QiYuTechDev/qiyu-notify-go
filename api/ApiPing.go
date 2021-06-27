package api

import "net/http"

type ApiPing struct {
}

// Get
// 健康检查
func (d ApiPing) Get() (interface{}, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/ping"

	pathArgs := map[string]string{}
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
