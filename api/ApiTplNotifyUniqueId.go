package api

import "net/http"

type ApiTplNotifyUniqueId struct {
}

// Post
// 发送自定义模版通知消息
func (d ApiTplNotifyUniqueId) Post(pathArgs map[string]string, bearer string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "POST"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/tpl/notify/{unique_id}"

	queryArgs := map[string]string{}

	if bearer != "" {
		r.Header.Set("Authorization", "bearer "+bearer)
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
