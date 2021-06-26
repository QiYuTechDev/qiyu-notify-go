package api

import "encoding/json"
import "net/http"
import "bytes"
import "io"
import "../dt"

type ApiTplApp struct {
}

// Put
// 创建自定义模版APP
func (d ApiTplApp) Put(body dt.TplAppCreateDt, bearer string) (*dt.TplAppInfoDt, error) {
	r := new(http.Request)
	r.Method = "PUT"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/tpl/app"

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

	ret := new(dt.TplAppInfoDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}

// Delete
// 删除自定义模版推送配置
func (d ApiTplApp) Delete(queryArgs map[string]string, bearer string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "DELETE"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/tpl/app"

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
