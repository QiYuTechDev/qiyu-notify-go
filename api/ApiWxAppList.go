package api

import "encoding/json"
import "net/http"
import "github.com/QiYuTechDev/qiyu-notify-go/dt"

type ApiWxAppList struct {
}

// Get
// 企业微信配置列表
func (d ApiWxAppList) Get(bearer string) (*[]dt.WxAppInfoDt, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/wx/app/list"

	pathArgs := map[string]string{}
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

	ret := new([]dt.WxAppInfoDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}
