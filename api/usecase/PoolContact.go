/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"bitbucket.ef.network/go/sdk"
)

type PoolContact struct {
	restClient sdk.RestClientInterface
}

func NewPoolContact(restClient sdk.RestClientInterface) *PoolContact {
	return &PoolContact{restClient}
}

// GET: Queries
// @param string UserId
// @param string|null PoolId
// @param string|null PoolContactType full|create

type ListPoolContactsForUserParameters struct {
	UserId          string
	PoolId          *string
	PoolContactType *string
}

func (t *PoolContact) ListPoolContactsForUser(p *ListPoolContactsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.PoolContactType != nil {
		queryParameters.Add(`poolContactType`, *p.PoolContactType)
	}

	return t.restClient.Get(
		`/v2/PoolContact/UseCase/ListPoolContactsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
