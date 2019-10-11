/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type UserIdentifier struct {
	restClient rest.RestClientInterface
}

func NewUserIdentifier(restClient rest.RestClientInterface) *UserIdentifier {
	return &UserIdentifier{restClient}
}

// GET: Queries

// POST: Commands

type SetUserIdentifierParameters struct {
	UserIdentifierId string
	Identifier       string
	PoolId           *string
}

func (t *UserIdentifier) SetUserIdentifier(p *SetUserIdentifierParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userIdentifierId`, p.UserIdentifierId)
	queryParameters.Add(`identifier`, p.Identifier)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Post(
		`/v2/UserIdentifier/UseCase/SetUserIdentifier`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *UserIdentifier) SetUserIdentifierWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/UserIdentifier/UseCase/SetUserIdentifier`,
		data,
		nil,
		nil,
	)
}
