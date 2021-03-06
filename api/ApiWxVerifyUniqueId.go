package api

import "net/http"

type ApiWxVerifyUniqueId struct {
}

// Get
// 企业微信验证
// pathArgs 字段有: unique_id
// queryArgs 字段有: timestamp msg_signature nonce echostr
func (d ApiWxVerifyUniqueId) Get(pathArgs map[string]string, queryArgs map[string]string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/wx/verify/{unique_id}"

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
