/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type OAuth struct {
	restClient rest.RestClientInterface
}

func NewOAuth(restClient rest.RestClientInterface) *OAuth {
	return &OAuth{restClient}
}

// GET: Queries

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

func (t *OAuth) CreateGhostAccessTokenWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/OAuth/UseCase/CreateGhostAccessToken`,
		data,
		nil,
		nil,
	)
}
