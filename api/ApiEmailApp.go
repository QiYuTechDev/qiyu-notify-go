package api

import "net/http"

type ApiEmailApp struct {
}

// Put
// 添加电子邮箱
// queryArgs 字段有: email
func (d ApiEmailApp) Put(queryArgs map[string]string, bearer string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "PUT"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/email/app"

	pathArgs := map[string]string{}

	r.Header.Set("Authorization", "bearer "+bearer)

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

// Delete
// 删除电子邮箱推送配置
// queryArgs 字段有: unique_id
func (d ApiEmailApp) Delete(queryArgs map[string]string, bearer string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "DELETE"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/email/app"

	pathArgs := map[string]string{}

	r.Header.Set("Authorization", "bearer "+bearer)

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
