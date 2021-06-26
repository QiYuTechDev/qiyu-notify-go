package api

import (
	"io"
	"net/http"
)

func sendHttpRequest(r *http.Request) (io.ReadCloser, error) {
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
