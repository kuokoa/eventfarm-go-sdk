/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type OAuth2 struct {
	restClient rest.RestClientInterface
}

func NewOAuth2(restClient rest.RestClientInterface) *OAuth2 {
	return &OAuth2{restClient}
}

// GET: Queries

// POST: Commands

type CreateOAuthClientParameters struct {
	Name         string
	RedirectUrls []string
	Identifier   *string
	Secret       *string
}

func (t *OAuth2) CreateOAuthClient(p *CreateOAuthClientParameters) (r *http.Response, err error) {
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
		`/v2/OAuth2/UseCase/CreateOAuthClient`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *OAuth2) CreateOAuthClientWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/OAuth2/UseCase/CreateOAuthClient`,
		data,
		nil,
		nil,
	)
}

type RevokeAccessTokenParameters struct {
	Identifier string
}

func (t *OAuth2) RevokeAccessToken(p *RevokeAccessTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`identifier`, p.Identifier)

	return t.restClient.Post(
		`/v2/OAuth2/UseCase/RevokeAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *OAuth2) RevokeAccessTokenWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/OAuth2/UseCase/RevokeAccessToken`,
		data,
		nil,
		nil,
	)
}

type SetRedirectUrlsForOAuthClientParameters struct {
	Identifier   string
	RedirectUrls []string
}

func (t *OAuth2) SetRedirectUrlsForOAuthClient(p *SetRedirectUrlsForOAuthClientParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`identifier`, p.Identifier)
	for i := range p.RedirectUrls {
		queryParameters.Add(`redirectUrls[]`, p.RedirectUrls[i])
	}

	return t.restClient.Post(
		`/v2/OAuth2/UseCase/SetRedirectUrlsForOAuthClient`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *OAuth2) SetRedirectUrlsForOAuthClientWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/OAuth2/UseCase/SetRedirectUrlsForOAuthClient`,
		data,
		nil,
		nil,
	)
}
