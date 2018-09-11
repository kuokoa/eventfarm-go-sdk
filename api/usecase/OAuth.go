package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"bitbucket.ef.network/go/sdk"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type OAuth struct {
	restClient sdk.RestClientInterface
}

func NewOAuth(restClient sdk.RestClientInterface) *OAuth {
	return &OAuth{restClient}
}

// GET: Queries
// @param string OauthAccessTokenId

type GetOAuthAccessTokenParameters struct {
	OauthAccessTokenId string
}

func (t *OAuth) GetOAuthAccessToken(p *GetOAuthAccessTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`oauthAccessTokenId`, p.OauthAccessTokenId)

	return t.restClient.Get(
		`/v2/OAuth/UseCase/GetOAuthAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string Email

type CreateGhostAccessTokenParameters struct {
	Email string
}

func (t *OAuth) CreateGhostAccessToken(p *CreateGhostAccessTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, p.Email)

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

type CreateOAuthClientParameters struct {
	Name         string
	RedirectUrls []string
	Identifier   *string
	Secret       *string
}

func (t *OAuth) CreateOAuthClient(p *CreateOAuthClientParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, p.Name)
	for i := range p.RedirectUrls {
		queryParameters.Add(`redirectUrls[]`, p.RedirectUrls[i])
	}
	if p.Identifier != nil {
		queryParameters.Add(`identifier`, *p.Identifier)
	}
	if p.Secret != nil {
		queryParameters.Add(`secret`, *p.Secret)
	}

	return t.restClient.Post(
		`/v2/OAuth/UseCase/CreateOAuthClient`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Identifier

type RevokeAccessTokenParameters struct {
	Identifier string
}

func (t *OAuth) RevokeAccessToken(p *RevokeAccessTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`identifier`, p.Identifier)

	return t.restClient.Post(
		`/v2/OAuth/UseCase/RevokeAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Identifier
// @param array RedirectUrls

type SetRedirectUrlsForOAuthClientParameters struct {
	Identifier   string
	RedirectUrls []string
}

func (t *OAuth) SetRedirectUrlsForOAuthClient(p *SetRedirectUrlsForOAuthClientParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`identifier`, p.Identifier)
	for i := range p.RedirectUrls {
		queryParameters.Add(`redirectUrls[]`, p.RedirectUrls[i])
	}

	return t.restClient.Post(
		`/v2/OAuth/UseCase/SetRedirectUrlsForOAuthClient`,
		&queryParameters,
		nil,
		nil,
	)
}
