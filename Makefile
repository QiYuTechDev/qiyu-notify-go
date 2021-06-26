gen-go-source-code:
	rm -f code.zip || true

	curl -X 'POST' \
	  'https://code.qiyutech.tech/code/gen/url' \
	  -H 'accept: */*' \
	  -H 'Content-Type: application/x-www-form-urlencoded' \
	  -d 'openapi_url=https%3A%2F%2Fnotify.qiyutech.tech%2Fapi%2Fopenapi.json&dt_class_name=&dt_prefix=&stub_class_name=&stub_prefix=&stub_postfix=&dt_postfix=&php_namespace=&language=go&path_base_url=https%3A%2F%2Fnotify.qiyutech.tech' \
	  --output code.zip

	unzip code.zip && rm -f code.zip
