package rest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type HttpRestClient struct {
	timeout       time.Duration
	baseUri       string
	EnableLogging bool
}

const defaultTimeout = 20 * time.Second

func NewHttpRestClient(baseUri string) *HttpRestClient {
	return &HttpRestClient{
		timeout: defaultTimeout,
		baseUri: baseUri,
	}
}

func (restClient *HttpRestClient) Get(
	url string,
	queryParameters *url.Values,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	url = restClient.baseUri + url

	if queryParameters != nil && len(*queryParameters) > 0 {
		url += `?` + queryParameters.Encode()
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.Header.Add(`User-Agent`, `Core/EventFarm/golang/`+version)
	//req.Header.Add(`Accept-Encoding`, `gzip`)

	if restClient.EnableLogging {
		LogRequest(req)
	}

	client := new(http.Client)
	client.Timeout = defaultTimeout
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	if restClient.EnableLogging {
		LogResponse(resp)
	}

	return
}

func (restClient *HttpRestClient) Post(
	url string,
	formParameters *url.Values,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	url = restClient.baseUri + url
	body := bytes.NewBufferString(formParameters.Encode())

	req, err := http.NewRequest(`POST`, url, body)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.Header.Add(`User-Agent`, `Core/EventFarm/golang/`+version)
	req.Header.Add(`Content-Type`, `application/x-www-form-urlencoded;charset=UTF-8)`)
	//req.Header.Add(`Accept-Encoding`, `gzip`)

	if restClient.EnableLogging {
		LogRequest(req)
	}

	client := new(http.Client)
	client.Timeout = defaultTimeout
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	if restClient.EnableLogging {
		LogResponse(resp)
	}

	return
}

func (restClient *HttpRestClient) PostJSON(
	url string,
	data *map[string]interface{},
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	url = restClient.baseUri + url
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(data)

	req, err := http.NewRequest(`POST`, url, body)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.Header.Add(`User-Agent`, `Core/EventFarm/golang/`+version)
	req.Header.Add(`Content-Type`, `application/json;charset=UTF-8)`)
	//req.Header.Add(`Accept-Encoding`, `gzip`)

	if restClient.EnableLogging {
		LogRequest(req)
	}

	client := new(http.Client)
	client.Timeout = defaultTimeout
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	if restClient.EnableLogging {
		LogResponse(resp)
	}

	return
}

func (restClient *HttpRestClient) PostMultipart(
	url string,
	multipart map[string]string,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	url = restClient.baseUri + url

	req, err := http.NewRequest(`POST`, url, nil)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.Header.Add(`User-Agent`, `Core/EventFarm/golang/`+version)
	req.Header.Add(`Content-Type`, `application/x-www-form-urlencoded;charset=UTF-8)`)
	//req.Header.Add(`Accept-Encoding`, `gzip`)

	if restClient.EnableLogging {
		LogRequest(req)
	}

	client := new(http.Client)
	client.Timeout = defaultTimeout
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	if restClient.EnableLogging {
		LogResponse(resp)
	}

	return
}

func LogRequest(r *http.Request) {
	body, _ := httputil.DumpRequest(r, true)
	log.Println(`Request:`)
	log.Println(string(body))
}

func LogResponse(r *http.Response) {
	body, _ := httputil.DumpResponse(r, true)
	log.Println(`Response:`)
	log.Println(string(body))
}
