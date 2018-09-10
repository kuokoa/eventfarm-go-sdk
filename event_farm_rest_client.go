package gosdk

import (
	"net/http"
	"net/url"
	"time"
)

const version = `6.4.x`

type EventFarmRestClient struct {
	restClient            RestClientInterface
	accessTokenRestClient RestClientInterface
	clientId              string
	clientSecret          string
	oAuthAccessToken      *OAuthAccessToken
}

func NewEventFarmRestClient(
	restClient RestClientInterface,
	accessTokenRestClient RestClientInterface,
	clientId string,
	clientSecret string,
	oAuthAccessToken *OAuthAccessToken,
) *EventFarmRestClient {
	return &EventFarmRestClient{
		restClient:            restClient,
		accessTokenRestClient: accessTokenRestClient,
		clientId:              clientId,
		clientSecret:          clientSecret,
		oAuthAccessToken:      oAuthAccessToken,
	}
}

func (t *EventFarmRestClient) Get(
	url string,
	queryParameters *url.Values,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	newHeaders, err := t.getAuthorizationHeader(headers)
	if err != nil {
		return
	}

	return t.restClient.Get(
		url,
		queryParameters,
		newHeaders,
		timeout,
	)
}

func (t *EventFarmRestClient) Post(
	url string,
	formParameters *url.Values,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	newHeaders, err := t.getAuthorizationHeader(headers)
	if err != nil {
		return
	}

	return t.restClient.Post(
		url,
		formParameters,
		newHeaders,
		timeout,
	)
}

func (t *EventFarmRestClient) PostMultipart(
	url string,
	multipart map[string]string,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	newHeaders, err := t.getAuthorizationHeader(headers)
	if err != nil {
		return
	}

	return t.restClient.PostMultipart(
		url,
		multipart,
		newHeaders,
		timeout,
	)
}

func (t *EventFarmRestClient) getAuthorizationHeader(headers map[string]string) (newHeaders map[string]string, err error) {
	newHeaders = make(map[string]string)

	for k, v := range headers {
		newHeaders[k] = v
	}

	oAuthAccessToken, err := t.getOAuthAccessToken()
	if err != nil {
		return
	}

	newHeaders[`Authorization`] = oAuthAccessToken.getHeaderString()

	return
}

func (t *EventFarmRestClient) getOAuthAccessToken() (oAuthAccessToken *OAuthAccessToken, err error) {
	if t.oAuthAccessToken == nil {
		t.oAuthAccessToken, err = t.getClientCredentialsAccessToken()
		if err != nil {
			return
		}
	}

	if t.oAuthAccessToken.isExpired() {
		if t.oAuthAccessToken.refreshToken != `` {
			t.oAuthAccessToken, err = t.getRefreshToken(t.oAuthAccessToken.refreshToken)
			if err != nil {
				oAuthAccessToken, err = t.getClientCredentialsAccessToken()
				if err != nil {
					return
				}
			}
		} else {
			t.oAuthAccessToken, err = t.getClientCredentialsAccessToken()
			if err != nil {
				return
			}
		}
	}

	oAuthAccessToken = t.oAuthAccessToken
	return
}

func (t *EventFarmRestClient) getRefreshToken(refreshToken string) (oAuthAccessToken *OAuthAccessToken, err error) {
	values := url.Values{}
	values.Add(`grant_type`, `refresh_token`)
	values.Add(`refresh_token`, refreshToken)
	values.Add(`client_id`, t.clientId)
	values.Add(`client_secret`, t.clientSecret)

	resp, err := t.accessTokenRestClient.Post(
		`/oauth/token.json`,
		&values,
		nil,
		nil,
	)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	oAuthAccessToken, err = NewOAuthAccessTokenFromResponse(resp)
	if err != nil {
		return
	}

	return
}

func (t *EventFarmRestClient) getClientCredentialsAccessToken() (oAuthAccessToken *OAuthAccessToken, err error) {
	values := url.Values{}
	values.Add(`grant_type`, `client_credentials`)
	values.Add(`client_id`, t.clientId)
	values.Add(`client_secret`, t.clientSecret)

	resp, err := t.accessTokenRestClient.Post(
		`/oauth/token.json`,
		&values,
		nil,
		nil,
	)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	oAuthAccessToken, err = NewOAuthAccessTokenFromResponse(resp)
	if err != nil {
		return
	}

	return
}
