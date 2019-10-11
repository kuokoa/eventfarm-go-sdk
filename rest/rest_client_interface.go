package rest

import (
	"net/http"
	"net/url"
	"time"
)

type RestClientInterface interface {
	Get(
		url string,
		queryParameters *url.Values,
		headers map[string]string,
		timeout *time.Duration,
	) (resp *http.Response, err error)

	Post(
		url string,
		formParameters *url.Values,
		headers map[string]string,
		timeout *time.Duration,
	) (resp *http.Response, err error)

	PostJSON(
		url string,
		data *map[string]interface{},
		headers map[string]string,
		timeout *time.Duration,
	) (resp *http.Response, err error)

	PostMultipart(
		url string,
		multipart map[string]string,
		headers map[string]string,
		timeout *time.Duration,
	) (resp *http.Response, err error)
}
