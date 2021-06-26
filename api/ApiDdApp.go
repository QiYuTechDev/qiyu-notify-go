package api

import "encoding/json"
import "net/http"
import "bytes"
import "io"
import "../dt"

type ApiDdApp struct {
}

// Put
// 添加钉钉推送配置
func (d ApiDdApp) Put(body dt.DdAppCreateDt, bearer string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "PUT"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/dd/app"

	pathArgs := map[string]string{}
	queryArgs := map[string]string{}

	if bearer != "" {
		r.Header.Set("Authorization", "bearer "+bearer)
	}

	if t, err := json.Marshal(body); err == nil {
		r.Body = io.NopCloser(bytes.NewReader(t))
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

// Delete
// 删除钉钉推送配置
func (d ApiDdApp) Delete(queryArgs map[string]string, bearer string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "DELETE"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/dd/app"

	pathArgs := map[string]string{}

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
