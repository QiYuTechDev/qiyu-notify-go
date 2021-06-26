package api

import "encoding/json"
import "net/http"
import "github.com/QiYuTechDev/qiyu-notify-go/dt"

type ApiEmailAppList struct {
}

// Get
// 获取所有的电子邮箱配置
func (d ApiEmailAppList) Get(bearer string) (*[]dt.EmailAppInfoDt, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/email/app/list"

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

	ret := new([]dt.EmailAppInfoDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}
