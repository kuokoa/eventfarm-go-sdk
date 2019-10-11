/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"github.com/eventfarm/go-sdk/rest"
	"net/http"
	"net/url"
	"strconv"
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
