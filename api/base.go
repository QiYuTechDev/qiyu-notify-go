package api

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func sendHttpRequest(r *http.Request) (io.ReadCloser, error) {
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func buildUrl(baseUrl, pathUrl string, params map[string]string) (*url.URL, error) {
	parts := strings.Split(pathUrl, "/")
	out := make([]string, len(parts))
	for idx, p := range parts {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			out[idx] = params[p[1:len(p)-1]]
		} else {
			out[idx] = p
		}
	}
	realPath := strings.Join(out, "/")
	fullUrl := baseUrl + realPath
	return url.Parse(fullUrl)
}
