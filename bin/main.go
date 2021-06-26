package main

import (
	"net/http"
	"net/url"
)

func main() {

	r := new(http.Request)
	url := new(url.URL)
	r.URL = url

	print("main")
}
