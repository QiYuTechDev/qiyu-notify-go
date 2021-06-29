package api

import "encoding/json"
import "net/http"
import "github.com/QiYuTechDev/qiyu-notify-go/dt"

type ApiDdMsgList struct {
}

// Get
// 获取钉钉APP最新消息列表
// queryArgs 字段有: unique_id
func (d ApiDdMsgList) Get(queryArgs map[string]string, bearer string) (*[]dt.DdMsgDt, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/dd/msg/list"

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

	ret := new([]dt.DdMsgDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}
