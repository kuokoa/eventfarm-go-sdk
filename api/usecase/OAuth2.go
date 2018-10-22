/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk"
)

type OAuth2 struct {
	restClient sdk.RestClientInterface
}

func NewOAuth2(restClient sdk.RestClientInterface) *OAuth2 {
	return &OAuth2{restClient}
}

// GET: Queries

// POST: Commands
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

// @param string Identifier
// @param array RedirectUrls

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
