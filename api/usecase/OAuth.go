package usecase

import (
	"gosdk"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type OAuth struct {
	restClient gosdk.RestClientInterface
}

func NewOAuth(restClient gosdk.RestClientInterface) *OAuth {
	return &OAuth{restClient}
}

// GET: Queries
// @param string OauthAccessTokenId
func (t *OAuth) GetOAuthAccessToken(OauthAccessTokenId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`oauthAccessTokenId`, OauthAccessTokenId)

	return t.restClient.Get(
		`/v2/OAuth/UseCase/GetOAuthAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string Email
func (t *OAuth) CreateGhostAccessToken(Email string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, Email)

	return t.restClient.Post(
		`/v2/OAuth/UseCase/CreateGhostAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Name
// @param array RedirectUrls
// @param string|null Identifier
// @param string|null Secret
func (t *OAuth) CreateOAuthClient(Name string, RedirectUrls []string, Identifier *string, Secret *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, Name)
	for i := range RedirectUrls {
		queryParameters.Add(`redirectUrls[]`, RedirectUrls[i])
	}
	if Identifier != nil {
		queryParameters.Add(`identifier`, *Identifier)
	}
	if Secret != nil {
		queryParameters.Add(`secret`, *Secret)
	}

	return t.restClient.Post(
		`/v2/OAuth/UseCase/CreateOAuthClient`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Identifier
func (t *OAuth) RevokeAccessToken(Identifier string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`identifier`, Identifier)

	return t.restClient.Post(
		`/v2/OAuth/UseCase/RevokeAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Identifier
// @param array RedirectUrls
func (t *OAuth) SetRedirectUrlsForOAuthClient(Identifier string, RedirectUrls []string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`identifier`, Identifier)
	for i := range RedirectUrls {
		queryParameters.Add(`redirectUrls[]`, RedirectUrls[i])
	}

	return t.restClient.Post(
		`/v2/OAuth/UseCase/SetRedirectUrlsForOAuthClient`,
		&queryParameters,
		nil,
		nil,
	)
}
