package api

import "encoding/json"
import "net/http"
import "github.com/QiYuTechDev/qiyu-notify-go/dt"

type ApiWxMsgList struct {
}

// Get
// 企业微信消息列表
func (d ApiWxMsgList) Get(queryArgs map[string]string, bearer string) (*[]dt.WxMsgDt, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/wx/msg/list"

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

	r.Header.Set("Accept", "application/json")

	ret := new([]dt.WxMsgDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}
