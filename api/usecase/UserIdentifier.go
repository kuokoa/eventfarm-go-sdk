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

type UserIdentifier struct {
	restClient rest.RestClientInterface
}

func NewUserIdentifier(restClient rest.RestClientInterface) *UserIdentifier {
	return &UserIdentifier{restClient}
}

// GET: Queries

// POST: Commands
// @param string UserIdentifierId
// @param string Identifier
// @param string|null PoolId

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
