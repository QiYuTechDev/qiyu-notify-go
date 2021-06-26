package api

import "encoding/json"
import "net/http"
import "../dt"

type ApiTplAppList struct {
}

// Get
// 获取所有的自定义模版推送的配置
func (d ApiTplAppList) Get(bearer string) (*[]dt.TplAppInfoDt, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/tpl/app/list"

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

	ret := new([]dt.TplAppInfoDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}
