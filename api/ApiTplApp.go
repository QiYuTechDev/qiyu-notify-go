package api

import "encoding/json"
import "net/http"
import "bytes"
import "io"
import "github.com/QiYuTechDev/qiyu-notify-go/dt"

type ApiTplApp struct {
}

// Get
// 获取自定义模版APP详情
// queryArgs 字段有: unique_id
func (d ApiTplApp) Get(queryArgs map[string]string, bearer string) (*dt.TplAppInfoDt, error) {
	r := new(http.Request)
	r.Method = "GET"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/tpl/app"

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

	ret := new(dt.TplAppInfoDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
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

	r.Header.Set("Authorization", "bearer "+bearer)

	if t, err := json.Marshal(body); err == nil {
		r.Body = io.NopCloser(bytes.NewReader(t))
		r.Header.Set("Content-Type", "application/json")
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

	r.Header.Set("Accept", "application/json")

	ret := new(dt.TplAppInfoDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}

// Delete
// 删除自定义模版推送配置
// queryArgs 字段有: unique_id
func (d ApiTplApp) Delete(queryArgs map[string]string, bearer string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "DELETE"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/tpl/app"

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

// Patch
// 更新模版APP
func (d ApiTplApp) Patch(body dt.TplAppUpdateDt, bearer string) (*dt.TplAppInfoDt, error) {
	r := new(http.Request)
	r.Method = "PATCH"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/tpl/app"

	pathArgs := map[string]string{}
	queryArgs := map[string]string{}

	r.Header.Set("Authorization", "bearer "+bearer)

	if t, err := json.Marshal(body); err == nil {
		r.Body = io.NopCloser(bytes.NewReader(t))
		r.Header.Set("Content-Type", "application/json")
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

	r.Header.Set("Accept", "application/json")

	ret := new(dt.TplAppInfoDt)

	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, err
}
