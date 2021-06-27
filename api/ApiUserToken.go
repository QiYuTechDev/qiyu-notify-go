package api

import "encoding/json"
import "net/http"
import "github.com/QiYuTechDev/qiyu-notify-go/dt"

type ApiUserToken struct {
}

// Get
// 获取一个新的Token
func (d ApiUserToken) Get() (*dt.UserTokenDt, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/user/token"

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

	r.Header.Set("Accept", "application/json")

	ret := new(dt.UserTokenDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}
