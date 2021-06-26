package api

import "encoding/json"
import "net/http"
import "bytes"
import "io"
import "../dt"

type ApiWxChatwootUniqueId struct {
}

// Post
// 企业微信 Chatwoot WebHook 回调
func (d ApiWxChatwootUniqueId) Post(pathArgs map[string]string, body dt.ChatwootWebHookDt, bearer string) (interface{}, error) {
	r := new(http.Request)
	r.Method = "POST"

	baseUrl := "https://notify.qiyutech.tech"
	pathUrl := "/api/wx/chatwoot/{unique_id}"

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
